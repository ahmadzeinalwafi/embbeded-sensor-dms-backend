package repository

import (
	"context"
	"database/sql"
	entity "dms/internal/domain/entities"
	MySQLConnector "dms/internal/infrastructure/database/mysql"
	"fmt"
	"testing"

	_ "github.com/go-sql-driver/mysql"
)

func setupTestDB() *sql.DB {
	// Initialize a connection to the test database
	return MySQLConnector.GetConnection()
}

func TestUserRepository_Insert(t *testing.T) {
	userRepository := NewUserRepository(setupTestDB())

	ctx := context.Background()
	user := entity.User{
		User_Id:        "test_user_01",
		Name:           "Test User",
		Email:          "testuser@example.com",
		Password_Hash:  "hashed_password",
	}

	result, err := userRepository.Insert(ctx, user)
	if err != nil {
		t.Fatalf("Error inserting user: %v", err)
	}

	fmt.Printf("Inserted user: %+v\n", result)
	if result.User_Id == "" {
		t.Error("Expected User_Id to be populated after insert")
	}
}

func TestUserRepository_FindById(t *testing.T) {
	userRepository := NewUserRepository(setupTestDB())
	ctx := context.Background()
	testUserID := "test_user_01"

	userDevice, err := userRepository.FindById(ctx, testUserID)
	if err != nil {
		t.Fatalf("Error finding user by ID: %v", err)
	}

	fmt.Printf("Found user device: %+v\n", userDevice)
	if userDevice.User_Id != testUserID {
		t.Errorf("Expected User_Id to be %s, got %s", testUserID, userDevice.User_Id)
	}
}

func TestUserRepository_DeleteById(t *testing.T) {
	userRepository := NewUserRepository(setupTestDB())

	ctx := context.Background()
	testUserID := "test_user_01"

	err := userRepository.DeleteById(ctx, testUserID)
	if err != nil {
		t.Fatalf("Error deleting user by ID: %v", err)
	}

	fmt.Println("Deleted user successfully")

	// Verify the user was deleted
	_, err = userRepository.FindById(ctx, testUserID)
	if err == nil {
		t.Error("Expected error when finding deleted user, but got none")
	} else {
		fmt.Printf("Verified user deletion: %v\n", err)
	}
}
