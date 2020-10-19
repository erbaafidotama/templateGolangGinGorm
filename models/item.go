package models

import (
	"gorm.io/gorm"
)

type Item struct {
	gorm.Model
	TransactionItem []TransactionItem
	ItemName        string `sql:"type:varchar(250);"`
	Quantity        int
	Price           float32 `sql:"type:decimal(10,2);"`
	InfoDesc        string  `sql:"type:longtext;"`
}
