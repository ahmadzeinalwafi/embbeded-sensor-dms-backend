package usecase

import (
	"context"
	"github.com/go-playground/validator/v10"
	aggregate "dms/internal/domain/aggregates"
	entity "dms/internal/domain/entities"
	event "dms/internal/domain/events"
	repository "dms/internal/domain/repositories"
	tools "dms/tools"
	"fmt"
)

type userContractImpl struct {
	repo repository.UserRepository
}

func NewUserUseCase(repo repository.UserRepository) event.UserService {
	return &userContractImpl{
		repo: repo,
	}
}

// CreateUser implements the logic for creating a new user.
func (u *userContractImpl) CreateUser(ctx context.Context, user aggregate.EnteredUserInformation) (entity.User, error) {
	validate := validator.New()
	err := validate.Struct(user)
	if err != nil {
		return entity.User{}, fmt.Errorf("invalid user input: %w", err)
	}

	newUserEntity := entity.User{
		User_Id:       tools.GenerateShortID(),
		Name:          user.Name,
		Email:         user.Email,
		Password_Hash: user.Password_Hash,
	}

	_, err = u.repo.Insert(ctx, newUserEntity)
	if err != nil {
		return entity.User{}, fmt.Errorf("error when creating user: %w", err)
	}

	return newUserEntity, nil
}

// GetUserInfo fetches user details by user ID.
func (u *userContractImpl) GetUserInfo(ctx context.Context, user_id string) (aggregate.UserInformation, error) {
	if user_id == "" {
		return aggregate.UserInformation{}, fmt.Errorf("user ID cannot be empty")
	}

	userEntity, err := u.repo.FindById(ctx, user_id)
	if err != nil {
		return aggregate.UserInformation{}, fmt.Errorf("error when fetching user information: %w", err)
	}

	return aggregate.UserInformation{
		User_Id:    userEntity.User_Id,
		Name:       userEntity.Name,
		Email:      userEntity.Email,
		Created_At: userEntity.Created_At,
	}, nil
}

// DeleteUserById deletes a user by their ID.
func (u *userContractImpl) DeleteUserById(ctx context.Context, user_id string) error {
	if user_id == "" {
		return fmt.Errorf("user ID cannot be empty")
	}

	err := u.repo.DeleteById(ctx, user_id)
	if err != nil {
		return fmt.Errorf("error when deleting user: %w", err)
	}

	return nil
}
