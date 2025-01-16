package main

import (
	usecase "dms/internal/application/usecases"
	MySQLConnector "dms/internal/infrastructure/database/mysql"
	repository "dms/internal/infrastructure/repositories"
	"dms/internal/interface/http/router"
	"log"
	"net/http"
)

func main() {
	// Establish database connection
	db := MySQLConnector.GetConnection()

	// Initialize repository
	userDeviceRepository := repository.NewUserDeviceRepository(db)

	// Initialize use case
	deviceService := usecase.NewDeviceUseCase(userDeviceRepository)

	// Setup router
	mux := router.SetupRouter(deviceService)

	// Start the server
	log.Println("Server started on :8888")
	if err := http.ListenAndServe(":8888", mux); err != nil {
		log.Fatalf("Server failed: %v", err)
	}
}
