package main

import (
	usecase "golang_api/internal/application/usecases"
	MySQLConnector "golang_api/internal/infrastructure/database/mysql"
	"golang_api/internal/interface/http/router"
	"log"
	"net/http"
)

func main() {
	// Establish database connection
	db := MySQLConnector.GetConnection()

	// Initialize the use case
	deviceService := usecase.NewDeviceUseCase(db)

	// Setup router
	mux := router.SetupRouter(deviceService)

	// Start the server
	log.Println("Server started on :8888")
	if err := http.ListenAndServe(":8888", mux); err != nil {
		log.Fatalf("Server failed: %v", err)
	}
}
