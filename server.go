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

	v1 := router.Group("api/v1") // /api/v1
	{
		sewaAset := v1.Group("/sewaAset") // /api/v1/sewaAset
		{
			sewaAset.POST("/login", routes.Login)                                 // /api/v1/sewaAset/login
			sewaAset.GET("/users", middleware.IsAuth(), routes.GetUser)           // /api/v1/sewaAset/users
			sewaAset.POST("/users", middleware.IsAuth(), routes.PostUser)         // /api/v1/sewaAset/users
			sewaAset.PUT("/users/:id", middleware.IsAuth(), routes.UpdateUser)    // /api/v1/sewaAset/users/:id
			sewaAset.DELETE("/users/:id", middleware.IsAuth(), routes.DeleteUser) // /api/v1/sewaAset/users/:id
		}
	}

	// router.Run(":8080") // if you want to run on port 8080
	router.Run()
}
