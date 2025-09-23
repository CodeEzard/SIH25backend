package handlers

import (
	"encoding/json"
	"errors"
	"log"
	"net/http"
	"strconv"
	"vericred/internal/db"
	"vericred/internal/middleware"
	"vericred/internal/models"

	"gorm.io/gorm"
)

func CreateUniversity(w http.ResponseWriter, r *http.Request) {
	log.Println("CreateUniversity called")
	var body map[string]any
	if err := json.NewDecoder(r.Body).Decode(&body); err != nil {
		http.Error(w, "invalid JSON body", http.StatusBadRequest)
		return
	}

	acad_email, _ := body["AcadEmail"].(string)
	org_name, _ := body["OrgName"].(string)
	org_type, _ := body["OrgType"].(string)
	org_url, _ := body["OrgUrl"].(string)
	org_desc, _ := body["OrgDesc"].(string)
	country, _ := body["Country"].(string)
	state, _ := body["State"].(string)
	city, _ := body["City"].(string)
	total_students_str, _ := body["TotalStudents"].(string)
	address, _ := body["Address"].(string)
	postal_code, _ := body["PostalCode"].(string)

	students := 0
	if total_students_str != "" {
		if v, err := strconv.Atoi(total_students_str); err == nil {
			students = v
		}
	}

	metamaskAddress, ok := r.Context().Value(middleware.MetamaskAddressKey).(string)
	if !ok || metamaskAddress == "" {
		http.Error(w, "metamaskAddress is missing or invalid", http.StatusBadRequest)
		return
	}

	// Block if this wallet already has a student profile
	var existingUser models.Users
	if err := db.DB.Where("metamask_address = ?", metamaskAddress).First(&existingUser).Error; err == nil && existingUser.ID != 0 {
		w.WriteHeader(http.StatusConflict)
		_ = json.NewEncoder(w).Encode(map[string]any{
			"error":        "account_conflict",
			"message":      "Wallet already registered as a student. Use a different wallet to create a university account.",
			"account_type": "student",
		})
		return
	}

	// If org already exists for this wallet, return it (idempotent)
	var existingOrg models.Organization
	err := db.DB.Where("metamask_address = ?", metamaskAddress).First(&existingOrg).Error
	if err == nil {
		w.WriteHeader(http.StatusOK)
		_ = json.NewEncoder(w).Encode(map[string]any{
			"account_type": "university",
			"organization": existingOrg,
			"authStatus": map[string]any{
				"isAuthenticated": true,
				"accountType":    "university",
			},
		})
		return
	} else if !errors.Is(err, gorm.ErrRecordNotFound) {
		http.Error(w, "database error", http.StatusInternalServerError)
		return
	}

	org := models.Organization{
		MetamaskAddress: metamaskAddress,
		AcadEmail:       acad_email,
		OrgName:         org_name,
		OrgType:         org_type,
		OrgUrl:          org_url,
		OrgDesc:         org_desc,
		Country:         country,
		State:           state,
		City:            city,
		TotalStudents:   students,
		Address:         address,
		PostalCode:      postal_code,
	}
	if err := db.DB.Create(&org).Error; err != nil {
		http.Error(w, "failed to create organization", http.StatusInternalServerError)
		return
	}

	// Update account to point to this organization
	var account models.Accounts
	if err := db.DB.Where("metamask_address = ?", metamaskAddress).First(&account).Error; err != nil {
		http.Error(w, "account not found for wallet", http.StatusInternalServerError)
		return
	}
	account.OwnerID = org.ID
	account.OwnerType = "university"
	account.AccountType = "university"
	if err := db.DB.Save(&account).Error; err != nil {
		http.Error(w, "failed to update account", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
	_ = json.NewEncoder(w).Encode(map[string]any{
		"account_type": "university",
		"organization": org,
		"authStatus": map[string]any{
			"isAuthenticated": true,
			"accountType":    "university",
		},
	})
}

func ShowOrg(w http.ResponseWriter, r *http.Request) {
	metamaskAddress, ok := r.Context().Value(middleware.MetamaskAddressKey).(string)
	if !ok || metamaskAddress == "" {
		http.Error(w, "metamaskAddress is missing or invalid", http.StatusBadRequest)
		return
	}
	var org models.Organization
	res := db.DB.Where("metamask_address = ?", metamaskAddress).First(&org)
	if errors.Is(res.Error, gorm.ErrRecordNotFound) {
		w.WriteHeader(http.StatusNotFound)
		_ = json.NewEncoder(w).Encode(map[string]string{"error": "organization not found"})
		return
	} else if res.Error != nil {
		http.Error(w, "database error", http.StatusInternalServerError)
		return
	}
	_ = json.NewEncoder(w).Encode(org)
}

func AllOrgs(w http.ResponseWriter, r *http.Request) {
	var orgs []models.Organization
	result := db.DB.Limit(10).Find(&orgs)
	if result.Error != nil {
		w.WriteHeader(http.StatusNotFound)
		_ = json.NewEncoder(w).Encode(map[string]string{"error": "organization not found"})
		return
	}
	_ = json.NewEncoder(w).Encode(orgs)
}

func SpecificUniversity(w http.ResponseWriter, r *http.Request) {
	var body map[string]any
	if err := json.NewDecoder(r.Body).Decode(&body); err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}
	address, ok := body["metamask_address"].(string)
	if !ok || address == "" {
		http.Error(w, "Invalid or missing metamask_address", http.StatusBadRequest)
		return
	}
	var org models.Organization
	res := db.DB.Raw("SELECT * FROM organizations WHERE metamask_address = ?", address).Scan(&org)
	if errors.Is(res.Error, gorm.ErrRecordNotFound) {
		w.WriteHeader(http.StatusNotFound)
		_ = json.NewEncoder(w).Encode(map[string]string{"error": "organization not found"})
		return
	} else if res.Error != nil {
		http.Error(w, "Database error", http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusOK)
	_ = json.NewEncoder(w).Encode(org)
}