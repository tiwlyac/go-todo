package main

import (
	"github.com/gin-gonic/gin"
	"go-todo/service"
	"go-todo/database"
)

const port = ":1234"

func main() {
	database.Create()
	router := gin.Default()
	v1 := router.Group("/api/todos") 
	{
		v1.GET("/", service.GetHandler)
		v1.GET("/:id", service.GetByIDHandler)
		v1.POST("/", service.PostHandler)
		v1.PUT("/:id", service.PutByIDHandler)
		v1.DELETE("/:id", service.DeleteByIDHandler)
	}
	router.Run(port)
}


