package repository

import (
	"context"
	"fmt"
	entity "golang_api/internal/domain/user_device/entities"
	MySQLConnector "golang_api/internal/infrastructure/database/mysql"
	"testing"

	_ "github.com/go-sql-driver/mysql"
)

func TestCommentInsert(t *testing.T) {
	userDeviceRepository := NewUserDeviceRepository(MySQLConnector.GetConnection())

	ctx := context.Background()
	user_device := entity.UserDevice{
		User_Id:   "alpha01",
		Device_Id: "alpha01",
	}

	result, err := userDeviceRepository.Insert(ctx, user_device)
	if err != nil {
		panic(err)
	}
	fmt.Println(result)
}

func TestFindByUserId(t *testing.T) {
	userDeviceRepository := NewUserDeviceRepository(MySQLConnector.GetConnection())

	user_device, err := userDeviceRepository.FindById(context.Background(), "alpha01")
	if err != nil {
		panic(err)
	}

	fmt.Println(user_device)
}

func TestFindAll(t *testing.T) {
	userDeviceRepository := NewUserDeviceRepository(MySQLConnector.GetConnection())

	user_device, err := userDeviceRepository.FindAll(context.Background(), 5)
	if err != nil {
		panic(err)
	}

	for _, comment := range user_device {
		fmt.Println(comment)
	}
}
