package models

import (
	"time"

	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	TransactionItem []TransactionItem
	NIK             string
	FullName        string
	DateBirth       time.Time `json:"date_birth"`
	AdminRole       bool      `gorm:"default:0"`
}
