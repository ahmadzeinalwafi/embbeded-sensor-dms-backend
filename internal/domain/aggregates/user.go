package aggregate

import "time"

type EnteredUserInformation struct {
	Name          string
	Email         string
	Password_Hash string
}

type UserInformation struct {
	User_Id       string
	Name          string
	Email         string
	Created_At    time.Time
}
