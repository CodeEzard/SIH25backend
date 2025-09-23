package handlers

import (
	"encoding/json"
	"errors"
	"log"
	"net/http"
	"vericred/internal/db"
	"vericred/internal/middleware"
	"vericred/internal/models"

	"gorm.io/gorm"
)

func CreateUser(w http.ResponseWriter, r *http.Request) {
	log.Println("CreateUser called")
	var body map[string]any
	if err := json.NewDecoder(r.Body).Decode(&body); err != nil {
		http.Error(w, "invalid JSON body", http.StatusBadRequest)
		return
	}

	email, _ := body["email"].(string)
	firstName, _ := body["firstName"].(string)
	lastName, _ := body["lastName"].(string)
	studentEmail, _ := body["studentEmail"].(string)

	metamaskAddress, ok := r.Context().Value(middleware.MetamaskAddressKey).(string)
	if !ok || metamaskAddress == "" {
		http.Error(w, "metamaskAddress is missing or invalid", http.StatusBadRequest)
		return
	}

	// Block if this wallet already mapped to a university
	var existingOrg models.Organization
	if err := db.DB.Where("metamask_address = ?", metamaskAddress).First(&existingOrg).Error; err == nil && existingOrg.ID != 0 {
		w.WriteHeader(http.StatusConflict)
		_ = json.NewEncoder(w).Encode(map[string]any{
			"error":        "account_conflict",
			"message":      "Wallet already registered as a university. Use a different wallet to create a student account.",
			"account_type": "university",
		})
		return
	}

	// If a user already exists for this wallet, return it (idempotent)
	var existingUser models.Users
	err := db.DB.Where("metamask_address = ?", metamaskAddress).First(&existingUser).Error
	if err == nil {
		w.WriteHeader(http.StatusOK)
		_ = json.NewEncoder(w).Encode(map[string]any{
			"account_type": "student",
			"user":         existingUser,
			"authStatus": map[string]any{
				"isAuthenticated": true,
				"accountType":    "student",
			},
		})
		return
	} else if !errors.Is(err, gorm.ErrRecordNotFound) {
		http.Error(w, "database error", http.StatusInternalServerError)
		return
	}

	// Create user first
	newUser := models.Users{
		MetamaskAddress: metamaskAddress,
		Email:           email,
		FirstName:       firstName,
		LastName:        lastName,
		StudentEmail:    studentEmail,
		IsVerified:      true,
	}
	if err := db.DB.Create(&newUser).Error; err != nil {
		http.Error(w, "failed to create user", http.StatusInternalServerError)
		return
	}

	// Update existing account to point to this user
	var account models.Accounts
	if err := db.DB.Where("metamask_address = ?", metamaskAddress).First(&account).Error; err != nil {
		http.Error(w, "account not found for wallet", http.StatusInternalServerError)
		return
	}
	account.OwnerID = newUser.ID
	account.OwnerType = "user"
	account.AccountType = "student"
	if err := db.DB.Save(&account).Error; err != nil {
		http.Error(w, "failed to update account", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
	_ = json.NewEncoder(w).Encode(map[string]any{
		"account_type": "student",
		"user":         newUser,
		"authStatus": map[string]any{
			"isAuthenticated": true,
			"accountType":    "student",
		},
	})
}

func ShowUser(w http.ResponseWriter, r *http.Request) {
	metamaskAddress, ok := r.Context().Value(middleware.MetamaskAddressKey).(string)
	if !ok || metamaskAddress == "" {
		http.Error(w, "metamaskAddress is missing or invalid", http.StatusBadRequest)
		return
	}
	log.Println("User logged in...")
	var user models.Users
	res := db.DB.Where("metamask_address = ?", metamaskAddress).First(&user)
	if errors.Is(res.Error, gorm.ErrRecordNotFound) {
		w.WriteHeader(http.StatusNotFound)
		_ = json.NewEncoder(w).Encode(map[string]string{"error": "user not found"})
		return
	} else if res.Error != nil {
		http.Error(w, "database error", http.StatusInternalServerError)
		return
	}
	_ = json.NewEncoder(w).Encode(map[string]any{
		"account_type": "student",
		"user":         user,
		"authStatus": map[string]any{
			"isAuthenticated": true,
			"accountType":    "student",
		},
	})
}

func AllUsers(w http.ResponseWriter, r *http.Request) {
	var users []models.Users
	result := db.DB.Limit(10).Find(&users)
	if result.Error != nil {
		w.WriteHeader(http.StatusNotFound)
		_ = json.NewEncoder(w).Encode(map[string]string{"error": "user not found"})
		return
	}
	_ = json.NewEncoder(w).Encode(users)
}

func SearchUser(w http.ResponseWriter, r *http.Request) {
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
	var user models.Users
	res := db.DB.Raw("SELECT * FROM users WHERE metamask_address = ?", address).Scan(&user)
	if errors.Is(res.Error, gorm.ErrRecordNotFound) {
		w.WriteHeader(http.StatusNotFound)
		_ = json.NewEncoder(w).Encode(map[string]string{"error": "user not found"})
		return
	} else if res.Error != nil {
		http.Error(w, "Database error", http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusOK)
	_ = json.NewEncoder(w).Encode(user)
}