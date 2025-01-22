package main

import (
	"context"
	influxdb "dms/internal/infrastructure/persistance/influxdb"
	mongodb "dms/internal/infrastructure/persistance/mongodb"
	"log"
	"os"
	"os/signal"
	"syscall"

	usecase "dms/internal/application/usecases"
	mqtt "dms/internal/infrastructure/messaging/paho"
	repository "dms/internal/infrastructure/repositories"
	consumer "dms/internal/interface/mqtt/consumer"
)

func main() {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	go func() {
		sigChan := make(chan os.Signal, 1)
		signal.Notify(sigChan, os.Interrupt, syscall.SIGTERM)
		<-sigChan
		log.Println("Termination signal received, shutting down gracefully...")
		cancel()
	}()

	mongo := mongodb.MongoDBConnector()
	influx, err := influxdb.InfluxDBConnector()
	if err != nil {
		panic(err)
	}
	repo := repository.NewHistoricalDeviceRecordsRepository(influx, "dms_bucket", "dms_org")
	deviceConfigRepository := repository.NewDeviceConfigRepository(mongo, "sensors", "configuration")

	// Create use case
	messageProcessor := usecase.NewMessageProcessorUseCase(repo, deviceConfigRepository)

	// Initialize MQTT connection
	mqttConnection, err := mqtt.NewMQTTConnection("localhost:1883")
	if err != nil {
		log.Fatalf("Failed to initialize MQTT connection: %v", err)
	}

	if err != nil {
		log.Fatalf("Failed to initialize MQTT connection: %v", err)
	}
	consumer := consumer.NewMQTTConsumer(mqttConnection.Client, "devices/data", messageProcessor)
	err = consumer.Start()
	if err != nil {
		log.Fatalf("Failed to start MQTT consumer: %v", err)
	}
	// Wait for shutdown signal
	<-ctx.Done()
	log.Println("Application stopped")
}
