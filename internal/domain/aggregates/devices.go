package aggregate

import "time"

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