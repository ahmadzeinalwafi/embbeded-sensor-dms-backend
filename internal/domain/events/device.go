package event

import (
	"context"
	aggregate "golang_api/internal/domain/aggregates"
)

type DeviceService interface {
	CreateDevice(ctx context.Context, config aggregate.EnteredDeviceInformation) (aggregate.EnteredDeviceInformation, error)
	GetUserDevices(ctx context.Context, user_id string) ([]aggregate.DeviceInformation, error)
}
