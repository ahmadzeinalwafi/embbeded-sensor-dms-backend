package usecase

import (
	"context"
	entity "dms/internal/domain/entities"
	repository "dms/internal/domain/repositories"
	tools "dms/tools"
	"log"
	"fmt"
)

type MessageProcessorUseCase interface {
	ProcessMessage(ctx context.Context, messageData entity.HistoricalDeviceRecords) error
}

type MessageProcessor struct {
	historicalDeviceRecordsRepo repository.HistoricalDeviceRecordsRepository
	deviceConfigRepo repository.DeviceConfigRepository
}

func NewMessageProcessorUseCase(DeviceRecordsRepo repository.HistoricalDeviceRecordsRepository, DeviceConfigRepo repository.DeviceConfigRepository) *MessageProcessor {
	return &MessageProcessor{
		historicalDeviceRecordsRepo: DeviceRecordsRepo,
		deviceConfigRepo: DeviceConfigRepo,
	}
}

func (m *MessageProcessor) ProcessMessage(ctx context.Context, data entity.HistoricalDeviceRecords) error {
	deviceConfigInfo, err := m.deviceConfigRepo.FindByDeviceId(ctx, data.Device_Id)
	if err != nil {
		return fmt.Errorf("error retrieving device config: %w", err)
	}

	log.Printf("config: %s", deviceConfigInfo)

	// Prepare a new map for the converted fields
	convertedFields := make(map[string]interface{})

	// Iterate over the device configuration fields and perform type conversion
	for field, fieldType := range deviceConfigInfo.Fields {
		value, exists := data.Fields[field]
		if !exists {
			return fmt.Errorf("field %s is missing in device records", field)
		}

		// Perform type conversion based on the field type
		switch fieldType {
		case "float16", "float32", "float64":
			convertedFields[field], err = tools.ToFloat(value)
			if err != nil {
				return fmt.Errorf("error converting field %s to %s: %w", field, fieldType, err)
			}
		case "int8", "int16", "int32", "int64":
			convertedFields[field], err = tools.ToInt(value)
			if err != nil {
				return fmt.Errorf("error converting field %s to %s: %w", field, fieldType, err)
			}
		default:
			return fmt.Errorf("unsupported field type %s for field %s", fieldType, field)
		}
	}

	// Create the entity with the converted fields
	entityDeviceRecords := entity.HistoricalDeviceRecords{
		Device_Id: data.Device_Id,
		Fields:    convertedFields,
	}

	log.Printf("result: %s", convertedFields)

	// Write the data to the repository
	err = m.historicalDeviceRecordsRepo.WriteData(ctx, entityDeviceRecords)
	if err != nil {
		return fmt.Errorf("error when creating device records: %w", err)
	}

	return nil
}
