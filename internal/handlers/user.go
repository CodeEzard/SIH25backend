package handlers

import (
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"net/http"
	"vericred/internal/db"
	"vericred/internal/middleware"
	"vericred/internal/models"

	"gorm.io/gorm"
)

func CreateUser(w http.ResponseWriter, r *http.Request) {
	log.Println("CreateUser functin called.")
	var body map[string]any
	json.NewDecoder(r.Body).Decode(&body)

	email := body["email"].(string)
	first_name := body["firstName"].(string)
	last_name := body["lastName"].(string)
	student_id := body["studentEmail"].(string)

	metamaskAddress, ok := r.Context().Value(middleware.MetamaskAddressKey).(string)
	if !ok || metamaskAddress == "" {
		http.Error(w, "metamaskAddress is missing or invalid", http.StatusBadRequest)
		return
	}

	user := models.Users{
		MetamaskAddress: metamaskAddress,
		Email:           email,
		FirstName:       first_name,
		LastName:        last_name,
		StudentEmail:    student_id,
	}

	res := db.DB.Where("metamask_address = ?", metamaskAddress).First(&user)

	if errors.Is(res.Error, gorm.ErrRecordNotFound) {
		user.IsVerified = true
		var account models.Accounts

		if err := db.DB.Where("metamask_address = ?", metamaskAddress).First(&account).Error; err != nil {
			log.Fatalf("account not found: %v", err)
		}
		fmt.Println("account before modification: ", account)

		account.OwnerID = user.ID
		account.OwnerType = "user"

		fmt.Println("account after modification: ", account)
		user.Account = account
		if err := db.DB.Save(&account).Error; err != nil {
			log.Fatalf("Failed to update account: %v", err)
		}

		createResult := db.DB.Create(&user)
		
		if createResult.Error != nil {
			http.Error(w, "Failed to create account", http.StatusInternalServerError)
			return
		}


	} else if res.Error != nil {
		http.Error(w, "Database error", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(user)
}

func ShowUser(w http.ResponseWriter, r *http.Request) {
	metamaskAddress, ok := r.Context().Value(middleware.MetamaskAddressKey).(string)
	if !ok || metamaskAddress == "" {
		http.Error(w, "metamaskAddress is missing or invalid", http.StatusBadRequest)
	}
	log.Println("User logged in...")
	var user models.Users
	res := db.DB.Where("metamask_address = ?", metamaskAddress).First(&user)
	if res.Error == gorm.ErrRecordNotFound {
		w.WriteHeader(http.StatusNotFound)
		json.NewEncoder(w).Encode(map[string]string{"error": "user not found"})
		return
	} else if res.Error != nil {
		http.Error(w, "Database error", http.StatusInternalServerError)
		return
	}
	json.NewEncoder(w).Encode(user)
}

func AllUsers(w http.ResponseWriter, r *http.Request) {
	var orgs []models.Users
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

func SearchUser(w http.ResponseWriter, r *http.Request) {
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

	var user models.Users

	res := db.DB.Raw("SELECT * FROM users WHERE metamask_address = ?", address).Scan(&user)
	
	if res.Error == gorm.ErrRecordNotFound {
		w.WriteHeader(http.StatusNotFound)
		json.NewEncoder(w).Encode(map[string]string{"error": "organization not found"})
		return
	} else if res.Error != nil {
		http.Error(w, "Database error", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(user)
}