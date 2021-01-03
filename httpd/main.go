package main

import (
	"fmt"
	"io"
	"os"

	"go-mongo/handler"
	"go-mongo/middleware"

	"github.com/gin-gonic/gin"

	"log"
	//	"time"
)

func setupLogOutput() {

	f, _ := os.Create("gin.log")
	gin.DefaultWriter = io.MultiWriter(f, os.Stdout)
}


func main() {
	fmt.Println("hello world")

	setupLogOutput()

	r := gin.New()

	r.Use(gin.Recovery(), middleware.Logger())

	

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
		productRoutes.GET("/product/:id",handler.SingleProduct)
		productRoutes.PUT("/update/:id",handler.UpdateProduct)
		productRoutes.DELETE("/delete/:id",handler.DeleteProduct)
	}


	if err := r.Run(":5000"); err != nil {
		log.Fatal(err.Error())
	}
}
