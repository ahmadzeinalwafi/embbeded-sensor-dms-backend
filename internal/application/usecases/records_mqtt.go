package usecase

import (
	"context"
	entity "dms/internal/domain/entities"
	repository "dms/internal/domain/repositories"
	tools "dms/tools"
	"fmt"
	"log"
)

type MessageProcessorUseCase interface {
	ProcessMessage(ctx context.Context, messageData entity.HistoricalDeviceRecords) error
}

type MessageProcessor struct {
	historicalDeviceRecordsRepo repository.HistoricalDeviceRecordsRepository
	deviceConfigRepo            repository.DeviceConfigRepository
}

func NewMessageProcessorUseCase(DeviceRecordsRepo repository.HistoricalDeviceRecordsRepository, DeviceConfigRepo repository.DeviceConfigRepository) *MessageProcessor {
	return &MessageProcessor{
		historicalDeviceRecordsRepo: DeviceRecordsRepo,
		deviceConfigRepo:            DeviceConfigRepo,
	}
}

func (m *MessageProcessor) ProcessMessage(ctx context.Context, data entity.HistoricalDeviceRecords) error {
	data.Fields = tools.ToLowerCaseKeyMap(data.Fields)

	deviceConfigInfo, err := m.deviceConfigRepo.FindByDeviceId(ctx, data.Device_Id)
	if err != nil {
		return fmt.Errorf("error retrieving device config: %w", err)
	}

	log.Printf("config: %s", deviceConfigInfo)

	convertedFields, err := tools.ConvertFields(data.Fields, deviceConfigInfo.Fields)
	if err != nil {
		return err
	}

	entityDeviceRecords := entity.HistoricalDeviceRecords{
		Device_Id: data.Device_Id,
		Fields:    convertedFields,
	}

	log.Printf("result: %s", convertedFields)

	err = m.historicalDeviceRecordsRepo.WriteData(ctx, entityDeviceRecords)
	if err != nil {
		return fmt.Errorf("error when creating device records: %w", err)
	}

	return nil
}
