package main

import (
	"log"
	"net/http"

	"wsu-senior-project/endpoint"
	"wsu-senior-project/service"
	"wsu-senior-project/transport"
)

func main() {
	// Create service
	svc := service.NewHelloService()

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
