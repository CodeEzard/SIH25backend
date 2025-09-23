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
	fmt.Println("GEMINI OUTPUT: ", pc)
	
	// Fetch possible matches. Prefer exact roll match, but also allow fuzzy fallback by name/university
	var candidates []models.LegacyCredential
	if strings.TrimSpace(pc.RegisterNumber) != "" {
		_ = db.DB.Preload("University").Where("roll_number = ?", pc.RegisterNumber).Find(&candidates).Error
	}
	if len(candidates) == 0 {
		// Fallback: try by partial name/university to aid matching
		nameLike := "%" + strings.ToLower(strings.TrimSpace(pc.StudentName)) + "%"
		uniLike := "%" + strings.ToLower(strings.TrimSpace(pc.UniversityName)) + "%"
		_ = db.DB.Preload("University").Joins("JOIN organizations ON organizations.id = legacy_credentials.university_id").
			Where("LOWER(legacy_credentials.student_name) LIKE ? OR LOWER(organizations.org_name) LIKE ?", nameLike, uniLike).
			Limit(10).Find(&candidates).Error
	}

	if len(candidates) == 0 {
		writeJSONResp(w, http.StatusOK, map[string]any{
			"status":  "Not_Found",
			"message": "No matching record was found for the provided details.",
		})
		return
	}

	// Choose best candidate by combined similarity
	metric := metrics.NewJaroWinkler()
	bestIdx := -1
	bestScore := -1.0
	bestNameSim := 0.0
	bestUniSim := 0.0
	for i, rec := range candidates {
		officialUni := strings.TrimSpace(rec.University.OrgName)
		// Canonicalize university names to mitigate OCR typos and location suffixes
		uniSim := strutil.Similarity(canonUniName(pc.UniversityName), canonUniName(officialUni), metric)
		nameSim := strutil.Similarity(norm(pc.StudentName), norm(rec.StudentName), metric)
		rollBonus := 0.0
		if strings.EqualFold(strings.TrimSpace(pc.RegisterNumber), strings.TrimSpace(rec.RollNumber)) {
			rollBonus = 0.10
		}
		score := 0.45*uniSim + 0.4*nameSim + rollBonus
		if score > bestScore {
			bestScore = score
			bestIdx = i
			bestNameSim = nameSim
			bestUniSim = uniSim
		}
	}

	rec := candidates[bestIdx]
	officialUni := strings.TrimSpace(rec.University.OrgName)
	rollMatch := strings.EqualFold(strings.TrimSpace(pc.RegisterNumber), strings.TrimSpace(rec.RollNumber))

	data := map[string]any{
		"student_name_ocr":        pc.StudentName,
		"register_number":         pc.RegisterNumber,
		"course_name":             pc.CourseName,
		"year_of_passing":         pc.YearOfPassing,
		"university_name_ocr":     strings.TrimSpace(pc.UniversityName),
		"official_university":     officialUni,
		"official_student_name":   rec.StudentName,
		"record":                  rec,
		"name_confidence":         bestNameSim,
		"university_confidence":   bestUniSim,
		"roll_number_exact_match": rollMatch,
	}

	// Adaptive verification policy (A):
	// - Strict: roll exact + name >= 0.95 + university >= 0.90
	// - Adaptive: if roll exact + name >= 0.98, accept university >= 0.88
	strictName := bestNameSim >= 0.95
	strictUni := bestUniSim >= 0.90
	veryHighName := bestNameSim >= 0.98
	adaptiveUni := bestUniSim >= 0.88
	if rollMatch && ((strictName && strictUni) || (veryHighName && adaptiveUni)) {
		writeJSONResp(w, http.StatusOK, map[string]any{
			"status":             "Verified",
			"overall_confidence": bestScore,
			"data":               data,
		})
		return
	}

	// Explain exactly what failed
	reasons := []string{}
	if !rollMatch { reasons = append(reasons, "Roll number does not match the official record") }
	if !(strictName || veryHighName) { reasons = append(reasons, "Student name does not closely match the official record") }
	if !(strictUni || (veryHighName && adaptiveUni)) { reasons = append(reasons, "Institution name does not closely match the official record") }

	writeJSONResp(w, http.StatusOK, map[string]any{
		"status":             "Potentially_Tampered",
		"overall_confidence": bestScore,
		"message":            strings.Join(reasons, "; "),
		"data":               data,
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

// norm collapses spaces, removes punctuation, and lowercases for robust comparisons
func norm(s string) string {
	s = strings.ToLower(strings.TrimSpace(s))
	// Replace common punctuation with space
	replacer := strings.NewReplacer(".", " ", ",", " ", "-", " ", "_", " ", "'", " ", "\"", " ")
	s = replacer.Replace(s)
	// Collapse multiple spaces
	parts := strings.Fields(s)
	return strings.Join(parts, " ")
}

// canonUniName normalizes and fixes common OCR typos, and removes location suffixes/campus markers (C)
func canonUniName(s string) string {
	n := norm(s)
	// Fix frequent OCR/typo variants
	repls := map[string]string{
		"tichology": "technology",
		"techology": "technology",
		"technolgy": "technology",
		"institue":  "institute",
		"instittute": "institute",
		"inistute":  "institute",
	}
	for from, to := range repls {
		n = strings.ReplaceAll(n, from, to)
	}
	// Remove common location/campus words
	stops := map[string]struct{}{
		"mesra": {},
		"ranchi": {},
		"jharkhand": {},
		"india": {},
		"campus": {},
	}
	toks := strings.Fields(n)
	keep := make([]string, 0, len(toks))
	for _, t := range toks {
		if _, blocked := stops[t]; blocked { continue }
		keep = append(keep, t)
	}
	return strings.Join(keep, " ")
}
