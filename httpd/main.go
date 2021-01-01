package main

import (
	"fmt"
	"context"
	"go-mongo/handler"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	//"go.mongodb.org/mongo-driver/mongo/readpref"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"log"
	"time"
)

type Product struct {
	ID   primitive.ObjectID `json:"_id,omitempty" bson:"_id,omitempty"`
	Title string            `json:"title,omitempty" bson:"title,omitempty"`
	Description  string     `json:"description,omitempty" bson:"description,omitempty"`
}

var client *mongo.Client

func CreateProduct() gin.HandlerFunc {

	return func(c *gin.Context) {
		var reqBody Product

		if err := c.ShouldBindJSON(&reqBody); err != nil {
			c.JSON(422, gin.H{
				"error":   true,
				"message": "invalid request body",
			})
			return
		}
		collection :=client.Database("demo").Collection("product")
		ctx,_ := context.WithTimeout(context.Background(), 10*time.Second)
		result,_ :=collection.InsertOne(ctx, reqBody)
		


		c.JSON(200, gin.H{
		"message": "successfully added User",
		"data": result,
		})
	}
}

func main() {
	fmt.Println("hello world")
	
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	
	client, _ = mongo.Connect(ctx, options.Client().ApplyURI("mongodb://localhost:27017"))

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
		productRoutes.POST("/add",CreateProduct())

	}


	if err := r.Run(":5000"); err != nil {
		log.Fatal(err.Error())
	}
}
