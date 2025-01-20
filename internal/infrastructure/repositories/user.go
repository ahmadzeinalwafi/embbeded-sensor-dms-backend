package repository

import (
	"context"
	"database/sql"
	entity "dms/internal/domain/entities"
	repository "dms/internal/domain/repositories"
	"errors"
	"fmt"
)

type UserRepositoryImpl struct {
	DB *sql.DB
}

func NewUserRepository(db *sql.DB) repository.UserRepository {
	return &UserRepositoryImpl{DB: db}
}

func (repository *UserRepositoryImpl) Insert(ctx context.Context, user entity.User) (entity.User, error) {
	script := "INSERT INTO users(user_id, name, email, password_hash) VALUES (?,?,?,?)"
	result, err := repository.DB.ExecContext(ctx,
		script,
		user.User_Id,
		user.Name,
		user.Email,
		user.Password_Hash)
	if err != nil {
		return user, err
	}
	id, err := result.LastInsertId()
	if err != nil {
		return user, err
	}
	user.User_Id = fmt.Sprint(id)
	return user, nil
}

func (repository *UserRepositoryImpl) FindByUserId(ctx context.Context, user_id string) (entity.User, error) {
	script := "SELECT user_id, name, email, password_hash, created_at FROM users WHERE user_id = ? LIMIT 1"
	var user entity.User

	// Use QueryRowContext to fetch a single record
	err := repository.DB.QueryRowContext(ctx, script, user_id).Scan(
		&user.User_Id,
		&user.Name,
		&user.Email,
		&user.Password_Hash,
		&user.Created_At,
	)
	if err != nil {
		if err == sql.ErrNoRows {
			return user, errors.New("no user found with id " + user_id)
		}
		return user, err
	}

	return user, nil
}

func (repository *UserRepositoryImpl) FindByEmail(ctx context.Context, email string) (entity.User, error) {
	script := "SELECT user_id, name, email, password_hash, created_at FROM users WHERE email = ? LIMIT 1"
	var user entity.User

	// Use QueryRowContext to fetch a single record
	err := repository.DB.QueryRowContext(ctx, script, email).Scan(
		&user.User_Id,
		&user.Name,
		&user.Email,
		&user.Password_Hash,
		&user.Created_At,
	)
	if err != nil {
		if err == sql.ErrNoRows {
			return user, errors.New("no user found with email " + email)
		}
		return user, err
	}

	return user, nil
}

func (repository *UserRepositoryImpl) DeleteById(ctx context.Context, user_id string) error {
	script := "DELETE FROM users WHERE user_id = ?"
	result, err := repository.DB.ExecContext(ctx, script, user_id)
	if err != nil {
		return err
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return err
	}

	if rowsAffected == 0 {
		return errors.New("no user found with id " + user_id)
	}

	return nil
}
