package handlers

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"
	"vericred/internal/db"
	"vericred/internal/middleware"
	"vericred/internal/models"

	"gorm.io/gorm"
)


func CreateUniversity(w http.ResponseWriter, r *http.Request) {

	var body map[string]any
    err := json.NewDecoder(r.Body).Decode(&body)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	acad_email, ok := body["AcadEmail"].(string)
	if !ok {
		http.Error(w, "Invalid or missing AcadEmail", http.StatusBadRequest)
		return
	}
	org_name, ok := body["OrgName"].(string)
	if !ok {
		http.Error(w, "Invalid or missing OrgName", http.StatusBadRequest)
		return
	}
	org_type, ok := body["OrgType"].(string)
	if !ok {
		http.Error(w, "Invalid or missing OrgType", http.StatusBadRequest)
		return
	}
	org_url, ok := body["OrgUrl"].(string)
	if !ok {
		http.Error(w, "Invalid or missing OrgUrl", http.StatusBadRequest)
		return
	}
	org_desc, ok := body["OrgDesc"].(string)
	if !ok {
		http.Error(w, "Invalid or missing OrgDesc", http.StatusBadRequest)
		return
	}
	country, ok := body["Country"].(string)
	if !ok {
		http.Error(w, "Invalid or missing Country", http.StatusBadRequest)
		return
	}
	state, ok := body["State"].(string)
	if !ok {
		http.Error(w, "Invalid or missing State", http.StatusBadRequest)
		return
	}
	city, ok := body["City"].(string)
	if !ok {
		http.Error(w, "Invalid or missing City", http.StatusBadRequest)
		return
	}
	total_students, ok := body["TotalStudents"].(string) // Expecting a string
	if !ok {
		http.Error(w, "Invalid or missing TotalStudents", http.StatusBadRequest)
		return
	}

	students, err := strconv.Atoi(total_students) // Convert string to int
	if err != nil {
		http.Error(w, "TotalStudents must be a valid integer", http.StatusBadRequest)
		return
	}
	
	address, ok := body["Address"].(string)
	if !ok {
		http.Error(w, "Invalid or missing Address", http.StatusBadRequest)
		return
	}
	postal_code, ok := body["PostalCode"].(string)
	if !ok {
		http.Error(w, "Invalid or missing PostalCode", http.StatusBadRequest)
		return
	}


	metamaskAddress, ok := r.Context().Value(middleware.MetamaskAddressKey).(string)
	if !ok || metamaskAddress == "" {
		http.Error(w, "metamaskAddress is missing or invalid", http.StatusBadRequest)
		return
	}

	org := models.Organization{
		MetamaskAddress: metamaskAddress,
		AcadEmail: acad_email,
		OrgName: org_name,
		OrgType: org_type,
		OrgUrl: org_url,
		OrgDesc: org_desc,
		Country: country,
		State: state,
		City: city,
		TotalStudents: students,
		Address: address,
		PostalCode: postal_code,
	}

	
	res := db.DB.Where("metamask_address = ?", metamaskAddress).First(&org)

	if res.Error == gorm.ErrRecordNotFound {
		org.IsVerified = true
		createResult := db.DB.Create(&org)
        if createResult.Error != nil {
            http.Error(w, "Failed to create account", http.StatusInternalServerError)
            return
        }
		
		var account models.Accounts

		if err := db.DB.Where("metamask_address = ?", metamaskAddress).First(&account).Error; err != nil {
			log.Fatalf("account not found: %v", err)
		}

		account.OwnerID = org.ID
		account.OwnerType = "university"

		if err := db.DB.Save(&account).Error; err != nil {
			log.Fatalf("Failed to update account: %v", err)
		}
	
	} else if res.Error != nil {
        http.Error(w, "Database error", http.StatusInternalServerError)
        return
    }
	
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(org)
}

func ShowOrg(w http.ResponseWriter, r *http.Request) {
	metamaskAddress, ok := r.Context().Value(middleware.MetamaskAddressKey).(string)
	if !ok || metamaskAddress == "" {
		http.Error(w, "metamaskAddress is missing or invalid", http.StatusBadRequest)
		return
	}
	var org models.Organization
	res := db.DB.Where("metamask_address = ?", metamaskAddress).First(&org)
	if res.Error == gorm.ErrRecordNotFound {
		fmt.Println("inside not found university")
		w.WriteHeader(http.StatusNotFound)
		json.NewEncoder(w).Encode(map[string]string{"error": "organization not found"})
		return
	} else if res.Error != nil {
		fmt.Println("Database error")
		http.Error(w, "Database error", http.StatusInternalServerError)
		return
	}
	json.NewEncoder(w).Encode(org)
}

func AllOrgs(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Allorgs function called.")
	var orgs []models.Organization
	result := db.DB.Limit(10).Find(&orgs) 
	if result.Error != nil {
		w.WriteHeader(http.StatusNotFound)
		json.NewEncoder(w).Encode(map[string]string{"error": "user not found"})
		return
	} 
	json.NewEncoder(w).Encode(orgs)
	// for _, org := range orgs {
	// 	fmt.Printf("Organization: %+v\n", org)
	// 
}

func SpecificUniversity(w http.ResponseWriter, r *http.Request) {
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

	var org models.Organization
	
	res := db.DB.Raw("SELECT * FROM organizations WHERE metamask_address = ?", address).Scan(&org)
	
	if res.Error == gorm.ErrRecordNotFound {
		w.WriteHeader(http.StatusNotFound)
		json.NewEncoder(w).Encode(map[string]string{"error": "organization not found"})
		return
	} else if res.Error != nil {
		http.Error(w, "Database error", http.StatusInternalServerError)
		return
	}
	fmt.Print(org)
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(org)
}