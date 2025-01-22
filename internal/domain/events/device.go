package event

import (
	"context"
	dto "dms/internal/domain/data_transfer_object"
	entity "dms/internal/domain/entities"
)

type DeviceService interface {
	CreateDevice(ctx context.Context, device dto.EnteredDeviceInformation) (entity.Device, error)
	SetupDevice(ctx context.Context, device dto.FieldsDeviceConfig, device_id string) (entity.DeviceConfig, error)
	CreateRecordsDevice(ctx context.Context, deviceRecords dto.FieldsDeviceRecords, device_id string) (entity.HistoricalDeviceRecords, error)
	ReadRecordsDevice(ctx context.Context, device_id string) ([]map[string]interface{}, error)
	FindInfoByDeviceId(ctx context.Context, device_id string) (dto.DeviceInformation, error)
	FindAssosiatedUserByDeviceId(ctx context.Context, device_id string) ([]dto.AssosiatedUserInfo, error)
	DeleteDevice(ctx context.Context, device_id string) error
}
