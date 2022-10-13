package main

import (
	"httpReq/controllers"
	"httpReq/models"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()
	models.ConnectDatabase()
	router.GET("/api/product", controllers.Index)
	router.GET("/api/product/:id", controllers.Show)
	router.POST("/api/product", controllers.Create)
	router.PUT("/api/product/:id", controllers.Update)
	router.DELETE("/api/product", controllers.Delete)

	router.Run()

}
