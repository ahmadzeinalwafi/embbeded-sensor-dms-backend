package entity

import "time"

type UserDevice struct {
	Id         int32
	Created_At time.Time
	User_Id    string
	Device_Id  string
}
