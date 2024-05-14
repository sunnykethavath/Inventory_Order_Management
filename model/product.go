package model

import "time"

type Product struct {
	ID           uint `json:"id" gorm:"primaryKey"`
	CreatedAt    time.Time
	Name         string `json:"name"`
	Price        uint   `json:"price"`
	SerialNumber string `json:"serial_number"`
}
