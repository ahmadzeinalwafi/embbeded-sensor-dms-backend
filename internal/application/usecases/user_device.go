package usecase

import (
	"context"
	aggregate "dms/internal/domain/aggregates"
	entity "dms/internal/domain/entities"
	event "dms/internal/domain/events"
	repository "dms/internal/domain/repositories"
	"fmt"
)

type deviceContractImpl struct {
	repo repository.UserDeviceRepository
}

func NewDeviceUseCase(repo repository.UserDeviceRepository) event.DeviceService {
	return &deviceContractImpl{
		repo: repo,
	}
}

func (u *deviceContractImpl) CreateDevice(ctx context.Context, user_device aggregate.EnteredDeviceInformation) (aggregate.EnteredDeviceInformation, error) {
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

func (u *deviceContractImpl) GetUserDevices(ctx context.Context, user_id string) ([]aggregate.DeviceInformation, error) {
	userDevices, err := u.repo.FindByUserId(ctx, user_id)

	if err != nil {
		panic(fmt.Sprintf("Error when getting the device data by user id: %v", err))
	}

	var deviceInfos []aggregate.DeviceInformation

	for _, userDevice := range userDevices {
		deviceInfo := aggregate.DeviceInformation{
			DeviceId:   userDevice.Device_Id,
			Created_At: userDevice.Created_At,
			User_Id:    userDevice.User_Id,
		}
		deviceInfos = append(deviceInfos, deviceInfo)
	}

	return deviceInfos, nil
}
