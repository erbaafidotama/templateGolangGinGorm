package middleware

import (
	"fmt"
	"os"
	"strings"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)

func IsAuth() gin.HandlerFunc {
	fmt.Println("AUTH")
	return checkJWT(false)
}

func checkJWT(middlewareAdmin bool) gin.HandlerFunc {
	return func(c *gin.Context) {
		authHeader := c.Request.Header.Get("Authorization")
		bearerToken := strings.Split(authHeader, " ")

		if len(bearerToken) == 2 {
			token, err := jwt.Parse(bearerToken[1], func(token *jwt.Token) (interface{}, error) {
				// Don't forget to validate the alg is what you expect:
				if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
					return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
				}

				// hmacSampleSecret is a []byte containing your secret, e.g. []byte("my_secret_key")
				return []byte(os.Getenv("JWT_SECRET")), nil
			})

			if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
				adminRole := bool(claims["admin_role"].(bool))
				c.Set("jwt_user_id", claims["user_id"])
				// c.Set("jwt_isAdmin", claims["user_role"])

				if middlewareAdmin == true && adminRole == false {
					c.JSON(422, gin.H{"msg": "Only Admin Allowed"})
					c.Abort()
					return
				}
			} else {
				c.JSON(422, gin.H{"msg": "Invalid Token", "error": err})
				c.Abort()
				return
			}
		} else {
			c.JSON(422, gin.H{"msg": "Authorization Token not provided"})
			c.Abort()
			return
		}
	}
}

func CORSMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
		c.Writer.Header().Set("Access-Control-Allow-Credentials", "true")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, accept, origin, Cache-Control, X-Requested-With")
		c.Writer.Header().Set("Access-Control-Allow-Methods", "POST, OPTIONS, GET, PUT")

		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(204)
			return
		}

		c.Next()
	}
}
