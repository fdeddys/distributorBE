package main

import (
	"fmt"

	"github.com/gin-gonic/gin"
	_ "github.com/jinzhu/gorm/dialects/postgres"

	controller "distributorBE/controller"
	_ "distributorBE/database"
)

func main() {

	gin.SetMode(gin.ReleaseMode)

	router := gin.New()
	router.Use(middleware())

	route := router.Group("/api")
	{
		route.GET("/ping", controller.Ping)
	}

	router.Run()
}

func middleware() gin.HandlerFunc {
	return func(c *gin.Context) {

		fmt.Println("------------------ Start Middleware ------------------")

		fmt.Println("------------------ Finish Middleware ------------------")
		c.Next()
		return
	}
}
