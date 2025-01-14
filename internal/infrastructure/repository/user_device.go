package repository

import (
	"context"
	"database/sql"
	"errors"
	repository "golang_api/internal/domain/user_device"
	entity "golang_api/internal/domain/user_device/entities"
)

type userDevicesRepositoryImpl struct {
	DB *sql.DB
}

func NewUserDeviceRepository(db *sql.DB) repository.UserDeviceRepository {
	return &userDevicesRepositoryImpl{DB: db}
}

func (repository *userDevicesRepositoryImpl) Insert(ctx context.Context, user_device entity.UserDevice) (entity.UserDevice, error) {
	script := "INSERT INTO user_devices(user_id, device_id) VALUES (?,?)"
	result, err := repository.DB.ExecContext(ctx, script, user_device.User_Id, user_device.Device_Id)
	if err != nil {
		return user_device, err
	}
	id, err := result.LastInsertId()
	if err != nil {
		return user_device, err
	}
	user_device.Id = int32(id)
	return user_device, nil
}

func (repository *userDevicesRepositoryImpl) FindByUserId(ctx context.Context, id string) (entity.UserDevice, error) {
	script := "SELECT id, created_at, user_id, device_id FROM user_devices WHERE user_id = ? LIMIT 1"
	rows, err := repository.DB.QueryContext(ctx, script, id)
	user_device := entity.UserDevice{}
	if err != nil {
		return user_device, err
	}
	defer rows.Close()
	if rows.Next() {
		rows.Scan(&user_device.Id, &user_device.User_Id, &user_device.Device_Id)
		return user_device, nil
	} else {
		return user_device, errors.New("Id " + id + " Not Found")
	}
}

func (repository *userDevicesRepositoryImpl) FindAll(ctx context.Context, limit int32) ([]entity.UserDevice, error) {
	script := "SELECT id, created_at, user_id, device_id FROM user_devices LIMIT ?"
	rows, err := repository.DB.QueryContext(ctx, script, limit)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var user_devices []entity.UserDevice
	for rows.Next() {
		comment := entity.UserDevice{}
		rows.Scan(&comment.Id, &comment.User_Id, &comment.Device_Id)
		user_devices = append(user_devices, comment)
	}
	return user_devices, nil
}

func (repository *userDevicesRepositoryImpl) DeleteBySensorId(ctx context.Context, id string) error {
	script := "DELETE FROM user_devices WHERE device_id = ?"
	_, err := repository.DB.ExecContext(ctx, script, "alpha01")
	if err != nil {
		return err
	}
	return nil
}
