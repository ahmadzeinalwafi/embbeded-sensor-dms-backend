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

type userContractImpl struct {
	repo repository.UserRepository

}

func NewUserUseCase(repo repository.UserRepository) event.UserService {
	return &userContractImpl{
		repo: repo,
	}
}

func (u *userContractImpl) CreateUser(ctx context.Context, user aggregate.EnteredUserInformation) (entity.User, error) {
	hashed_password, err := tools.HashPassword(user.Password)
	if err != nil {
		panic(err)
	}

	newUserEntity := entity.User{
		User_Id:       tools.GenerateShortID(),
		Name:          user.Name,
		Email:         user.Email,
		Password_Hash: hashed_password,
	}

	_, err = u.repo.Insert(ctx, newUserEntity)
	if err != nil {
		return entity.User{}, fmt.Errorf("error when creating user: %w", err)
	}

	return newUserEntity, nil
}

func (u *userContractImpl) GetUserInfo(ctx context.Context, user_id string) (aggregate.UserInformation, error) {
	if user_id == "" {
		return aggregate.UserInformation{}, fmt.Errorf("user ID cannot be empty")
	}

	userEntity, err := u.repo.FindByUserId(ctx, user_id)
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

func (u *userContractImpl) GetUserToken(ctx context.Context, credential aggregate.UserCredential) (aggregate.AuthUserInformation, error) {
	userEntity, err := u.repo.FindByEmail(ctx, credential.Email)
	if err != nil {
		return aggregate.AuthUserInformation{}, fmt.Errorf("fetching user info: %w", err)
	}

	if status, err := tools.VerifyPassword(credential.Password, userEntity.Password_Hash); err != nil {
		return aggregate.AuthUserInformation{}, fmt.Errorf("error when verify password: %w", err)
	} else if !status {
		return aggregate.AuthUserInformation{}, fmt.Errorf("wrong password")
	}

	token, err := tools.GenerateToken(userEntity.User_Id, userEntity.Email, 1 * time.Hour)
	if err != nil {
		return aggregate.AuthUserInformation{}, fmt.Errorf("generate token: %w", err)
	}

	return aggregate.AuthUserInformation{
		User_Id: userEntity.User_Id,
		Name:    userEntity.Name,
		Email:   userEntity.Email,
		Token:   token,
	}, nil
}

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
