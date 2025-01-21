package repository

import (
	"context"
	"database/sql"
	entity "dms/internal/domain/entities"
	repository "dms/internal/domain/repositories"
	"errors"
	"fmt"
)

type DeviceRepositoryImpl struct {
	DB *sql.DB
}

func NewDeviceRepository(db *sql.DB) repository.DeviceRepository {
	return &DeviceRepositoryImpl{DB: db}
}

func (repository *DeviceRepositoryImpl) Insert(ctx context.Context, device entity.Device) (entity.Device, error) {
	script := "INSERT INTO devices(device_id, name, type, location, token, status, description) VALUES (?,?,?,?,?,?,?)"
	result, err := repository.DB.ExecContext(ctx,
		script,
		device.Device_Id,
		device.Name,
		device.Type,
		device.Location,
		device.Token,
		device.Status,
		device.Description,
	)
	if err != nil {
		return device, err
	}
	id, err := result.LastInsertId()
	if err != nil {
		return device, err
	}
	device.Device_Id = fmt.Sprint(id)
	return device, nil
}

func (repository *DeviceRepositoryImpl) FindInfoByDeviceId(ctx context.Context, device_id string) (entity.Device, error) {
	script := "SELECT device_id, name, type, location, token, status, description, created_at FROM devices WHERE device_id = ? LIMIT 1"
	var device entity.Device

	err := repository.DB.QueryRowContext(ctx, script, device_id).Scan(
		&device.Device_Id,
		&device.Name,
		&device.Type,
		&device.Location,
		&device.Token,
		&device.Status,
		&device.Description,
		&device.Created_At,
	)
	if err != nil {
		if err == sql.ErrNoRows {
			return device, errors.New("no device found with id " + device_id)
		}
		return device, err
	}

	return device, nil
}

func (repository *DeviceRepositoryImpl) FindAssosiatedUserByDeviceId(ctx context.Context, device_id string) ([]entity.User, error) {
	script := `
		SELECT u.user_id, u.name, u.email, u.created_at 
		FROM users u
		INNER JOIN user_devices du ON u.user_id = du.user_id
		WHERE du.device_id = ?
	`
	rows, err := repository.DB.QueryContext(ctx, script, device_id)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var users []entity.User
	for rows.Next() {
		var user entity.User
		err := rows.Scan(
			&user.User_Id,
			&user.Name,
			&user.Email,
			&user.Created_At,
		)
		if err != nil {
			return nil, err
		}
		users = append(users, user)
	}
	if len(users) == 0 {
		return nil, errors.New("no users associated with device id " + device_id)
	}

	return users, nil
}

func (repository *DeviceRepositoryImpl) DeleteByDeviceId(ctx context.Context, device_id string) error {
	script := "DELETE FROM devices WHERE device_id = ?"
	result, err := repository.DB.ExecContext(ctx, script, device_id)
	if err != nil {
		return err
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return err
	}

	if rowsAffected == 0 {
		return errors.New("no device found with id " + device_id)
	}

	return nil
}
