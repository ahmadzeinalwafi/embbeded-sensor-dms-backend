package usecase

import (
	"context"
	"database/sql"
	"fmt"
	"golang_api/internal/application/contracts"
	entity "golang_api/internal/domain/user_device/entities"
	"golang_api/internal/infrastructure/repositories"
)

type deviceContractImpl struct {
	DB *sql.DB
}

func NewDeviceUseCase(db *sql.DB) contracts.DeviceServiceContract {
	return &deviceContractImpl{DB: db}
}

func (repository *deviceContractImpl) CreateDevice(ctx context.Context, user_device contracts.EnteredDeviceInformation) (contracts.EnteredDeviceInformation, error) {
	userDeviceRepository := repositories.NewUserDeviceRepository(repository.DB)
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
	userDeviceRepository := repositories.NewUserDeviceRepository(repository.DB)
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
