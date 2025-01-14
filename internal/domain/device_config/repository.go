package repository

import (
	"context"
	entity "golang_api/internal/domain/device_config/entities"
)

type SensorRepository interface {
	Insert(ctx context.Context, sensor *entity.SensorConfig) error
	FindByID(ctx context.Context, sensorID string) (*entity.SensorConfig, error)
}
