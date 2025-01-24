package usecase

import (
	"context"
	dto "dms/internal/domain/data_transfer_object"
	entity "dms/internal/domain/entities"
	event "dms/internal/domain/events"
	repository "dms/internal/domain/repositories"
	tools "dms/tools"
	"fmt"
	"strings"
	"time"
)

type deviceContractImpl struct {
	deviceRepo                  repository.DeviceRepository
	userDeviceRepo              repository.UserDeviceRepository
	deviceConfigRepo            repository.DeviceConfigRepository
	historicalDeviceRecordsRepo repository.HistoricalDeviceRecordsRepository
}

func NewDeviceUseCase(
	deviceRepo repository.DeviceRepository,
	userDeviceRepo repository.UserDeviceRepository,
	deviceConfigRepo repository.DeviceConfigRepository,
	historicalDeviceRecordsRepo repository.HistoricalDeviceRecordsRepository) event.DeviceService {
	return &deviceContractImpl{
		deviceRepo:                  deviceRepo,
		userDeviceRepo:              userDeviceRepo,
		deviceConfigRepo:            deviceConfigRepo,
		historicalDeviceRecordsRepo: historicalDeviceRecordsRepo,
	}
}

func (d *deviceContractImpl) CreateDevice(ctx context.Context, device dto.EnteredDeviceInformation) (entity.Device, error) {
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

func (d *deviceContractImpl) SetupDevice(ctx context.Context, deviceConfig dto.FieldsDeviceConfig, device_id string) (entity.DeviceConfig, error) {
	lowercaseFields := make(map[string]interface{})
	for key, value := range deviceConfig.Fields {
		lowercaseFields[strings.ToLower(key)] = value
	}

	entityDeviceConfig := entity.DeviceConfig{
		Device_Id: device_id,
		Fields:    deviceConfig.Fields,
	}

	err := d.deviceConfigRepo.Insert(ctx, entityDeviceConfig)
	if err != nil {
		return entity.DeviceConfig{}, fmt.Errorf("error when creating device configuration: %w", err)
	}

	err = d.historicalDeviceRecordsRepo.CreateMeasurement(ctx, entityDeviceConfig)
	if err != nil {
		return entity.DeviceConfig{}, fmt.Errorf("error when creating device measurement: %w", err)
	}

	return entityDeviceConfig, nil
}

func (d *deviceContractImpl) CreateRecordsDevice(ctx context.Context, deviceRecords dto.FieldsDeviceRecords, device_id string) (entity.HistoricalDeviceRecords, error) {
	deviceRecords.Fields = tools.ToLowerCaseKeyMap(deviceRecords.Fields)

	deviceConfigInfo, err := d.deviceConfigRepo.FindByDeviceId(ctx, device_id)
	if err != nil {
		return entity.HistoricalDeviceRecords{}, fmt.Errorf("error retrieving device config: %w", err)
	}

	convertedFields, err := tools.ConvertFields(deviceRecords.Fields, deviceConfigInfo.Fields)
	if err != nil {
		return entity.HistoricalDeviceRecords{}, fmt.Errorf("error when converting data type: %w", err)
	}

	entityDeviceRecords := entity.HistoricalDeviceRecords{
		Device_Id: device_id,
		Fields:    convertedFields,
	}

	err = d.historicalDeviceRecordsRepo.WriteData(ctx, entityDeviceRecords)
	if err != nil {
		return entity.HistoricalDeviceRecords{}, fmt.Errorf("error when creating device records: %w", err)
	}

	return entityDeviceRecords, nil
}

func (d *deviceContractImpl) ReadRecordsDevice(ctx context.Context, device_id string) ([]map[string]interface{}, error) {
	data, err := d.historicalDeviceRecordsRepo.ReadData(ctx, device_id)
	if err != nil {
		return []map[string]interface{}{}, fmt.Errorf("error when reading device records: %w", err)
	}

	return data, nil
}

func (d *deviceContractImpl) FindInfoByDeviceId(ctx context.Context, device_id string) (dto.DeviceInformation, error) {
	if device_id == "" {
		return dto.DeviceInformation{}, fmt.Errorf("device ID cannot be empty")
	}

	deviceEntity, err := d.deviceRepo.FindInfoByDeviceId(ctx, device_id)
	if err != nil {
		return dto.DeviceInformation{}, fmt.Errorf("error when fetching device information: %w", err)
	}

	return dto.DeviceInformation{
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

func (d *deviceContractImpl) FindAssosiatedUserByDeviceId(ctx context.Context, device_id string) ([]dto.AssosiatedUserInfo, error) {
	if device_id == "" {
		return nil, fmt.Errorf("device ID cannot be empty")
	}

	users, err := d.deviceRepo.FindAssosiatedUserByDeviceId(ctx, device_id)
	if err != nil {
		return nil, fmt.Errorf("error when fetching associated users: %w", err)
	}

	var associatedUsers []dto.AssosiatedUserInfo
	for _, user := range users {
		associatedUsers = append(associatedUsers, dto.AssosiatedUserInfo{
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
