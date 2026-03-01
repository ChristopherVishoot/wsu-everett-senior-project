package main

import (
	"log"
	"net/http"

	"wsu-senior-project/database"
	"wsu-senior-project/endpoint"
	"wsu-senior-project/service"
	"wsu-senior-project/transport"
)

func main() {
	// Connect to database
	cfg := database.NewConfigFromEnv()
	db, err := database.Connect(cfg)
	if err != nil {
		log.Fatal("Failed to connect to database:", err)
	}
	defer db.Close()
	log.Println("Connected to database")

	// Create service
	svc := service.NewService(db)

	// Create endpoints
	endpoints := endpoint.MakeEndpoints(svc)

	// Create HTTP handler
	handler := transport.NewHTTPHandler(endpoints)

	// Start server
	log.Println("Starting server on :8080")
	if err := http.ListenAndServe(":8080", handler); err != nil {
		log.Fatal(err)
	}
}
