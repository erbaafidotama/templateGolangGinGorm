package main

import (
	"sewaAset/config"
	"sewaAset/middleware"
	"sewaAset/routes"

	"github.com/gin-gonic/gin"
)

func main() {
	//when server.go start, it will be run function InitDB (connecting to database)
	config.InitDB()

	router := gin.New()
	router.Use(middleware.CORSMiddleware())

	router.POST("/login", routes.Login)                                 // /api/v1/sewaAset/login
	router.GET("/users", middleware.IsAuth(), routes.GetUser)           // /api/v1/sewaAset/users
	router.POST("/users", middleware.IsAuth(), routes.PostUser)         // /api/v1/sewaAset/users
	router.PUT("/users/:id", middleware.IsAuth(), routes.UpdateUser)    // /api/v1/sewaAset/users/:id
	router.DELETE("/users/:id", middleware.IsAuth(), routes.DeleteUser) // /api/v1/sewaAset/users/:id

	// router.Run(":8080") // if you want to run on port 8080
	router.Run()
}
