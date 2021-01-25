package main

import (
	"fmt"
	"go-mongo/handler"
	"go-mongo/middleware"
	"io"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"

	//	gindump "github.com/tpKeeper/gin-dump"

	"log"
	//"time"
)

func setupLogOutput() {

	f, _ := os.Create("gin.log")
	gin.DefaultWriter = io.MultiWriter(f, os.Stdout)
}


func main() {
	fmt.Println("hello world")

	setupLogOutput()

	r := gin.New()


	r.Use(gin.Recovery(), middleware.Logger(), )


	// Set a lower memory limit for multipart forms (default is 32 MiB)
	r.MaxMultipartMemory = 8 << 20 // 8 MiB
	r.StaticFS("/asset", http.Dir("./public")) ///***if i give '/' instead of '/asset' then the '/' will conflict with the next route like '/ping' because Static uses *filename wildcard so it will conflict with the next /ping route                              

	r.GET("/ping", handler.PingGet())


   // r.Use(middleware.TokenVerify())

	userRoutes := r.Group("/user")
	{
		// userRoutes.GET("/", middleware.BasicAuth(),handler.GetUsers())
		userRoutes.POST("/add", middleware.BasicAuth(),handler.CreateUser)
		userRoutes.PUT("/update/:id", middleware.BasicAuth(),handler.UpdateUser)
		// userRoutes.DELETE("/delete/:id", middleware.BasicAuth(),handler.DeleteUser())
		userRoutes.POST("/auth/login",handler.LoginUser)

	}

	productRoutes := r.Group("/product")

	productRoutes.Use(middleware.AdminAuth())
	{
		
		productRoutes.POST("/add",middleware.TokenVerify(),middleware.IsAdmin(),handler.CreateProduct)
		productRoutes.GET("/products",handler.GetProducts)
		productRoutes.GET("/product/:id",handler.SingleProduct)
		productRoutes.PUT("/update/:id",middleware.TokenVerify(),middleware.IsAdmin(),handler.UpdateProduct)
		productRoutes.DELETE("/delete/:id",middleware.TokenVerify(),middleware.IsAdmin(),handler.DeleteProduct)
		productRoutes.GET("/latest/",handler.LatestProducts)
		productRoutes.GET("/price",handler.PriceBasedProducts)
		productRoutes.GET("/title",handler.TitleBasedProduct)

	}


	r.POST("/upload", handler.FileHandler)

	r.NoRoute(func(c *gin.Context) {
		c.String(http.StatusNotFound, "Not Found")
	})

	if err := r.Run(":5000"); err != nil {
		log.Fatal(err.Error())
	}
}
