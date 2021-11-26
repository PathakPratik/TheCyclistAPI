package main

import (
	"fmt"
	"log"
	"net/http"

	// "github.com/TheCyclistGoServer/API"
	"github.com/TheCyclistGoServer/AwsDynamoDb"
)

func main() {
	// Init ServeMux Router
	mux := router()

	// Connect to DynamoDb
	AwsDynamoDb.InitDynamoDb()

	// Called Once to Store Data of Events in DynamoDB
	// API.GetEvents()

	// Start Server
	fmt.Printf("Starting server at port 8080\n")
	// port := os.Getenv("PORT")
	// log.Fatal(http.ListenAndServe(":"+port, mux))
	log.Fatal(http.ListenAndServe(":8080", mux))
}
