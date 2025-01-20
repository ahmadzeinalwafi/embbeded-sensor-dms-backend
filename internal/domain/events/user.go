package event

import (
	"context"
	aggregate "dms/internal/domain/aggregates"
	entity "dms/internal/domain/entities"
)

type UserService interface {
	CreateUser(ctx context.Context, user aggregate.EnteredUserInformation) (entity.User, error)
	GetUserInfo(ctx context.Context, user_id string) (aggregate.UserInformation, error)
	GetUserToken(ctx context.Context, credential aggregate.UserCredential) (aggregate.AuthUserInformation, error)
	DeleteUserById(ctx context.Context, user_id string) error
}
