package handlers

import (
    "net/http"
    "github.com/go-chi/chi/v5"
    "github.com/skip2/go-qrcode"
    "strings"
)

// GET /api/credential/{id}/qrcode
func GetCredentialQRCode(w http.ResponseWriter, r *http.Request) {
    // Extract credential ID from URL path
    pathParts := strings.Split(r.URL.Path, "/")
    if len(pathParts) < 4 {
        http.Error(w, "Invalid path", http.StatusBadRequest)
        return
    }
    credID := chi.URLParam(r, "id")

    // Data to encode in QR (could be a URL or credential ID)
    data := "https://yourdomain.com/credential/" + credID

    // Generate QR code as PNG
    png, err := qrcode.Encode(data, qrcode.Medium, 256)
    if err != nil {
        http.Error(w, "Failed to generate QR code", http.StatusInternalServerError)
        return
    }

    w.Header().Set("Content-Type", "image/png")
    w.WriteHeader(http.StatusOK)
    w.Write(png)
}