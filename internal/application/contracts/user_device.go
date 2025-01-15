package contracts

import (
	"context"
	"time"
)

type EnteredDeviceInformation struct {
	DeviceName  string
	Description string
	Location    string
}

type DeviceInformation struct {
	DeviceId   string
	Created_At time.Time
	User_Id    string
}

type DeviceServiceContract interface {
	CreateDevice(ctx context.Context, config EnteredDeviceInformation) (EnteredDeviceInformation, error)
	GetUserDevices(ctx context.Context, user_id string) ([]DeviceInformation, error)
}
