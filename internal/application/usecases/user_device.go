package usecase

import (
	"context"
	"fmt"
	"golang_api/internal/application/contracts"
	entity "golang_api/internal/domain/user_device/entities"
	MySQLConnector "golang_api/internal/infrastructure/database/mysql"
	"golang_api/internal/infrastructure/repositories"
)

type deviceContractImpl struct {
}

func NewDeviceUseCase() contracts.DeviceServiceContract {
	return &deviceContractImpl{}
}

func (repository *deviceContractImpl) CreateDevice(ctx context.Context, user_device contracts.EnteredDeviceInformation) (contracts.EnteredDeviceInformation, error) {
	userDeviceRepository := repositories.NewUserDeviceRepository(MySQLConnector.GetConnection())
	newEntities := entity.UserDevice{
		User_Id:   user_device.DeviceName,
		Device_Id: user_device.Location,
	}

	_, err := userDeviceRepository.Insert(ctx, newEntities)
	if err != nil {
		panic(fmt.Sprintf("Error finding user by ID: %v", err))
	}

	return user_device, nil
}

func (repository *deviceContractImpl) GetUserDevices(ctx context.Context, user_id string) ([]contracts.DeviceInformation, error) {
	userDeviceRepository := repositories.NewUserDeviceRepository(MySQLConnector.GetConnection())
	userDevices, err := userDeviceRepository.FindByUserId(ctx, user_id)

	if err != nil {
		panic(fmt.Sprintf("Error finding user by ID: %v", err))
	}

	var deviceInfos []contracts.DeviceInformation

	for _, userDevice := range userDevices {
		deviceInfo := contracts.DeviceInformation{
			DeviceId:   userDevice.Device_Id,
			Created_At: userDevice.Created_At,
			User_Id:    userDevice.User_Id,
		}
		deviceInfos = append(deviceInfos, deviceInfo)
	}

	return deviceInfos, nil
}
