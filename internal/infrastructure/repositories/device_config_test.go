package repositories

import (
	"context"
	"fmt"
	"testing"

	entity "golang_api/internal/domain/device_config/entities"
	"golang_api/internal/infrastructure/database/mongodb"
)

func TestSensorInsert(t *testing.T) {
	client := mongodb.MongoDBConnector()
	defer client.Disconnect(context.Background())

	sensorRepo := NewDeviceConfigRepository(client, "sensor_data", "sensors")

	device_config := &entity.DeviceConfig{
		Device_Id: "12345",
		Fields: map[string]interface{}{
			"temperature_celcius": "float32",
			"humidity":            "int8",
		},
	}

	// Insert sensor data into MongoDB
	err := sensorRepo.Insert(context.Background(), device_config)
	if err != nil {
		t.Fatalf("Error inserting device configuration: %v", err)
	}
	fmt.Printf("Inserted device configuration: %+v\n", device_config)
}

func TestFindBySensorID(t *testing.T) {
	client := mongodb.MongoDBConnector()
	defer client.Disconnect(context.Background())

	sensorRepo := NewDeviceConfigRepository(client, "sensor_data", "sensors")

	found_device_config, err := sensorRepo.FindByDeviceId(context.Background(), "12345")
	if err != nil {
		t.Fatalf("Error finding sensor: %v", err)
	}
	if found_device_config == nil {
		t.Fatalf("Device configuration not found")
	}

	fmt.Printf("Found device configuration: %+v\n", found_device_config)
}

func TestDeleteByDeviceId(t *testing.T) {
	client := mongodb.MongoDBConnector()
	defer client.Disconnect(context.Background())

	repo := NewDeviceConfigRepository(client, "sensor_data", "sensors")

	err := repo.DeleteByDeviceId(context.Background(), "12345")
	if err != nil {
		t.Fatalf("Error deleting device configuration by device id: %v", err)
	}
	fmt.Println("Successfully deleted device configuration by device id 12345")

	deletedSensor, err := repo.FindByDeviceId(context.Background(), "12345")
	if err != nil {
		t.Fatalf("Error finding device after deletion: %v", err)
	}
	if deletedSensor != nil {
		t.Fatalf("Device configuration was not deleted")
	}
}
