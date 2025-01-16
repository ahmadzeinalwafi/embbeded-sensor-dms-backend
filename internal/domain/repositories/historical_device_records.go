package repository

import (
	"context"
	entity "golang_api/internal/domain/entities"
)

type HistoricalDeviceRecordsRepository interface {
	CreateMeasurement(ctx context.Context, config *entity.DeviceConfigRecords) error
	WriteData(ctx context.Context, config *entity.DeviceConfigRecords) error
	ReadData(ctx context.Context, deviceID string) ([]map[string]interface{}, error)
	DeleteMeasurement(ctx context.Context, deviceID string) error
}