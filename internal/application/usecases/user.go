package usecase

import (
	"context"
	dto "dms/internal/domain/data_transfer_object"
	entity "dms/internal/domain/entities"
	event "dms/internal/domain/events"
	repository "dms/internal/domain/repositories"
	tools "dms/tools"
	"errors"
	"fmt"
	"time"

	"github.com/go-sql-driver/mysql"
)

type userContractImpl struct {
	repo repository.UserRepository
}

func NewUserUseCase(repo repository.UserRepository) event.UserService {
	return &userContractImpl{
		repo: repo,
	}
}

func (u *userContractImpl) CreateUser(ctx context.Context, user dto.EnteredUserInformation) (entity.User, error) {
	var mysqlErr *mysql.MySQLError
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
	if errors.As(err, &mysqlErr) && mysqlErr.Number == 1062 {
		return entity.User{}, fmt.Errorf("duplicate entry: %w", err)
	} else if err != nil {
		return entity.User{}, fmt.Errorf("error when creating user: %w", err)
	}

	return newUserEntity, nil
}

func (u *userContractImpl) GetUserInfo(ctx context.Context, user_id string) (dto.UserInformation, error) {
	if user_id == "" {
		return dto.UserInformation{}, fmt.Errorf("user ID cannot be empty")
	}

	userEntity, err := u.repo.FindByUserId(ctx, user_id)
	if err != nil {
		return dto.UserInformation{}, fmt.Errorf("error when fetching user information: %w", err)
	}

	return dto.UserInformation{
		User_Id:    userEntity.User_Id,
		Name:       userEntity.Name,
		Email:      userEntity.Email,
		Created_At: userEntity.Created_At,
	}, nil
}

func (u *userContractImpl) FindAssosiatedDevicesByUserId(ctx context.Context, user_id string) ([]dto.DeviceInformation, error) {
	if user_id == "" {
		return nil, fmt.Errorf("user ID cannot be empty")
	}

	// devices, err := u.repo.FindAssosiatedDevicesByUserId(ctx, user_id)
	devices, err := u.repo.FindAssosiatedDevicesByUserId(ctx, user_id)
	if err != nil {
		return nil, fmt.Errorf("error when fetching associated devices: %w", err)
	}

	var associatedDevices []dto.DeviceInformation
	if devices == nil {
		return associatedDevices, nil
	}

	for _, device := range devices {
		associatedDevices = append(associatedDevices, dto.DeviceInformation{
			Device_Id:   device.Device_Id,
			Name:        device.Name,
			Type:        device.Type,
			Location:    device.Location,
			Token:       device.Token,
			Status:      device.Status,
			Description: device.Description,
			Created_At:  device.Created_At,
		})
	}

	return associatedDevices, nil
}

func (u *userContractImpl) GetUserToken(ctx context.Context, credential dto.UserCredential) (dto.AuthUserInformation, error) {
	userEntity, err := u.repo.FindByEmail(ctx, credential.Email)
	if err != nil {
		return dto.AuthUserInformation{}, fmt.Errorf("fetching user info: %w", err)
	}

	if status, err := tools.VerifyPassword(credential.Password, userEntity.Password_Hash); err != nil {
		return dto.AuthUserInformation{}, fmt.Errorf("error when verify password: %w", err)
	} else if !status {
		return dto.AuthUserInformation{}, fmt.Errorf("wrong password")
	}

	token, err := tools.GenerateToken(userEntity.User_Id, userEntity.Email, 1*time.Hour)
	if err != nil {
		return dto.AuthUserInformation{}, fmt.Errorf("generate token: %w", err)
	}

	return dto.AuthUserInformation{
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
