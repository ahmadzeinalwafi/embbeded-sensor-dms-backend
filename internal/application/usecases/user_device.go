package usecase

import (
	"context"
	"fmt"
	"golang_api/internal/application/contracts"
	repository "golang_api/internal/domain/user_device"
	entity "golang_api/internal/domain/user_device/entities"
)

type deviceContractImpl struct {
	repo repository.UserDeviceRepository
}

func NewDeviceUseCase(repo repository.UserDeviceRepository) contracts.DeviceServiceContract {
	return &deviceContractImpl{
		repo: repo,
	}
}

func (u *deviceContractImpl) CreateDevice(ctx context.Context, user_device contracts.EnteredDeviceInformation) (contracts.EnteredDeviceInformation, error) {
	newEntities := entity.UserDevice{
		User_Id:   user_device.DeviceName,
		Device_Id: user_device.Location,
	}
	_, err := u.repo.Insert(ctx, newEntities)

	if err != nil {
		panic(fmt.Sprintf("Error when creating device: %v", err))
	}

	return user_device, nil
}

func (u *deviceContractImpl) GetUserDevices(ctx context.Context, user_id string) ([]contracts.DeviceInformation, error) {
	userDevices, err := u.repo.FindByUserId(ctx, user_id)

	if err != nil {
		panic(fmt.Sprintf("Error when getting the device data by user id: %v", err))
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
