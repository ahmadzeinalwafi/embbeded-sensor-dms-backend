package repository

import (
	"context"
	"testing"

	entity "dms/internal/domain/entities"
	"dms/internal/infrastructure/persistance/influxdb"
)

func TestCreateMeasurementHistoricalDeviceRecordsRepository(t *testing.T) {
	client, err := influxdb.InfluxDBConnector()
	if err != nil {
		t.Fatalf("Failed to connect to influxdb: %v", err)
	}

	defer client.Close()

	repo := NewHistoricalDeviceRecordsRepository(client, "dms_bucket", "dms_org")
	deviceConfig := entity.DeviceConfig{
		Device_Id: "device_123",
		Fields: map[string]interface{}{
			"temperature": "float64",
			"humidity":    "int8",
		},
	}

	err = repo.CreateMeasurement(context.Background(), deviceConfig)
	if err != nil {
		t.Fatalf("Failed to create measurement: %v", err)
	}
}

func TestWriteDataHistoricalDeviceRecordsRepository(t *testing.T) {
	client, err := influxdb.InfluxDBConnector()
	if err != nil {
		t.Fatalf("Failed to connect to influxdb: %v", err)
	}

	defer client.Close()

	repo := NewHistoricalDeviceRecordsRepository(client, "dms_bucket", "dms_org")

	deviceConfig := entity.HistoricalDeviceRecords{
		Device_Id: "device_123",
		Fields: map[string]interface{}{
			"temperature": float32(25.5),
			"humidity":    int8(80),
		},
	}

	deviceConfig.Fields["temperature"] = float64(26.5)
	deviceConfig.Fields["humidity"] = int64(85)
	err = repo.WriteData(context.Background(), deviceConfig)
	if err != nil {
		t.Fatalf("Failed to write data: %v", err)
	}
}

func TestReadDataHistoricalDeviceRecordsRepository(t *testing.T) {
	client, err := influxdb.InfluxDBConnector()
	if err != nil {
		t.Fatalf("Failed to connect to influxdb: %v", err)
	}

	defer client.Close()

	repo := NewHistoricalDeviceRecordsRepository(client, "dms_bucket", "dms_org")

	deviceConfig := &entity.HistoricalDeviceRecords{
		Device_Id: "device_123",
		Fields: map[string]interface{}{
			"temperature": float32(25.5),
			"humidity":    int8(80),
		},
	}

	data, err := repo.ReadData(context.Background(), deviceConfig.Device_Id)
	if err != nil {
		t.Fatalf("Failed to read data: %v", err)
	}
	t.Logf("Read data: %+v", data)
}

func TestDeleteHistoricalDeviceRecordsRepository(t *testing.T) {
	client, err := influxdb.InfluxDBConnector()
	if err != nil {
		t.Fatalf("Failed to connect to influxdb: %v", err)
	}

	defer client.Close()

	repo := NewHistoricalDeviceRecordsRepository(client, "dms_bucket", "dms_org")

	deviceConfig := &entity.HistoricalDeviceRecords{
		Device_Id: "device_123",
		Fields: map[string]interface{}{
			"temperature": float32(25.5),
			"humidity":    int8(80),
		},
	}

	err = repo.DeleteMeasurement(context.Background(), deviceConfig.Device_Id)
	if err != nil {
		t.Fatalf("Failed to delete measurement: %v", err)
	}
}
