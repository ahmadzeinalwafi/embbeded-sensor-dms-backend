package event

import (
	"context"
	aggregate "dms/internal/domain/aggregates"
)

type DeviceService interface {
	CreateDevice(ctx context.Context, config aggregate.EnteredDeviceInformation) (aggregate.EnteredDeviceInformation, error)
	GetUserDevices(ctx context.Context, user_id string) ([]aggregate.DeviceInformation, error)
}
