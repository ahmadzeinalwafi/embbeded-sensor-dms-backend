package entity

import "time"

type Device struct {
	Device_Id		string
	Name			string
	Type			string
	Location 		string
	Token 			string
	Status 			string
	Description		string
	Created_At    	time.Time
}
