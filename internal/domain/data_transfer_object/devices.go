package dto

import "time"

type DeviceInformation struct {
	Device_Id   string
	Name        string
	Type        string
	Location    string
	Token       string
	Status      string
	Description string
	Created_At  time.Time
}

type EnteredDeviceInformation struct {
	Name        string `validate:"required"`
	Type        string `validate:"required"`
	Location    string `validate:"required"`
	Status      string `validate:"required"`
	Description string `validate:"required"`
	Owner       string `validate:"required"`
}

type FieldsDeviceConfig struct {
	Fields map[string]interface{} `validate:"required"`
}

type AssosiatedUserInfo struct {
	User_Id string
	Email   string
}
