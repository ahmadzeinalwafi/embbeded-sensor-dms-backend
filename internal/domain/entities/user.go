package entity

import "time"

type User struct {
	User_Id       string
	Name          string
	Email         string
	Password_Hash string
	Created_At    time.Time
}
