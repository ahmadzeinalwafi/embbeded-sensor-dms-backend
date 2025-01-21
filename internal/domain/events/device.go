package event

import (
	"context"
	aggregate "dms/internal/domain/aggregates"
	entity "dms/internal/domain/entities"
)

type DeviceService interface {
	//This will make device that will call repositories 1. DeviceRepository.Insert 2. UserDeviceRepository.Insert
	CreateDevice(ctx context.Context, device aggregate.EnteredDeviceInformation) (entity.Device, error)
	// This will get information of device that will call repositories DeviceRepository.FindInfoByDeviceId
	FindInfoByDeviceId(ctx context.Context, device_id string) (aggregate.DeviceInformation, error)
	// This will get information of assosiated user within device that will call repositories FindAssosiatedUserByDeviceId
	FindAssosiatedUserByDeviceId(ctx context.Context, device_id string) ([]aggregate.AssosiatedUserInfo, error)
	// This will delete the device that will call repositories DeleteByDeviceId
	DeleteDevice(ctx context.Context, device_id string) error
}
