package main

import (
	"fmt"
	"go-mongo/handler"

	"github.com/gin-gonic/gin"
)

func main() {
	fmt.Println("hello world")

	r := gin.Default()

	r.GET("/ping", handler.PingGet())

	userRoutes := r.Group("/users")
	{
		userRoutes.GET("/", handler.GetUsers())
		userRoutes.POST("/add", handler.CreateUser())

	}
	if err := r.Run(":5000"); err != nil {
		fmt.Println("error occured")
	}
}
