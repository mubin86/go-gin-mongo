package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"go-mongo/handler"
  )

func main(){
	fmt.Println("hello worl")

	r := gin.Default()

	r.GET("/ping", handler.PingGet)

	r.Run() 
}