package main

import (
	"log"
	"net/http"

	"github.com/rs/cors"

	usecase "dms/internal/application/usecases"
	influxdb "dms/internal/infrastructure/persistance/influxdb"
	mongodb "dms/internal/infrastructure/persistance/mongodb"
	MySQLConnector "dms/internal/infrastructure/persistance/mysql"
	repository "dms/internal/infrastructure/repositories"
	"dms/internal/interface/http/handler"
	"dms/internal/interface/http/router"
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

	corsHandler := cors.New(cors.Options{
		AllowedOrigins:   []string{"*"},                                       
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"}, 
		AllowedHeaders:   []string{"Content-Type", "Authorization"},           
		AllowCredentials: true,
	}).Handler(mux)

	// Start the server
	log.Println("Server started on :8888")
	if err := http.ListenAndServe(":8888", corsHandler); err != nil {
		log.Fatalf("Server failed: %v", err)
	}
}
