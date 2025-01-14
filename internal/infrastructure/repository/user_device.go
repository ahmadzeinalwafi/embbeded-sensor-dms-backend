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

func (repository *userDevicesRepositoryImpl) Insert(ctx context.Context, comment entity.UserDevice) (entity.UserDevice, error) {
	script := "INSERT INTO user_devices(user_id, device_id) VALUES (?,?)"
	result, err := repository.DB.ExecContext(ctx, script, comment.User_Id, comment.Device_Id)
	if err != nil {
		return comment, err
	}
	id, err := result.LastInsertId()
	if err != nil {
		return comment, err
	}
	comment.Id = int32(id)
	return comment, nil
}

func (repository *userDevicesRepositoryImpl) FindById(ctx context.Context, id string) (entity.UserDevice, error) {
	script := "SELECT id, created_at, user_id, device_id FROM user_devices WHERE user_id = ? LIMIT 1"
	rows, err := repository.DB.QueryContext(ctx, script, id)
	comment := entity.UserDevice{}
	if err != nil {
		return comment, err
	}
	defer rows.Close()
	if rows.Next() {
		rows.Scan(&comment.Id, &comment.User_Id, &comment.Device_Id)
		return comment, nil
	} else {
		return comment, errors.New("Id " + id + " Not Found")
	}
}

func (repository *userDevicesRepositoryImpl) FindAll(ctx context.Context, limit int32) ([]entity.UserDevice, error) {
	script := "SELECT id, created_at, user_id, device_id FROM user_devices LIMIT ?"
	rows, err := repository.DB.QueryContext(ctx, script, limit)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var comments []entity.UserDevice
	for rows.Next() {
		comment := entity.UserDevice{}
		rows.Scan(&comment.Id, &comment.User_Id, &comment.Device_Id)
		comments = append(comments, comment)
	}
	return comments, nil
}
