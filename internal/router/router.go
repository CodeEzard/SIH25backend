package router

import (
	"fmt"
	"net/http"

	"vericred/internal/eth/ipfs"
	"vericred/internal/handlers"
	"vericred/internal/middleware"

	"github.com/go-chi/chi/v5"
)

func RegisterRouter() http.Handler {
	r := chi.NewRouter()
		
	r.Use(middleware.CORSMiddleware)
	r.Use(middleware.LoggingMiddleware)
	// Health-style GET for proxies expecting a GET at /getnonce
	r.Get("/getnonce", handlers.GetNonceHealth)
	// Actual nonce creation (POST)
	r.Post("/getnonce", handlers.GetNonce)
	r.Post("/auth/metamasklogin", handlers.LoginInMetamask)
	r.Get("/universities", handlers.AllOrgs)
	r.Get("/students", handlers.AllUsers)
	r.Post("/credmint", handlers.MintCredentials)
	r.Post("/showuser", handlers.SearchUser)
	r.Post("/usercreds", handlers.ShowSearchedUserCreds)
	r.Get("/transactions", handlers.ShowAllTransactions)
	r.Get("/credential/{id}/qrcode", handlers.GetCredentialQRCode)
	r.Get("/kaithheathcheck", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, "ok")
	})
	// pending request (public create by student via body wallets)
	r.Post("/api/pending/request", handlers.CreatePendingRequest)
	r.Post("/api/specific-university", handlers.SpecificUniversity)

	r.Group(func(r chi.Router) {
		r.Use(middleware.AuthMiddleware)
		r.Post("/api/create/user", handlers.CreateUser)
		r.Post("/api/create/org", handlers.CreateUniversity)
		r.Get("/dashboard", handlers.ShowUser)
		r.Get("/university", handlers.ShowOrg)
		r.Post("/api/uploadtoipfs", ipfs.CreateJSONFileAndStoreToIPFS)
		r.Get("/api/creds", handlers.UserCreds)
		r.Post("/transactionhash", handlers.SetTransactionInfo)
		// pending requests for org
		r.Get("/api/pending/for-org", handlers.ListPendingRequestsForOrg)
		r.Patch("/api/pending/approve", handlers.ApprovePendingRequest)
		// r.Get("/university", handlers.ShowUniversity)
	})
	return r
}
