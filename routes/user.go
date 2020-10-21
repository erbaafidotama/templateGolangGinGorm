package routes

import (
	"sewaAset/config"
	"sewaAset/models"
	"time"

	"github.com/gin-gonic/gin"
)

func GetUser(c *gin.Context) {
	users := []models.User{}

	// select * from User
	if err := config.DB.Find(&users).Error; err != nil {
		// return error
		c.JSON(404, gin.H{
			"status":  "error",
			"message": "record not found",
		})
		c.Abort()
		return
	}

	// return complete
	c.JSON(200, gin.H{
		"message": "GET data user",
		"data":    users,
	})
}

func PostUser(c *gin.Context) {
	var roleAdmin bool
	// convert string date to date db
	dateStr := c.PostForm("date_birth")
	format := "2006-01-02"
	date, _ := time.Parse(format, dateStr)

	// check admin role
	if c.PostForm("admin_role") == "true" {
		roleAdmin = true
	}

	// make object from form body
	items := models.User{
		NIK:       c.PostForm("nik"),
		FullName:  c.PostForm("full_name"),
		DateBirth: date,
		AdminRole: roleAdmin,
	}

	// crete data to db
	config.DB.Create(&items)

	c.JSON(200, gin.H{
		"status": "berhasil post",
		"data":   items,
	})
}

func UpdateUser(c *gin.Context) {
	var roleAdmin bool

	// get id from url
	userId := c.Param("id")

	var dataUser models.User
	if err := config.DB.Where("id = ?", userId).First(&dataUser).Error; err != nil {
		c.JSON(404, gin.H{
			"status":  "error",
			"message": "record not found",
		})
		c.Abort()
		return
	}

	// convert string date to date db
	dateStr := c.PostForm("date_birth")
	format := "2006-01-02"
	date, _ := time.Parse(format, dateStr)

	if c.PostForm("admin_role") == "true" {
		roleAdmin = true
	}

	config.DB.Model(&dataUser).Where("id = ?", userId).Updates(models.User{
		FullName:  c.PostForm("full_name"),
		DateBirth: date,
		AdminRole: roleAdmin,
	})

	c.JSON(200, gin.H{
		"status": "Success",
		"data":   dataUser,
	})
}

func DeleteUser(c *gin.Context) {
	// get id from url
	userId := c.Param("id")

	var dataUser models.User
	if err := config.DB.Where("id = ?", userId).First(&dataUser).Error; err != nil {
		c.JSON(404, gin.H{
			"status":  "error",
			"message": "record not found",
		})
		c.Abort()
		return
	}

	config.DB.Where("id = ?", userId).Delete(&dataUser)

	c.JSON(200, gin.H{
		"status": "Success Delete",
		"data":   dataUser,
	})
}
