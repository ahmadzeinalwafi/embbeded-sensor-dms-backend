package main

import (
	usecase "dms/internal/application/usecases"
	MySQLConnector "dms/internal/infrastructure/database/mysql"
	repository "dms/internal/infrastructure/repositories"
	"dms/internal/interface/http/handler"
	"dms/internal/interface/http/router"
	"log"
	"net/http"
)

func main() {
	// Establish database connection
	db := MySQLConnector.GetConnection()

	// Initialize repositories
	userDeviceRepository := repository.NewUserDeviceRepository(db)
	userRepository := repository.NewUserRepository(db)
	deviceRepository := repository.NewDeviceRepository(db)

	// Initialize use cases
	deviceService := usecase.NewDeviceUseCase(deviceRepository, userDeviceRepository)
	userService := usecase.NewUserUseCase(userRepository)

	// Initialize handlers
	deviceHandler := handler.NewDeviceHandler(deviceService)
	userHandler := handler.NewUserHandler(userService)

	// Setup base router
	mux := router.NewRouter()

	// Add route groups
	router.AddDeviceRoutes(mux, deviceHandler)
	router.AddUserRoutes(mux, userHandler)

	// Start the server
	log.Println("Server started on :8888")
	if err := http.ListenAndServe(":8888", mux); err != nil {
		log.Fatalf("Server failed: %v", err)
	}
}
