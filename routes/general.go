package routes

import (
	"fmt"
	"os"
	"sewaAset/config"
	"sewaAset/models"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)

func Login(c *gin.Context) {
	var userData models.User

	nik := c.PostForm("nik")
	dateStr := c.PostForm("date_birth")
	format := "2006-01-02"
	date, _ := time.Parse(format, dateStr)

	if err := config.DB.Where("nik = ? AND date_birth = ?", nik, date).First(&userData).Error; err != nil {
		c.JSON(404, gin.H{
			"status":  "error",
			"message": "record not found",
		})
		c.Abort()
		return
	}

	fmt.Println(userData)
	var jwtToken = createToken(&userData)

	c.JSON(200, gin.H{
		"data":    userData,
		"token":   jwtToken,
		"message": "Berhasil Login",
	})
}

func createToken(user *models.User) string {
	jwtToken := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"user_id":    user.ID,
		"admin_role": user.AdminRole,
		"exp":        time.Now().AddDate(0, 0, 1).Unix(),
		"iat":        time.Now().Unix(),
	})

	tokenString, err := jwtToken.SignedString([]byte(os.Getenv("JWT_SECRET")))

	if err != nil {
		fmt.Println(err)
	}

	return tokenString
}
