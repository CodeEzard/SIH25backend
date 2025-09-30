package handlers

import (
	"encoding/json"
	"net/http"

	"vericred/internal/db"
	"vericred/internal/middleware"
	"vericred/internal/models"
)

// AuthMe returns the current wallet's auth status and account type
// GET /api/v1/auth/me (protected)
func AuthMe(w http.ResponseWriter, r *http.Request) {
	addr, ok := r.Context().Value(middleware.MetamaskAddressKey).(string)
	if !ok || addr == "" {
		http.Error(w, "metamaskAddress is missing or invalid", http.StatusBadRequest)
		return
	}

	// Fetch account (may be unknown type initially)
	var acc models.Accounts
	_ = db.DB.Where("metamask_address = ?", addr).First(&acc).Error

	// Check presence of profiles
	var user models.Users
	var org models.Organization
	_ = db.DB.Where("metamask_address = ?", addr).First(&user).Error
	_ = db.DB.Where("metamask_address = ?", addr).First(&org).Error

	hasUser := user.ID != 0
	hasOrg := org.ID != 0

	accountType := acc.AccountType
	if accountType == "" || accountType == "unknown" {
		if hasUser {
			accountType = "student"
		} else if hasOrg {
			accountType = "university"
		} else {
			accountType = "unknown"
		}
	}

	_ = json.NewEncoder(w).Encode(map[string]any{
		"address":                addr,
		"account_type":           accountType,
		"has_user_profile":       hasUser,
		"has_university_profile": hasOrg,
	})
}
