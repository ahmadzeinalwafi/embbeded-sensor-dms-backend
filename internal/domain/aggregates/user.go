package aggregate

import "time"

type EnteredUserInformation struct {
	Name          string `validate:"required"`
	Email         string `validate:"required,email"`
	Password string `validate:"required"`
}

type UserInformation struct {
	User_Id    string
	Name       string
	Email      string
	Created_At time.Time
}
