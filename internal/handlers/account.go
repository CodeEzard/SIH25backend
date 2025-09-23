package handlers

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"vericred/internal/db"
	"vericred/internal/models"
	"vericred/pkg"
	"vericred/redisdb"

	"gorm.io/gorm"
)


func LoginInMetamask(w http.ResponseWriter, r *http.Request) {
    var body map[string]interface{}

    json.NewDecoder(r.Body).Decode(&body)
    
    metamaskAddress, ok := body["metamask_address"].(string)
    if !ok {
        http.Error(w, "metamask_address field required", http.StatusBadRequest)
        return
    }
    
    signature, ok := body["signature"].(string)
    if !ok {
        http.Error(w, "signature field required", http.StatusBadRequest)
        return
    }
    
    // fmt.Printf("Metamask Address: %s\nSignature: %s", metamaskAddress, signature)

    rdb := redisdb.GetRedisInstance()
    nonce, err := rdb.RedisGetNonce(metamaskAddress)

    if err != nil {
        fmt.Println("Error:", err)
        return
    }
    
    if nonce == "" {
        log.Println("Nonce missing")
        http.Error(w, "nonce missing", http.StatusUnauthorized)
        return
    }

    verified, err := pkg.VerifySignature(metamaskAddress, nonce, signature)
    
    if err != nil || !verified {
        log.Println("Invalid signature")
        http.Error(w, "Invalid signature", http.StatusUnauthorized)
        return
    }

    acc := models.Accounts{
        MetamaskAddress: metamaskAddress,
        AccountType:     "unknown", // defer role until profile creation
    }

    var existingAcc models.Accounts
    result := db.DB.First(&existingAcc, "metamask_address = ?", metamaskAddress)    
    
    if result.Error == gorm.ErrRecordNotFound {
        createResult := db.DB.Create(&acc)
        if createResult.Error != nil {
            http.Error(w, "Failed to create account", http.StatusInternalServerError)
            return
        }

    } else if result.Error != nil {
        http.Error(w, "Database error", http.StatusInternalServerError)
        return
    }

    tokenStr, err := pkg.CreateToken(acc.MetamaskAddress)
    log.Println("tokenStr: ", tokenStr)
    if err != nil {
        http.Error(w, "Failed to create token", http.StatusInternalServerError)
        return
    } 
    
    w.WriteHeader(http.StatusOK)
    fmt.Fprint(w, tokenStr)
}