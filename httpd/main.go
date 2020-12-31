package main

import (
	"fmt"
	"go-mongo/handler"
	"github.com/gin-gonic/gin"
	"log"
)

func main() {
	fmt.Println("hello world")

	r := gin.Default()

	r.GET("/ping", handler.PingGet())

	userRoutes := r.Group("/users")
	{
		userRoutes.GET("/", handler.GetUsers())
		userRoutes.POST("/add", handler.CreateUser())
		userRoutes.PUT("/update/:id", handler.EditUser())

	}
	if err := r.Run(":5000"); err != nil {
		log.Fatal(err.Error())
	}
}
