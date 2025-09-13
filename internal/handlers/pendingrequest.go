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

type createPendingPayload struct {
	StudentWallet    string `json:"student_wallet"`
	UniversityWallet string `json:"university_wallet"`
}

// POST /api/pending/request
// Body: { student_wallet, university_wallet }
func CreatePendingRequest(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	fmt.Println("Inside pending request")

	var body createPendingPayload
	if err := json.NewDecoder(r.Body).Decode(&body); err != nil {
		fmt.Println("Invalid JSON in request body")
		http.Error(w, "invalid json", http.StatusBadRequest)
		return
	}
	if body.StudentWallet == "" || body.UniversityWallet == "" {
		fmt.Println("Missing student_wallet or university_wallet in request body")
		http.Error(w, "missing student_wallet or university_wallet", http.StatusBadRequest)
		return
	}

	// find user by student wallet
	var user models.Users
	if err := db.DB.Where("metamask_address = ?", body.StudentWallet).First(&user).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			fmt.Println("Student not registered")
			http.Error(w, "student not registered", http.StatusBadRequest)
			return
		}
		fmt.Println("Database error finding student")
		log.Printf("db error finding student: %v", err)
		http.Error(w, "database error", http.StatusInternalServerError)
		return
	}

	// find organization by university wallet
	var org models.Organization
	if err := db.DB.Where("metamask_address = ?", body.UniversityWallet).First(&org).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			fmt.Println("University not registered")
			http.Error(w, "university not registered", http.StatusBadRequest)
			return
		}
		fmt.Println("Database error finding organization")
		log.Printf("db error finding organization: %v", err)
		http.Error(w, "database error", http.StatusInternalServerError)
		return
	}

	// prevent duplicate pending requests
	var existing models.PendingRequest
	if err := db.DB.Where("requester_id = ? AND organization_id = ? AND is_approved = ?", user.ID, org.ID, false).First(&existing).Error; err == nil {
		fmt.Println("Duplicate pending request found")
		w.WriteHeader(http.StatusConflict)
		json.NewEncoder(w).Encode(existing)
		return
	} else if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
		fmt.Println("Database error checking duplicates")
		log.Printf("db error checking duplicates: %v", err)
		http.Error(w, "database error", http.StatusInternalServerError)
		return
	}

	pending := models.PendingRequest{
		RequesterID:    user.ID,
		OrganizationID: org.ID,
		IsApproved:     false,
	}
	if err := db.DB.Create(&pending).Error; err != nil {
		fmt.Println("Failed to create pending request")
		log.Printf("failed to create pending request: %v", err)
		http.Error(w, "failed to create request", http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(pending)
}

// GET /api/pending/for-org
// Auth required: uses org address from context to list its pending requests
func ListPendingRequestsForOrg(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	fmt.Println("Inside all pending req")
	orgAddr, ok := r.Context().Value(middleware.MetamaskAddressKey).(string)
	if !ok || orgAddr == "" {
		fmt.Println("Unauthorized access: missing or invalid org address in context")
		http.Error(w, "unauthorized", http.StatusUnauthorized)
		return
	}
	var org models.Organization
	if err := db.DB.Where("metamask_address = ?", orgAddr).First(&org).Error; err != nil {
		fmt.Println("Organization not found")
		http.Error(w, "organization not found", http.StatusUnauthorized)
		return
	}

	var reqs []models.PendingRequest
	if err := db.DB.Where("organization_id = ? AND is_approved = ?", org.ID, false).
		Preload("Requester", func(db *gorm.DB) *gorm.DB {
			if db == nil {
				fmt.Println("Failed to preload Requester: DB instance is nil")
				log.Println("failed to preload Requester: DB instance is nil")
				http.Error(w, "database error", http.StatusInternalServerError)
				return nil
			}
			return db
		}).
		Preload("Organization", func(db *gorm.DB) *gorm.DB {
			if db == nil {
				fmt.Println("Failed to preload Organization: DB instance is nil")
				log.Println("failed to preload Organization: DB instance is nil")
				http.Error(w, "database error", http.StatusInternalServerError)
				return nil
			}
			return db
		}).
		Order("created_at DESC").
		Find(&reqs).Error; err != nil {
		fmt.Println("Failed to list pending requests")
		log.Printf("failed to list pending: %v", err)
		http.Error(w, "database error", http.StatusInternalServerError)
		return
	}
	json.NewEncoder(w).Encode(reqs)
}

type approvePendingPayload struct {
	StudentWallet string `json:"student_wallet"`
}

// PATCH /api/pending/approve
// Body: { student_wallet }
// Marks pending request as approved for the org in context and the given student
func ApprovePendingRequest(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	orgAddr, ok := r.Context().Value(middleware.MetamaskAddressKey).(string)
	if !ok || orgAddr == "" {
		fmt.Println("Unauthorized access: missing or invalid org address in context")
		http.Error(w, "unauthorized", http.StatusUnauthorized)
		return
	}
	var body approvePendingPayload
	if err := json.NewDecoder(r.Body).Decode(&body); err != nil {
		fmt.Println("Invalid JSON in request body")
		http.Error(w, "invalid json", http.StatusBadRequest)
		return
	}
	if body.StudentWallet == "" {
		fmt.Println("Missing student_wallet in request body")
		http.Error(w, "missing student_wallet", http.StatusBadRequest)
		return
	}

	var org models.Organization
	if err := db.DB.Where("metamask_address = ?", orgAddr).First(&org).Error; err != nil {
		fmt.Println("Organization not found")
		http.Error(w, "organization not found", http.StatusUnauthorized)
		return
	}
	var user models.Users
	if err := db.DB.Where("metamask_address = ?", body.StudentWallet).First(&user).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			fmt.Println("Student not registered")
			http.Error(w, "student not registered", http.StatusBadRequest)
			return
		}
		fmt.Println("Database error finding student")
		http.Error(w, "database error", http.StatusInternalServerError)
		return
	}

	// update all matching pending requests to approved
	res := db.DB.Model(&models.PendingRequest{}).
		Where("requester_id = ? AND organization_id = ? AND is_approved = ?", user.ID, org.ID, false).
		Update("is_approved", true)
	if res.Error != nil {
		fmt.Println("Failed to approve pending request")
		log.Printf("failed to approve pending: %v", res.Error)
		http.Error(w, "failed to update", http.StatusInternalServerError)
		return
	}
	if res.RowsAffected == 0 {
		fmt.Println("No pending request found to approve")
		http.Error(w, "no pending request found", http.StatusNotFound)
		return
	}

	json.NewEncoder(w).Encode(map[string]any{"updated": res.RowsAffected})
}
