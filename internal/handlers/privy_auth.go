package handlers

import (
	"crypto/ecdsa"
	"crypto/elliptic"
	"crypto/rsa"
	"encoding/base64"
	"encoding/json"
	"errors"
	"fmt"
	"math/big"
	"net/http"
	"os"
	"strconv"
	"strings"
	"sync"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"gorm.io/gorm"

	"vericred/internal/db"
	"vericred/internal/models"
	"vericred/pkg"
)

// JWKS caching for Privy
var (
	jwksCache struct {
		mu        sync.RWMutex
		keys      map[string]any // kid -> public key (RSA/ECDSA)
		fetchedAt time.Time
		url       string
	}
)

type jwks struct { Keys []jwk `json:"keys"` }

type jwk struct {
	Kty string `json:"kty"`
	Kid string `json:"kid"`
	Alg string `json:"alg"`
	Use string `json:"use"`
	// RSA
	N   string `json:"n"`
	E   string `json:"e"`
	// EC
	Crv string `json:"crv"`
	X   string `json:"x"`
	Y   string `json:"y"`
}

func getPrivyJWKSURL() (string, error) {
	u := os.Getenv("PRIVY_JWKS_URL")
	if u == "" {
		return "", errors.New("PRIVY_JWKS_URL not set")
	}
	return u, nil
}

func fetchJWKS(url string) (map[string]any, error) {
	client := &http.Client{Timeout: 10 * time.Second}
	resp, err := client.Get(url)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	if resp.StatusCode != 200 {
		return nil, fmt.Errorf("jwks fetch failed: %s", resp.Status)
	}
	var doc jwks
	if err := json.NewDecoder(resp.Body).Decode(&doc); err != nil {
		return nil, err
	}
	out := make(map[string]any)
	for _, k := range doc.Keys {
		switch strings.ToUpper(k.Kty) {
		case "RSA":
			if k.N == "" || k.E == "" || k.Kid == "" { continue }
			pub, err := rsaFromJWK(k.N, k.E)
			if err != nil { continue }
			out[k.Kid] = pub
		case "EC":
			// Expect ES256 (P-256)
			if k.Crv == "" || k.X == "" || k.Y == "" || k.Kid == "" { continue }
			pub, err := ecdsaFromJWK(k.Crv, k.X, k.Y)
			if err != nil { continue }
			out[k.Kid] = pub
		}
	}
	if len(out) == 0 { return nil, errors.New("empty jwks") }
	return out, nil
}

func rsaFromJWK(nB64, eB64 string) (*rsa.PublicKey, error) {
	nBytes, err := base64.RawURLEncoding.DecodeString(nB64)
	if err != nil { return nil, err }
	eBytes, err := base64.RawURLEncoding.DecodeString(eB64)
	if err != nil { return nil, err }
	// Convert eBytes big-endian to int
	eInt := 0
	for _, b := range eBytes { eInt = eInt<<8 | int(b) }
	if eInt == 0 { eInt = 65537 }
	return &rsa.PublicKey{N: new(big.Int).SetBytes(nBytes), E: eInt}, nil
}

func ecdsaFromJWK(crv, xB64, yB64 string) (*ecdsa.PublicKey, error) {
	var curve elliptic.Curve
	switch crv {
	case "P-256":
		curve = elliptic.P256()
	default:
		return nil, fmt.Errorf("unsupported curve: %s", crv)
	}
	xBytes, err := base64.RawURLEncoding.DecodeString(xB64)
	if err != nil { return nil, err }
	yBytes, err := base64.RawURLEncoding.DecodeString(yB64)
	if err != nil { return nil, err }
	X := new(big.Int).SetBytes(xBytes)
	Y := new(big.Int).SetBytes(yBytes)
	return &ecdsa.PublicKey{Curve: curve, X: X, Y: Y}, nil
}

func getPrivyKeyForKid(kid string) (any, error) {
	jwksCache.mu.RLock()
	if jwksCache.keys != nil && time.Since(jwksCache.fetchedAt) < time.Hour {
		if k, ok := jwksCache.keys[kid]; ok {
			jwksCache.mu.RUnlock()
			return k, nil
		}
	}
	jwksCache.mu.RUnlock()

	url, err := getPrivyJWKSURL()
	if err != nil { return nil, err }
	keys, err := fetchJWKS(url)
	if err != nil { return nil, err }
	jwksCache.mu.Lock()
	jwksCache.keys = keys
	jwksCache.fetchedAt = time.Now()
	jwksCache.url = url
	jwksCache.mu.Unlock()
	if k, ok := keys[kid]; ok { return k, nil }
	return nil, errors.New("kid not found in jwks")
}

// PrivyLogin handles POST /api/v1/auth/privy-login
// Body: { "privy_token": "..." }
func PrivyLogin(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var body map[string]any
	if err := json.NewDecoder(r.Body).Decode(&body); err != nil {
		http.Error(w, "invalid json", http.StatusBadRequest)
		return
	}
	tokStr, _ := body["privy_token"].(string)
	if strings.TrimSpace(tokStr) == "" {
		http.Error(w, "privy_token is required", http.StatusBadRequest)
		return
	}

	// Verify JWT with Privy JWKS (support RS256 and ES256)
	parsed, err := jwt.Parse(tokStr, func(t *jwt.Token) (interface{}, error) {
		if t.Method != jwt.SigningMethodRS256 && t.Method != jwt.SigningMethodES256 {
			return nil, errors.New("unexpected signing method (need RS256 or ES256)")
		}
		kid, _ := t.Header["kid"].(string)
		if kid == "" { return nil, errors.New("missing kid") }
		return getPrivyKeyForKid(kid)
	})
	if err != nil || !parsed.Valid {
		http.Error(w, "invalid privy token", http.StatusUnauthorized)
		return
	}

	claims, ok := parsed.Claims.(jwt.MapClaims)
	if !ok { http.Error(w, "invalid claims", http.StatusUnauthorized); return }

	// exp check
	now := time.Now().Unix()
	if exp, ok := getUnixFromClaim(claims["exp"]); !ok || now >= exp {
		http.Error(w, "token expired", http.StatusUnauthorized)
		return
	}
	// Optional issuer/audience checks via env
	if iss := os.Getenv("PRIVY_ISSUER"); iss != "" {
		if v, _ := claims["iss"].(string); v != iss { http.Error(w, "issuer mismatch", http.StatusUnauthorized); return }
	}
	if aud := os.Getenv("PRIVY_AUDIENCE"); aud != "" {
		if !audienceContains(claims["aud"], aud) { http.Error(w, "audience mismatch", http.StatusUnauthorized); return }
	}

	addr := extractAddressFromPrivyClaims(claims)
	if addr == "" {
		// Fallback: accept wallet_address provided by client if included (optional)
		if v, ok := body["wallet_address"].(string); ok { addr = v }
	}
	addr = strings.ToLower(strings.TrimSpace(addr))
	if !strings.HasPrefix(addr, "0x") || len(addr) != 42 {
		http.Error(w, "wallet address not found in token", http.StatusBadRequest)
		return
	}

	// Provision account if needed (same as MetaMask flow)
	acc := models.Accounts{ MetamaskAddress: addr, AccountType: "user" }
	var existing models.Accounts
	res := db.DB.First(&existing, "metamask_address = ?", addr)
	if errors.Is(res.Error, gorm.ErrRecordNotFound) {
		if err := db.DB.Create(&acc).Error; err != nil {
			http.Error(w, "failed to create account", http.StatusInternalServerError)
			return
		}
	} else if res.Error != nil {
		http.Error(w, "database error", http.StatusInternalServerError)
		return
	}

	// Issue our session JWT (same as MetaMask)
	signed, err := pkg.CreateToken(addr)
	if err != nil {
		http.Error(w, "failed to create token", http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusOK)
	fmt.Fprint(w, signed)
}

func extractAddressFromPrivyClaims(mc jwt.MapClaims) string {
	// Try common fields
	if v, ok := mc["wallet_address"].(string); ok { return v }
	if v, ok := mc["address"].(string); ok { return v }
	if v, ok := mc["eth_address"].(string); ok { return v }
	// Try nested user.wallet.address
	if u, ok := mc["user"].(map[string]any); ok {
		if v, ok := u["wallet_address"].(string); ok { return v }
		if w, ok := u["wallet"].(map[string]any); ok {
			if v, ok := w["address"].(string); ok { return v }
		}
		if ws, ok := u["wallets"].([]any); ok && len(ws) > 0 {
			if w0, ok := ws[0].(map[string]any); ok {
				if v, ok := w0["address"].(string); ok { return v }
			}
		}
	}
	return ""
}

func getUnixFromClaim(v any) (int64, bool) {
	switch t := v.(type) {
	case float64:
		return int64(t), true
	case json.Number:
		if i, err := t.Int64(); err == nil { return i, true }
	case string:
		if i, err := time.ParseDuration(t); err == nil { return time.Now().Add(i).Unix(), true }
		if i, err := strconv.ParseInt(t, 10, 64); err == nil { return i, true }
	}
	return 0, false
}

func audienceContains(v any, target string) bool {
	if s, ok := v.(string); ok { return s == target }
	if arr, ok := v.([]any); ok {
		for _, it := range arr {
			if s, ok := it.(string); ok && s == target { return true }
		}
	}
	return false
}
