package repository

import (
	"context"
	entity "dms/internal/domain/entities"
)

type HistoricalDeviceRecordsRepository interface {
	CreateMeasurement(ctx context.Context, config *entity.DeviceConfig) error
	WriteData(ctx context.Context, config *entity.HistoricalDeviceRecords) error
	ReadData(ctx context.Context, deviceID string) ([]map[string]interface{}, error)
	DeleteMeasurement(ctx context.Context, deviceID string) error
}
