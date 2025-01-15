package repository

import (
	"context"
	entity "golang_api/internal/domain/user_device/entities"
)

type UserDeviceRepository interface {
	Insert(ctx context.Context, comment entity.UserDevice) (entity.UserDevice, error)
	FindByUserId(ctx context.Context, id string) ([]entity.UserDevice, error)
	FindAll(ctx context.Context, limit int32) ([]entity.UserDevice, error)
	DeleteBySensorId(ctx context.Context, id string) error
}
