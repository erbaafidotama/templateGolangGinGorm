package main

import (
	"sewaAset/config"
	"sewaAset/routes"

	"github.com/gin-gonic/gin"
)

func main() {
	//when server.go start, it will be run function InitDB (connecting to database)
	config.InitDB()

	router := gin.Default()

	v1 := router.Group("api/v1") // /api/v1
	{
		sewaAset := v1.Group("/sewaAset") // /api/v1/sewaAset
		{
			sewaAset.GET("/users", routes.GetUser)   // /api/v1/sewaAset/users
			sewaAset.POST("/users", routes.PostUser) // /api/v1/sewaAset/users
		}
	}

	router.Run()
}
