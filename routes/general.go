package routes

import (
    "fmt"
    "os"
    "sewaAset/config"
    "sewaAset/models"
    "strconv"
    "time"

    "github.com/dgrijalva/jwt-go"
    "github.com/gin-gonic/gin"
    "golang.org/x/crypto/bcrypt"
)

func Login(c *gin.Context) {
    var userData models.User
    var req struct {
        Username string `json:"username"`
        Password string `json:"password"`
    }
    if err := c.ShouldBindJSON(&req); err != nil {
        c.JSON(400, gin.H{
            "status":  "error",
            "message": "invalid request",
        })
        c.Abort()
        return
    }

    if err := config.DB.Where("username = ?", req.Username).First(&userData).Error; err != nil {
        c.JSON(404, gin.H{
            "status":  "error",
            "message": "record not found",
        })
        c.Abort()
        return
    }

    if err := bcrypt.CompareHashAndPassword([]byte(userData.Password), []byte(req.Password)); err != nil {
        c.JSON(401, gin.H{
            "status":  "error",
            "message": "invalid credentials",
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
	// Read expiration from env, default to 30 minutes if not set
	expirationMinutes := 30
	if envExp := os.Getenv("JWT_EXPIRATION_MINUTES"); envExp != "" {
		if minutes, err := strconv.Atoi(envExp); err == nil {
			expirationMinutes = minutes
		}
	}

	jwtToken := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"user_id":    user.ID,
		"admin_role": user.AdminRole,
		"exp":        time.Now().Add(time.Duration(expirationMinutes) * time.Minute).Unix(),
		"iat":        time.Now().Unix(),
	})

	tokenString, err := jwtToken.SignedString([]byte(os.Getenv("JWT_SECRET")))

	if err != nil {
		fmt.Println(err)
	}

	return tokenString
}
