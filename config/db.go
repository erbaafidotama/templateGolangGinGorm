package config

import (
	"os"
	"sewaAset/models"

	"github.com/subosito/gotenv"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func InitDB() {
	// get ENV
	gotenv.Load()

	var err error
	dbName := os.Getenv("DB_NAME")

	// connect to db
	dsn := "root:pass@tcp(127.0.0.1:3306)/" + dbName + "?charset=utf8mb4&parseTime=True&loc=Local"
	DB, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}

	// migrate table
	DB.AutoMigrate(&models.User{})
	DB.AutoMigrate(&models.Item{})
	DB.AutoMigrate(&models.TransactionItem{})
}
