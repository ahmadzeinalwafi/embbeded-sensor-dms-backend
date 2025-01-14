package repository

import (
	"context"
	entity "golang_api/internal/domain/device_config/entities"
)

type DeviceConfigRepository interface {
	Insert(ctx context.Context, sensor *entity.DeviceConfig) error
	FindByDeviceId(ctx context.Context, sensorID string) (*entity.DeviceConfig, error)
	DeleteByDeviceId(ctx context.Context, deviceID string) error
}
