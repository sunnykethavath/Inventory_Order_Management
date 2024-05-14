package model

import "time"

// User Model
type User struct {
	ID        uint `json:"id" gorm:"primaryKey"`
	CreatedAt time.Time
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	PhoneNo   string `json:"phone_no"`
}
