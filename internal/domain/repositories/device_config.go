package repository

import (
	"context"
	entity "dms/internal/domain/entities"
)

type DeviceConfigRepository interface {
	Insert(ctx context.Context, sensor *entity.DeviceConfig) error
	FindByDeviceId(ctx context.Context, sensorID string) (*entity.DeviceConfig, error)
	DeleteByDeviceId(ctx context.Context, deviceID string) error
}
