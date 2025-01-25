package event

import (
	"context"
	dto "dms/internal/domain/data_transfer_object"
	entity "dms/internal/domain/entities"
)

type UserService interface {
	CreateUser(ctx context.Context, user dto.EnteredUserInformation) (entity.User, error)
	GetUserInfo(ctx context.Context, user_id string) (dto.UserInformation, error)
	FindAssosiatedDevicesByUserId(ctx context.Context, user_id string) ([]dto.DeviceInformation, error)
	GetUserToken(ctx context.Context, credential dto.UserCredential) (dto.AuthUserInformation, error)
	DeleteUserById(ctx context.Context, user_id string) error
}
