package repository

import (
	"context"
	entity "dms/internal/domain/entities"
)

type UserRepository interface {
	Insert(ctx context.Context, comment entity.User) (entity.User, error)
	FindByUserId(ctx context.Context, user_id string) (entity.User, error)
	FindByEmail(ctx context.Context, email string) (entity.User, error)
	DeleteById(ctx context.Context, id string) error
}
