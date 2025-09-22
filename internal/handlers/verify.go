package handlers

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
	"regexp"
	"strings"

	vision "cloud.google.com/go/vision/apiv1"
	visionpb "cloud.google.com/go/vision/v2/apiv1/visionpb"

	"vericred/internal/db"
	"vericred/internal/models"

	"github.com/adrg/strutil"
	"github.com/adrg/strutil/metrics"
	"google.golang.org/api/option"
	"gorm.io/gorm"
)

func writeJSONResp(w http.ResponseWriter, status int, payload any) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	_ = json.NewEncoder(w).Encode(payload)
}

// VerifyDocument: POST /api/v1/verify-document
// multipart/form-data with file field "certificate"
func VerifyDocument(w http.ResponseWriter, r *http.Request) {
	// Limit body to 10MB
	r.Body = http.MaxBytesReader(w, r.Body, 10<<20)
	if err := r.ParseMultipartForm(10 << 20); err != nil {
		writeJSONResp(w, http.StatusBadRequest, map[string]any{"status": "Bad_Request", "message": "failed to parse form or file too large"})
		return
	}

	file, _, err := r.FormFile("certificate")
	if err != nil {
		// Debug: list available file keys and try alternatives
		if r.MultipartForm != nil && r.MultipartForm.File != nil {
			keys := make([]string, 0, len(r.MultipartForm.File))
			for k := range r.MultipartForm.File {
				keys = append(keys, k)
			}
			fmt.Println("verify: available multipart file fields:", keys)
			alts := []string{"file", "upload", "image", "document", "cert", "certificateFile", "certificate[]", "files[]"}
			for _, a := range alts {
				if f2, _, err2 := r.FormFile(a); err2 == nil {
					file, err = f2, nil
					fmt.Println("verify: using alternative file field:", a)
					break
				}
			}
			if err != nil && len(keys) > 0 {
				k0 := keys[0]
				if f2, _, err2 := r.FormFile(k0); err2 == nil {
					file, err = f2, nil
					fmt.Println("verify: falling back to first file field:", k0)
				}
			}
		}
		if err != nil {
			writeJSONResp(w, http.StatusBadRequest, map[string]any{"status": "Bad_Request", "message": "missing file field 'certificate' (send multipart/form-data with field name 'certificate')"})
			return
		}
	}
	defer file.Close()

	imgBytes, err := io.ReadAll(file)
	if err != nil || len(imgBytes) == 0 {
		writeJSONResp(w, http.StatusBadRequest, map[string]any{"status": "Bad_Request", "message": "failed to read uploaded file"})
		return
	}

	// OCR with Google Vision
	ctx := context.Background()
	credPath := os.Getenv("GOOGLE_APPLICATION_CREDENTIALS")
	var client *vision.ImageAnnotatorClient
	if credPath != "" {
		client, err = vision.NewImageAnnotatorClient(ctx, option.WithCredentialsFile(credPath))
	} else {
		client, err = vision.NewImageAnnotatorClient(ctx)
	}
	msg := fmt.Sprintf("failed to init OCR client %s", err)
	if err != nil {
		writeJSONResp(w, http.StatusInternalServerError, map[string]any{"status": "Server_Error", "message": msg})
		return
	}
	defer client.Close()

	img := &visionpb.Image{Content: imgBytes}
	anns, err := client.DetectTexts(ctx, img, nil, 1)
	msg = fmt.Sprintf("could not extract text from image %s", err)
	if err != nil || len(anns) == 0 || anns[0].Description == "" {
		writeJSONResp(w, http.StatusBadRequest, map[string]any{"status": "Bad_Request", "message": msg})
		return
	}
	raw := anns[0].Description

	// Replace regex parser with Gemini-based parser
	pc, perr := ParseWithGemini(raw)
	if perr != nil {
		writeJSONResp(w, http.StatusBadRequest, map[string]any{"status": "Bad_Request", "message": perr.Error()})
		return
	}

	// Lookup record by register/roll number (treat as roll_number here)
	var rec models.LegacyCredential
	err = db.DB.Preload("University").Where("roll_number = ?", pc.RegisterNumber).First(&rec).Error
	if err == gorm.ErrRecordNotFound {
		writeJSONResp(w, http.StatusOK, map[string]any{
			"status":  "Not_Found",
			"message": "No matching record was found for the provided register number.",
		})
		return
	} else if err != nil {
		writeJSONResp(w, http.StatusInternalServerError, map[string]any{"status": "Server_Error", "message": "database error"})
		return
	}

	official := strings.TrimSpace(rec.University.OrgName)
	extracted := strings.TrimSpace(pc.UniversityName)

	// Fuzzy compare with Jaro-Winkler
	metric := metrics.NewJaroWinkler()
	conf := strutil.Similarity(strings.ToLower(extracted), strings.ToLower(official), metric)

	data := map[string]any{
		"student_name_ocr":    pc.StudentName,
		"register_number":     pc.RegisterNumber,
		"course_name":         pc.CourseName,
		"year_of_passing":     pc.YearOfPassing,
		"university_name_ocr": extracted,
		"official_university": official,
		"record":              rec,
	}

	if conf >= 0.85 {
		writeJSONResp(w, http.StatusOK, map[string]any{
			"status":           "Verified",
			"match_confidence": conf,
			"data":             data,
		})
		return
	}

	writeJSONResp(w, http.StatusOK, map[string]any{
		"status":           "Potentially_Tampered",
		"match_confidence": conf,
		"message":          "The institution name on the document does not match the official record.",
		"data":             data,
	})
}

var (
	rollRe = regexp.MustCompile(`(?i)\broll\s*(no\.?|number|num|#)?\s*[:\-]?\s*([A-Z0-9\-_/]+)`) // capture roll no formats
	nameRe = regexp.MustCompile(`(?i)\b(student\s*)?name\s*[:\-]?\s*([A-Za-z][A-Za-z .'-]{2,})`)
)

func parseCertificateText(raw string) (studentName, rollNumber, universityName string) {
	lines := strings.Split(raw, "\n")
	// Try regex extraction
	for _, ln := range lines {
		l := strings.TrimSpace(ln)
		if rollNumber == "" {
			if m := rollRe.FindStringSubmatch(l); len(m) >= 3 {
				rollNumber = strings.TrimSpace(m[2])
			}
		}
		if studentName == "" {
			if m := nameRe.FindStringSubmatch(l); len(m) >= 3 {
				studentName = strings.TrimSpace(m[2])
			}
		}
	}

	// Heuristic university name: pick the longest line containing keywords
	keywords := []string{"university", "institute", "college", "academy"}
	best := ""
	for _, ln := range lines {
		l := strings.TrimSpace(ln)
		ll := strings.ToLower(l)
		for _, kw := range keywords {
			if strings.Contains(ll, kw) {
				if len(l) > len(best) {
					best = l
				}
				break
			}
		}
	}
	universityName = best
	return
}
