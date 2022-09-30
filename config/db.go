package config

import (
	"fmt"
	"os"
	"sewaAset/models"

	"github.com/subosito/gotenv"
	// "gorm.io/driver/mysql" // if you want use mysql database
	"gorm.io/driver/postgres" // if you want use postgres database
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

var DB *gorm.DB

func InitDB() {
	// get ENV
	gotenv.Load()

	var err error
	dbName := os.Getenv("DB_NAME")
	dbUsername := os.Getenv("DB_USERNAME")
	dbPassword := os.Getenv("DB_PASSWORD")
	dbPort := os.Getenv("DB_PORT")

	// connect to mysql db
	// dsn := "root:pass@tcp(127.0.0.1:3306)/" + dbName + "?charset=utf8mb4&parseTime=True&loc=Local"
	// DB, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})

	// connect to postgresdb
	dsn := "user=" + dbUsername + " password=" + dbPassword + " dbname=" + dbName + " port=" + dbPort + " sslmode=disable TimeZone=Asia/Shanghai"
	fmt.Println(dsn)
	DB, err := gorm.Open(postgres.Open(dsn), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
	})

	if err != nil {
		panic("Connecting database failed:" + err.Error())
	}

	// migrate table
	DB.AutoMigrate(&models.User{})
	DB.AutoMigrate(&models.Item{})
	DB.AutoMigrate(&models.TransactionItem{})
}
