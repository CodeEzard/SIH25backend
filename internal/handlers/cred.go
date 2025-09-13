package handlers

import (
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"net/http"
	"strings"
	"time"
	"vericred/internal/db"
	"vericred/internal/middleware"
	"vericred/internal/models"

	"gorm.io/gorm"
)

func MintCredentials(w http.ResponseWriter, r *http.Request) {

	// fmt.Println("Inside the mint cred function")

	var body map[string]any
	if err := json.NewDecoder(r.Body).Decode(&body); err != nil {
		fmt.Println("Error in decoding: ", err)
		http.Error(w, "Failed to decode request body", http.StatusBadRequest)
		return
	}

	var cred models.Credential

	studentWallet, ok := body["student_wallet"].(string)
	if !ok || studentWallet == "" {
		fmt.Println("Invalid or missing 'student_wallet'")
		http.Error(w, "Invalid or missing 'student_wallet'", http.StatusBadRequest)
		return
	}
	cred.StudentWallet = studentWallet

	universityWallet, ok := body["university_wallet"].(string)
	if !ok || universityWallet == "" {
		fmt.Println("Invalid or missing 'university_wallet'")
		http.Error(w, "Invalid or missing 'university_wallet'", http.StatusBadRequest)
		return
	}
	cred.UniversityWallet = universityWallet

	degreeName, ok := body["degree_name"].(string)
	if !ok || degreeName == "" {
		fmt.Println("Invalid or missing 'degree_name'")
		http.Error(w, "Invalid or missing 'degree_name'", http.StatusBadRequest)
		return
	}
	cred.DegreeName = degreeName

	description, ok := body["description"].(string)
	if !ok {
		cred.Description = ""
	} else {
		cred.Description = description
	}

	credType, ok := body["type"].(string)
	if !ok || credType == "" {
		fmt.Println("Invalid or missing 'type'")
		http.Error(w, "Invalid or missing 'type'", http.StatusBadRequest)
		return
	}
	cred.Type = credType

	major, ok := body["major"].(string)
	if ok {
		cred.Major = major
	}

	// issued date: accept multiple formats (RFC3339, date-only, or datetime)
	issuedDateStr, ok := body["issued_date"].(string)
	if !ok || issuedDateStr == "" {
		http.Error(w, "Invalid or missing 'issued_date'", http.StatusBadRequest)
		return
	}

	var issuedDate time.Time
	var parseErr error

	issuedDate, parseErr = time.Parse(time.RFC3339, issuedDateStr)
	if parseErr != nil {
		issuedDate, parseErr = time.Parse("2006-01-02", issuedDateStr)
		if parseErr != nil {
			issuedDate, parseErr = time.Parse("2006-01-02 15:04:05", issuedDateStr)
			if parseErr != nil {
				http.Error(w, "Invalid 'issued_date' format; expected RFC3339 or YYYY-MM-DD", http.StatusBadRequest)
				return
			}
		}
	}

	issuedDate = issuedDate.UTC()
	cred.IssuedDate = issuedDate

	graduationDate, ok := body["graduation_date"].(string)
	if ok {
		cred.GraduationDate = graduationDate
	}

	ipfsLink, ok := body["ipfs_link"].(string)
	if !ok || ipfsLink == "" {
		fmt.Println("Invalid or missing 'ipfs_link'")
		http.Error(w, "Invalid or missing 'ipfs_link'", http.StatusBadRequest)
		return
	}
	cred.IPFSLink = strings.TrimSpace(ipfsLink)

	deanSig, ok := body["dean_sig"].(string)
	if !ok || deanSig == "" {
		fmt.Println("Invalid or missing 'dean_sig'")
		http.Error(w, "Invalid or missing 'dean_sig'", http.StatusBadRequest)
		return
	}
	cred.DeanSig = deanSig

	var org models.Organization
	if err := db.DB.Where("metamask_address = ?", cred.UniversityWallet).First(&org).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			fmt.Println("University not registered")
			http.Error(w, "University not registered", http.StatusBadRequest)
			return
		}
		fmt.Println("DB error finding university:", err)
		http.Error(w, "Database error", http.StatusInternalServerError)
		return
	}

	var user models.Users
	if err := db.DB.Where("metamask_address = ?", cred.StudentWallet).First(&user).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			fmt.Println("Student not registered")
			http.Error(w, "Student not registered", http.StatusBadRequest)
			return
		}
		fmt.Println("DB error finding student:", err)
		http.Error(w, "Database error", http.StatusInternalServerError)
		return
	}

	fmt.Println("UserID: ", user.ID)
	fmt.Println("OrgID: ", org.ID)
	cred.UserID = user.ID
	cred.OrganizationID = org.ID

	log.Printf("Creating credential id=%s user_id=%d org_id=%d", cred.ID, cred.UserID, cred.OrganizationID)

	tx := db.DB.Begin()
	if err := tx.Create(&cred).Error; err != nil {
		tx.Rollback()
		fmt.Println("Failed to create credential:", err)
		res := fmt.Sprint("Failed to create credential: ", err)
		http.Error(w, res, http.StatusInternalServerError)
		return
	}
	if err := tx.Commit().Error; err != nil {
		fmt.Println("Transaction commit failed:", err)
		http.Error(w, "Database error", http.StatusInternalServerError)
		return
	}

	if err := db.DB.Preload("User").Preload("Organization").Where("id = ?", cred.ID).First(&cred).Error; err != nil {
		fmt.Println("Preload error:", err)
		http.Error(w, "Error when preloading", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(cred)
}

func UserCreds(w http.ResponseWriter, r *http.Request) {
	metamaskAddress, ok := r.Context().Value(middleware.MetamaskAddressKey).(string)
	if !ok || metamaskAddress == "" {
		fmt.Println("metamaskAddress is missing or invalid")
		http.Error(w, "metamaskAddress is missing or invalid", http.StatusBadRequest)
		return
	}
	var cred []models.Credential
	res := db.DB.Where("student_wallet = ?", metamaskAddress).Find(&cred)
	if res.Error == gorm.ErrRecordNotFound {
		fmt.Println("User not found")
		w.WriteHeader(http.StatusNotFound)
		json.NewEncoder(w).Encode(map[string]string{"error": "user not found"})
		return
	} else if res.Error != nil {
		fmt.Println("Database error:", res.Error)
		http.Error(w, "Database error", http.StatusInternalServerError)
		return
	}
	json.NewEncoder(w).Encode(cred)
}



func ShowSearchedUserCreds(w http.ResponseWriter, r *http.Request) {
	var body map[string]any
	err := json.NewDecoder(r.Body).Decode(&body)
	
	if err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	address, ok := body["metamask_address"].(string)
	if !ok || address == "" {
		http.Error(w, "Invalid or missing metamask_address", http.StatusBadRequest)
		return
	}

	var cred []models.Credential
	res := db.DB.Where("student_wallet = ?", address).Find(&cred)
	if res.Error == gorm.ErrRecordNotFound {
		fmt.Println("User not found")
		w.WriteHeader(http.StatusNotFound)
		json.NewEncoder(w).Encode(map[string]string{"error": "user not found"})
		return
	} else if res.Error != nil {
		fmt.Println("Database error:", res.Error)
		http.Error(w, "Database error", http.StatusInternalServerError)
		return
	}
	json.NewEncoder(w).Encode(cred)
}
