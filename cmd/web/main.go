package main

import (
	usecase "dms/internal/application/usecases"
	influxdb "dms/internal/infrastructure/database/influxdb"
	mongodb "dms/internal/infrastructure/database/mongodb"
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
	mongo := mongodb.MongoDBConnector()
	influx, err := influxdb.InfluxDBConnector()
	if err != nil {
		panic(err)
	}

	// Initialize repositories
	userDeviceRepository := repository.NewUserDeviceRepository(db)
	userRepository := repository.NewUserRepository(db)
	deviceRepository := repository.NewDeviceRepository(db)
	deviceConfigRepository := repository.NewDeviceConfigRepository(mongo, "sensors", "configuration")
	historicalDeviceRecordsRepository := repository.NewHistoricalDeviceRecordsRepository(influx, "dms_bucket", "dms_org")

	// Initialize use cases
	deviceService := usecase.NewDeviceUseCase(deviceRepository, userDeviceRepository, deviceConfigRepository, historicalDeviceRecordsRepository)
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
