package repositories

import (
	"context"
	"fmt"
	entity "golang_api/internal/domain/user_device/entities"
	MySQLConnector "golang_api/internal/infrastructure/database/mysql"
	"testing"

	_ "github.com/go-sql-driver/mysql"
)

func TestUserDeviceInsert(t *testing.T) {
	userDeviceRepository := NewUserDeviceRepository(MySQLConnector.GetConnection())

	ctx := context.Background()
	user_device := entity.UserDevice{
		User_Id:   "alpha01",
		Device_Id: "alpha01",
	}

	result, err := userDeviceRepository.Insert(ctx, user_device)
	if err != nil {
		panic(fmt.Sprintf("Error inserting user device: %v", err))
	}
	fmt.Printf("Inserted user device: %+v\n", result)
}

func TestFindByUserId(t *testing.T) {
	userDeviceRepository := NewUserDeviceRepository(MySQLConnector.GetConnection())

	user_device, err := userDeviceRepository.FindByUserId(context.Background(), "alpha01")
	if err != nil {
		panic(fmt.Sprintf("Error finding user by ID: %v", err))
	}

	fmt.Printf("Found user device: %+v\n", user_device)
}

func TestFindAll(t *testing.T) {
	userDeviceRepository := NewUserDeviceRepository(MySQLConnector.GetConnection())

	user_devices, err := userDeviceRepository.FindAll(context.Background(), 5)
	if err != nil {
		panic(fmt.Sprintf("Error finding all user devices: %v", err))
	}

	for _, user_device := range user_devices {
		fmt.Printf("User device: %+v\n", user_device)
	}
}

func TestDeleteBySensorId(t *testing.T) {
	userDeviceRepository := NewUserDeviceRepository(MySQLConnector.GetConnection())

	err := userDeviceRepository.DeleteBySensorId(context.Background(), "alpha01")
	if err != nil {
		panic(fmt.Sprintf("Error deleting user device: %v", err))
	}
	fmt.Println("Deleted user device successfully")
}
