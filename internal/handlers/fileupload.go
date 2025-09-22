package handlers

import (
	"encoding/csv"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"mime/multipart"
	"net/http"
	"strconv"
	"strings"
	"time"

	"vericred/internal/db"
	"vericred/internal/middleware"
	"vericred/internal/models"
)

// writeJSON is a small helper to return JSON responses consistently.
// func writeJSON(w http.ResponseWriter, status int, payload any) {
// 	w.Header().Set("Content-Type", "application/json")
// 	w.WriteHeader(status)
// 	_ = json.NewEncoder(w).Encode(payload)
// }

// func writeError(w http.ResponseWriter, status int, msg string) {
// 	writeJSON(w, status, map[string]any{"error": msg})
// }

// BulkUploadHandler handles CSV bulk upload of legacy credentials by an authenticated university admin.
func BulkUploadHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	fmt.Println("Inside bulk upload handler.")
	// 1) Ensure auth context has metamask address, then resolve to Organization
	metamaskAddress, ok := r.Context().Value(middleware.MetamaskAddressKey).(string)
	if !ok || metamaskAddress == "" {
		http.Error(w, "unauthorized", http.StatusUnauthorized)
		fmt.Println("Error:", "unauthorized")
		return
	}

	var org models.Organization
	if err := db.DB.Where("metamask_address = ?", metamaskAddress).First(&org).Error; err != nil {
		http.Error(w, "organization not found", http.StatusForbidden)
		fmt.Println("Error:", "organization not found")
		return
	}

	// 2) Parse multipart with a 50MB limit
	if err := r.ParseMultipartForm(50 << 20); err != nil {
		http.Error(w, "failed to parse form", http.StatusBadRequest)
		fmt.Println("Error:", "failed to parse form")
		return
	}

	// Tolerant file field lookup: prefer "recordsCsv", but try alternatives and fallback to first file field.
	var file multipart.File
	var header *multipart.FileHeader
	var err error

	file, header, err = r.FormFile("recordsCsv")
	if err != nil {
		alts := []string{"records", "csv", "file", "upload", "records_file", "recordsCSV", "recordsCsv[]", "files[]"}
		available := []string{}
		if r.MultipartForm != nil && r.MultipartForm.File != nil {
			for k := range r.MultipartForm.File {
				available = append(available, k)
			}
		}
		fmt.Println("bulk-upload: available multipart file fields:", available)

		// Try alternatives (case-insensitive match against available keys)
		lookup := func(name string) (multipart.File, *multipart.FileHeader, error) {
			if f, h, e := r.FormFile(name); e == nil {
				return f, h, nil
			}
			// case-insensitive search across available keys
			lname := strings.ToLower(name)
			for _, k := range available {
				if strings.ToLower(k) == lname {
					return r.FormFile(k)
				}
			}
			return nil, nil, fmt.Errorf("not found")
		}
		for _, a := range alts {
			if f2, h2, e2 := lookup(a); e2 == nil {
				file, header, err = f2, h2, nil
				fmt.Println("bulk-upload: using alternative file field:", a)
				break
			}
		}
		// Fallback to the first available file field
		if err != nil && len(available) > 0 {
			k0 := available[0]
			if f2, h2, e2 := r.FormFile(k0); e2 == nil {
				file, header, err = f2, h2, nil
				fmt.Println("bulk-upload: falling back to first file field:", k0)
			}
		}
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			_ = json.NewEncoder(w).Encode(map[string]any{
				"error":               "recordsCsv file is required",
				"expected_field":      "recordsCsv",
				"available_file_keys": available,
			})
			fmt.Println("Error:", "recordsCsv file is required")
			return
		}
	}
	defer file.Close()

	// 3) CSV reader and header validation
	reader := csv.NewReader(file)
	reader.TrimLeadingSpace = true
	reader.FieldsPerRecord = -1 // allow variable-length; we'll validate

	requiredHeaders := []string{"student_name", "roll_number", "program", "major", "batch_year", "issued_date", "graduation_date"}
	headers, err := reader.Read()
	if err != nil {
		http.Error(w, "unable to read CSV header", http.StatusBadRequest)
		fmt.Println("Error:", "unable to read CSV header")
		return
	}
	for i := range headers {
		headers[i] = strings.TrimSpace(strings.ToLower(headers[i]))
	}
	if !equalStringSlices(headers, requiredHeaders) {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(map[string]any{
			"error":    "Invalid CSV format. Please use the provided template.",
			"expected": requiredHeaders,
			"got":      headers,
		})
		fmt.Println("Error:", "Invalid CSV format. Please use the provided template.")
		return
	}

	// 4) Begin transaction
	tx := db.DB.Begin()
	if tx.Error != nil {
		http.Error(w, "could not start transaction", http.StatusInternalServerError)
		fmt.Println("Error:", "could not start transaction")
		return
	}
	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
		}
	}()

	// 5) Read and insert rows
	var count int
	var duplicates int
	for {
		rec, err := reader.Read()
		if errors.Is(err, io.EOF) {
			break
		}
		if err != nil {
			tx.Rollback()
			http.Error(w, "failed to read CSV rows", http.StatusBadRequest)
			fmt.Println("Error:", "failed to read CSV rows")
			return
		}
		// Expect len(rec) == len(requiredHeaders)
		if len(rec) != len(requiredHeaders) {
			tx.Rollback()
			http.Error(w, "row does not match header length", http.StatusBadRequest)
			fmt.Println("Error:", "row does not match header length")
			return
		}

		studentName := strings.TrimSpace(rec[0])
		rollNumber := strings.TrimSpace(rec[1])
		program := strings.TrimSpace(rec[2])
		major := strings.TrimSpace(rec[3])
		batchYearStr := strings.TrimSpace(rec[4])
		issuedDateStr := strings.TrimSpace(rec[5])
		graduationDate := strings.TrimSpace(rec[6])

		var batchYear int
		if batchYearStr != "" {
			by, err := strconv.Atoi(batchYearStr)
			if err != nil {
				tx.Rollback()
				http.Error(w, "invalid batch_year", http.StatusBadRequest)
				fmt.Println("Error:", "invalid batch_year")
				return
			}
			batchYear = by
		}

		var issuedDatePtr *time.Time
		if issuedDateStr != "" {
			// expect YYYY-MM-DD
			if t, err := time.Parse("2006-01-02", issuedDateStr); err == nil {
				issuedDatePtr = &t
			} else {
				tx.Rollback()
				http.Error(w, "invalid issued_date (expected YYYY-MM-DD)", http.StatusBadRequest)
				fmt.Println("Error:", "invalid issued_date (expected YYYY-MM-DD)")
				return
			}
		}

		// Duplicate check: same roll_number for this university
		var dup int64
		if err := tx.Model(&models.LegacyCredential{}).
			Where("roll_number = ? AND university_id = ?", rollNumber, org.ID).
			Count(&dup).Error; err != nil {
			tx.Rollback()
			http.Error(w, "database error during duplicate check", http.StatusInternalServerError)
			fmt.Println("Error:", "database error during duplicate check")
			return
		}
		if dup > 0 {
			duplicates++
			continue
		}

		row := models.LegacyCredential{
			StudentName:    studentName,
			RollNumber:     rollNumber,
			Program:        program,
			Major:          major,
			BatchYear:      batchYear,
			IssuedDate:     issuedDatePtr,
			GraduationDate: graduationDate,
			UniversityID:   org.ID,
		}

		if err := tx.Create(&row).Error; err != nil {
			tx.Rollback()
			http.Error(w, "failed to insert row", http.StatusInternalServerError)
			fmt.Println("Error:", "failed to insert row")
			return
		}
		count++
	}

	if err := tx.Commit().Error; err != nil {
		tx.Rollback()
		http.Error(w, "failed to commit transaction", http.StatusInternalServerError)
		fmt.Println("Error:", "failed to commit transaction")
		return
	}

	json.NewEncoder(w).Encode(map[string]any{
		"message":             fmt.Sprintf("Successfully imported %d records. Skipped %d duplicates.", count, duplicates),
		"inserted":            count,
		"duplicates_skipped":  duplicates,
		"file":                header.Filename,
	})
}

func equalStringSlices(a, b []string) bool {
	if len(a) != len(b) {
		return false
	}
	for i := range a {
		if a[i] != b[i] {
			return false
		}
	}
	return true
}