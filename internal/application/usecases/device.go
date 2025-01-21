package usecase

import (
	"context"
	aggregate "dms/internal/domain/aggregates"
	entity "dms/internal/domain/entities"
	event "dms/internal/domain/events"
	repository "dms/internal/domain/repositories"
	tools "dms/tools"
	"fmt"
	"time"
)

type deviceContractImpl struct {
	deviceRepo     repository.DeviceRepository
	userDeviceRepo repository.UserDeviceRepository
}

func NewDeviceUseCase(deviceRepo repository.DeviceRepository, userDeviceRepo repository.UserDeviceRepository) event.DeviceService {
	return &deviceContractImpl{
		deviceRepo:     deviceRepo,
		userDeviceRepo: userDeviceRepo,
	}
}

func (d *deviceContractImpl) CreateDevice(ctx context.Context, device aggregate.EnteredDeviceInformation) (entity.Device, error) {
	device_id := tools.GenerateShortID()
	jwt_token, err := tools.GenerateToken(device_id, device.Name, 12*30*24*time.Hour)
	if err != nil {
		return entity.Device{}, fmt.Errorf("error creating device token: %w", err)
	}
	newDeviceEntity := entity.Device{
		Device_Id:   device_id,
		Name:        device.Name,
		Type:        device.Type,
		Location:    device.Location,
		Token:       jwt_token,
		Status:      device.Status,
		Description: device.Description,
	}

	_, err = d.deviceRepo.Insert(ctx, newDeviceEntity)
	if err != nil {
		return entity.Device{}, fmt.Errorf("error when creating device: %w", err)
	}

	newUserDeviceRelation := entity.UserDevice{
		Device_Id: newDeviceEntity.Device_Id,
		User_Id:   device.Owner,
	}

	_, err = d.userDeviceRepo.Insert(ctx, newUserDeviceRelation)
	if err != nil {
		return entity.Device{}, fmt.Errorf("error when associating device with owner: %w", err)
	}

	return newDeviceEntity, nil
}

func (d *deviceContractImpl) FindInfoByDeviceId(ctx context.Context, device_id string) (aggregate.DeviceInformation, error) {
	if device_id == "" {
		return aggregate.DeviceInformation{}, fmt.Errorf("device ID cannot be empty")
	}

	deviceEntity, err := d.deviceRepo.FindInfoByDeviceId(ctx, device_id)
	if err != nil {
		return aggregate.DeviceInformation{}, fmt.Errorf("error when fetching device information: %w", err)
	}

	return aggregate.DeviceInformation{
		Device_Id:   deviceEntity.Device_Id,
		Name:        deviceEntity.Name,
		Type:        deviceEntity.Type,
		Location:    deviceEntity.Location,
		Token:       deviceEntity.Token,
		Status:      deviceEntity.Status,
		Description: deviceEntity.Description,
		Created_At:  deviceEntity.Created_At,
	}, nil

}

func (d *deviceContractImpl) FindAssosiatedUserByDeviceId(ctx context.Context, device_id string) ([]aggregate.AssosiatedUserInfo, error) {
	if device_id == "" {
		return nil, fmt.Errorf("device ID cannot be empty")
	}

	users, err := d.deviceRepo.FindAssosiatedUserByDeviceId(ctx, device_id)
	if err != nil {
		return nil, fmt.Errorf("error when fetching associated users: %w", err)
	}

	var associatedUsers []aggregate.AssosiatedUserInfo
	for _, user := range users {
		associatedUsers = append(associatedUsers, aggregate.AssosiatedUserInfo{
			User_Id: user.User_Id,
			Email:   user.Email,
		})
	}

	return associatedUsers, nil
}

func (d *deviceContractImpl) DeleteDevice(ctx context.Context, device_id string) error {
	if device_id == "" {
		return fmt.Errorf("device ID cannot be empty")
	}

	err := d.deviceRepo.DeleteByDeviceId(ctx, device_id)
	if err != nil {
		return fmt.Errorf("error when deleting device: %w", err)
	}

	return nil
}
