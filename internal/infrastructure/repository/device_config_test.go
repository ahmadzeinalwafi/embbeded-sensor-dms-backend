package repository

import (
	"context"
	"fmt"
	"testing"

	entity "golang_api/internal/domain/device_config/entities"
	"golang_api/internal/infrastructure/database/mongodb"
)

// Test for inserting a sensor
func TestSensorInsert(t *testing.T) {
	client := mongodb.MongoDBConnector()
	defer client.Disconnect(context.Background())

	sensorRepo := NewSensorRepository(client, "sensor_data", "sensors")

	sensor := &entity.SensorConfig{
		SensorID: "12345",
		Fields: map[string]interface{}{
			"temperature_celcius": "float32",
			"humidity":            "int8",
		},
	}

	// Insert sensor data into MongoDB
	err := sensorRepo.Insert(context.Background(), sensor)
	if err != nil {
		t.Fatalf("Error inserting sensor: %v", err)
	}
	fmt.Printf("Inserted sensor: %+v\n", sensor)
}

func TestFindBySensorID(t *testing.T) {
	client := mongodb.MongoDBConnector()
	defer client.Disconnect(context.Background())

	sensorRepo := NewSensorRepository(client, "sensor_data", "sensors")

	foundSensor, err := sensorRepo.FindByID(context.Background(), "12345")
	if err != nil {
		t.Fatalf("Error finding sensor: %v", err)
	}
	if foundSensor == nil {
		t.Fatalf("Sensor not found")
	}

	fmt.Printf("Found sensor: %+v\n", foundSensor)
}
