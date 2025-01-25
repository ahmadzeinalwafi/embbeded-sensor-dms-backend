package dto

import "time"

type EnteredUserInformation struct {
	Name     string `validate:"required"`
	Email    string `validate:"required,email"`
	Password string `validate:"required"`
}

type UserInformation struct {
	User_Id    string
	Name       string
	Email      string
	Created_At time.Time
}

type AuthUserInformation struct {
	User_Id string
	Name    string
	Email   string
	Token   string
}

type UserCredential struct {
	Email    string `validate:"required,email"`
	Password string `validate:"required"`
}