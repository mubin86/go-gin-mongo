package main

import (
	"fmt"
	
	"go-mongo/handler"
	"github.com/gin-gonic/gin"

	"log"
//	"time"
)


func main() {
	fmt.Println("hello world")

	r := gin.Default()

	r.GET("/ping", handler.PingGet())

	userRoutes := r.Group("/user")
	{
		userRoutes.GET("/", handler.GetUsers())
		userRoutes.POST("/add", handler.CreateUser())
		userRoutes.PUT("/update/:id", handler.EditUser())
		userRoutes.DELETE("/delete/:id", handler.DeleteUser())

	}

	productRoutes := r.Group("/product")
	{
		productRoutes.POST("/add",handler.CreateProduct)
		productRoutes.GET("/products",handler.GetProducts)
	}


	if err := r.Run(":5000"); err != nil {
		log.Fatal(err.Error())
	}
}
