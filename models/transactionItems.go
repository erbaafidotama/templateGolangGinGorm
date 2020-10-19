package models

import (
	"time"

	"gorm.io/gorm"
)

type TransactionItem struct {
	gorm.Model
	UserID     uint
	ItemID     uint
	ItemAmount int
	RentAmount int
	OrderDate  time.Time `json:"order_date"`
}
