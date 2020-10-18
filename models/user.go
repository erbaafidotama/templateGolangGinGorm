package models

import (
	"time"

	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	NIK       string
	FullName  string
	DateBirth time.Time
	AdminRole bool `gorm:"default:0"`
}
