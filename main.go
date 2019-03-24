package main

import (
	"fmt"

	controller "distributorBE/controller"

	"github.com/gin-gonic/gin"
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
