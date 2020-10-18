package routes

import (
	"sewaAset/config"
	"sewaAset/models"
	"time"

	"github.com/gin-gonic/gin"
)

func GetUser(c *gin.Context) {
	// declare model User
	var user models.User

	// select * from User
	if err := config.DB.Find(&user).Error; err != nil {
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
		"data":    user,
	})
}

func PostUser(c *gin.Context) {

	// convert string date to date db
	dateStr := c.PostForm("date_birth")
	i, _ := time.Parse(time.RFC3339, dateStr)

	// make object from form body
	items := models.User{
		NIK:       c.PostForm("nik"),
		FullName:  c.PostForm("full_name"),
		DateBirth: i,
	}

	// crete data to db
	config.DB.Create(&items)

	c.JSON(200, gin.H{
		"status": "berhasil post",
		"data":   items,
	})
}
