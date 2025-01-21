package repository

import (
	"context"
	entity "dms/internal/domain/entities"
)

type DeviceRepository interface {
	Insert(ctx context.Context, device entity.Device) (entity.Device, error)
	FindInfoByDeviceId(ctx context.Context, device_id string) (entity.Device, error)
	FindAssosiatedUserByDeviceId(ctx context.Context, device_id string) ([]entity.User, error)
	DeleteByDeviceId(ctx context.Context, device_id string) error
}
