package handlers

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"os"
	"strconv"
	"strings"
	"time"

	"github.com/go-chi/chi/v5"
	"github.com/golang-jwt/jwt/v5"

	"vericred/internal/db"
	"vericred/internal/middleware"
	"vericred/internal/models"
)

type shareClaims struct {
	CredentialID string `json:"credential_id"`
	jwt.RegisteredClaims
}

type generateShareLinkReq struct {
	CredentialID   string `json:"credential_id"`
	ExpiresInHours int    `json:"expires_in_hours"`
}

type generateShareLinkResp struct {
	ShareableURL string `json:"shareable_url"`
}

func getShareSecret() ([]byte, error) {
	if s := os.Getenv("SHARE_TOKEN_SECRET"); s != "" {
		return []byte(s), nil
	}
	if s := os.Getenv("JWT_SECRET"); s != "" {
		return []byte(s), nil
	}
	return nil, errors.New("missing SHARE_TOKEN_SECRET/JWT_SECRET")
}

// POST /api/v1/credentials/generate-share-link (protected)
func GenerateShareLink(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	addr, ok := r.Context().Value(middleware.MetamaskAddressKey).(string)
	if !ok || addr == "" {
		http.Error(w, "unauthorized", http.StatusUnauthorized)
		return
	}

	// Be liberal in what we accept from the frontend
	var payload map[string]any
	if err := json.NewDecoder(r.Body).Decode(&payload); err != nil {
		http.Error(w, "invalid json", http.StatusBadRequest)
		return
	}

	credID := ""
	if v, ok := payload["credential_id"].(string); ok {
		credID = strings.TrimSpace(v)
	} else if v, ok := payload["credentialId"].(string); ok { // optional camelCase fallback
		credID = strings.TrimSpace(v)
	}
	if credID == "" {
		http.Error(w, "credential_id is required", http.StatusBadRequest)
		return
	}

	// expires_in_hours may come as number or string, and snake_case or camelCase
	parseHours := func(x any) (int, bool) {
		switch t := x.(type) {
		case float64:
			return int(t), true
		case json.Number:
			if i, err := strconv.Atoi(t.String()); err == nil {
				return i, true
			}
		case string:
			if i, err := strconv.Atoi(strings.TrimSpace(t)); err == nil {
				return i, true
			}
		}
		return 0, false
	}
	expires := 0
	if v, ok := payload["expires_in_hours"]; ok {
		if i, ok2 := parseHours(v); ok2 {
			expires = i
		}
	} else if v, ok := payload["expiresInHours"]; ok { // camelCase fallback
		if i, ok2 := parseHours(v); ok2 {
			expires = i
		}
	} else if v, ok := payload["duration"]; ok { // alias used by frontend
		if i, ok2 := parseHours(v); ok2 {
			expires = i
		}
	}
	// Enforce 1..168 hours to avoid immediately-expired tokens
	if expires < 1 || expires > 168 {
		http.Error(w, "expires_in_hours must be between 1 and 168", http.StatusBadRequest)
		return
	}

	// Verify ownership: credential must belong to this student wallet
	var cred models.Credential
	if err := db.DB.Where("id = ?", credID).First(&cred).Error; err != nil {
		http.Error(w, "credential not found", http.StatusNotFound)
		return
	}
	if cred.StudentWallet == "" || !equalCaseInsensitive(cred.StudentWallet, addr) {
		http.Error(w, "forbidden: not owner of credential", http.StatusForbidden)
		return
	}

	secret, err := getShareSecret()
	if err != nil {
		http.Error(w, "server misconfigured", http.StatusInternalServerError)
		return
	}

	exp := time.Now().Add(time.Duration(expires) * time.Hour)
	claims := shareClaims{
		CredentialID: credID,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(exp),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
		},
	}
	tok := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	signed, err := tok.SignedString(secret)
	if err != nil {
		http.Error(w, "failed to sign share token", http.StatusInternalServerError)
		return
	}

	base := os.Getenv("FRONTEND_BASE_URL")
	if base == "" {
		base = "http://localhost:3000"
	}
	url := fmt.Sprintf("%s/verify/%s?token=%s", trimRightSlash(base), credID, signed)
	_ = json.NewEncoder(w).Encode(generateShareLinkResp{ShareableURL: url})
}

// GET /api/v1/credential-info/{id}?token=...
func GetCredentialInfo(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	id := chi.URLParam(r, "id")
	if id == "" {
		http.Error(w, "missing id", http.StatusBadRequest)
		return
	}
	tokenStr := r.URL.Query().Get("token")
	if tokenStr == "" {
		http.Error(w, "This verification link is invalid or has expired.", http.StatusUnauthorized)
		return
	}

	secret, err := getShareSecret()
	if err != nil {
		http.Error(w, "server misconfigured", http.StatusInternalServerError)
		return
	}

	parsed, err := jwt.ParseWithClaims(tokenStr, &shareClaims{}, func(t *jwt.Token) (interface{}, error) {
		if _, ok := t.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, errors.New("unexpected signing method")
		}
		return secret, nil
	})
	fmt.Println("Parsed: ", parsed)
	if err != nil || !parsed.Valid {
		http.Error(w, "This verification link is invalid or has expired.", http.StatusUnauthorized)
		return
	}
	claims, ok := parsed.Claims.(*shareClaims)
	fmt.Println("claims: ", claims)
	if !ok || claims.CredentialID == "" || claims.ExpiresAt == nil || time.Now().After(claims.ExpiresAt.Time) {
		http.Error(w, "This verification link is invalid or has expired.", http.StatusUnauthorized)
		return
	}
	if claims.CredentialID != id {
		http.Error(w, "forbidden: id mismatch", http.StatusForbidden)
		return
	}

	var cred models.Credential
	if err := db.DB.Where("id = ?", id).First(&cred).Error; err != nil {
		http.Error(w, "credential not found", http.StatusNotFound)
		return
	}

	// Optionally fetch IPFS document (best-effort)
	var ipfs any
	if cred.IPFSLink != "" {
		client := &http.Client{Timeout: 10 * time.Second}
		if resp, e := client.Get(cred.IPFSLink); e == nil && resp != nil && resp.Body != nil {
			defer resp.Body.Close()
			_ = json.NewDecoder(resp.Body).Decode(&ipfs)
		}
	}

	_ = json.NewEncoder(w).Encode(map[string]any{
		"credential":     cred,
		"ipfs":           ipfs,
		"valid_until":    claims.ExpiresAt.Time,
	})
}

func equalCaseInsensitive(a, b string) bool {
	return strings.EqualFold(strings.TrimSpace(a), strings.TrimSpace(b))
}

func trimRightSlash(s string) string {
	for len(s) > 0 && s[len(s)-1] == '/' {
		s = s[:len(s)-1]
	}
	return s
}
