package event

import (
	"context"
	aggregate "dms/internal/domain/aggregates"
	entity "dms/internal/domain/entities"
)

type DeviceService interface {
	CreateDevice(ctx context.Context, device aggregate.EnteredDeviceInformation) (entity.Device, error)
	SetupDevice(ctx context.Context, device aggregate.FieldsDeviceConfig, device_id string) (entity.DeviceConfig, error)
	FindInfoByDeviceId(ctx context.Context, device_id string) (aggregate.DeviceInformation, error)
	FindAssosiatedUserByDeviceId(ctx context.Context, device_id string) ([]aggregate.AssosiatedUserInfo, error)
	DeleteDevice(ctx context.Context, device_id string) error
}
