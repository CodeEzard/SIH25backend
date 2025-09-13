package main

import (
	"fmt"
	"log"
	"net/http"

	"vericred/internal/db"
	"vericred/internal/logging"
	"vericred/internal/router"
)

func main() {
	logging.Init()

	db.Init()

	r := router.RegisterRouter()
	fmt.Println("Port :8080 is active....")
	if err := http.ListenAndServe(":8080", r); err != nil {
		log.Fatal(err)
	}
}
