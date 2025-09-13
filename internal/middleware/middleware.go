package middleware

import (
	"context"
	"fmt"
	"net/http"
	"strings"
	"vericred/internal/logging"
	"vericred/pkg"
)

type contextKey string

const MetamaskAddressKey contextKey = "metamaskAddress"

func AuthMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		authHeader := r.Header.Get("Authorization")
		if authHeader == "" {
			logging.Logger.Println("Auth header req.")
			http.Error(w, "Authorization header required", http.StatusUnauthorized)
			return
		}

		bearerToken := strings.Split(authHeader, " ")
		if len(bearerToken) != 2 || bearerToken[0] != "Bearer" {
			http.Error(w, "Invalid auth format", http.StatusUnauthorized)
			return
		}

		token := bearerToken[1]
		address, err := pkg.VerifyToken(token)
		if err != nil {
			fmt.Println("Error:", err)
			return
		}
		fmt.Println("Set the address in middleware: ", address)

		ctx := context.WithValue(r.Context(), MetamaskAddressKey, address)		
		next.ServeHTTP(w, r.WithContext(ctx))

	})
}

// func ProtectedHandler(w http.ResponseWriter, r *http.Request) {
// 	w.Header().Set("Content-Type", "application/json")
// 	tokenString := r.Header.Get("Authorization")
// 	if tokenString == "" {
// 		w.WriteHeader(http.StatusUnauthorized)
// 		fmt.Fprintf(w, "Missing auth header.")
// 		return
// 	}
// 	tokenString = tokenString[len("Bearer "):]

//     _, err := pkg.VerifyToken(tokenString)
// 	if err != nil {
// 		w.WriteHeader(http.StatusUnauthorized)
// 		fmt.Fprint(w, "Invalid token")
// 		return
// 	}

// 	fmt.Fprint(w, "welcome fucker.")
// }

func LoggingMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Printf("[%s] %s %s\n", r.Method, r.RemoteAddr, r.URL.Path)
		next.ServeHTTP(w, r)
	})
}

func CORSMiddleware(next http.Handler)  http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

        w.Header().Set("Access-Control-Allow-Origin", "*") // we need to sepcify the frontend shit when in production
        w.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")
        w.Header().Set("Access-Control-Allow-Headers", "Content-Type, Authorization, X-Requested-With")
        w.Header().Set("Access-Control-Expose-Headers", "Authorization")
        w.Header().Set("Access-Control-Allow-Credentials", "true")
        if r.Method == "OPTIONS" {
            w.WriteHeader(http.StatusOK)
            return
        }

        next.ServeHTTP(w, r)
	})


}

