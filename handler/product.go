package handler

import (
	"context"
	"fmt"
  "os"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type Product struct {
	ID   primitive.ObjectID `json:"_id,omitempty" bson:"_id,omitempty"`
	Title string            `json:"title,omitempty" bson:"title,omitempty"`
	Description  string     `json:"description,omitempty" bson:"description,omitempty"`
}

//context
var ctx = func() context.Context {
	return context.Background()
//	return context.WithTimeout(context.Background(), 10*time.Second)
}()

var dbUserName = os.Getenv("DB_USERNAME")
var dbPassword = os.Getenv("DB_PASSWORD")

//connect database
func connect() (*mongo.Database, error) {

	client, err := mongo.NewClient(options.Client().ApplyURI("mongodb+srv://dbUserName:dbPassword@cluster0.4xaod.mongodb.net/gomongo"))
	if err != nil {
		return nil, err
	}
	
	err = client.Connect(ctx)
	if err != nil {
		return nil, err
	}
	fmt.Println("connect database successfully")
	
	return client.Database("gomongo"), nil
}

func CreateProduct(c *gin.Context)  {

	db, _ := connect()

	//var reqBody Product 
	product := new(Product)
	if err :=c.BindJSON(&product); err != nil {
			c.JSON(422, gin.H{
				"error":   true,
				"message": "invalid request body",
			})
			return
		}
		
		fmt.Println(product.Title)
		
		res,_ :=db.Collection("product").InsertOne(ctx, bson.M{"title": product.Title, "description": product.Description})

		fmt.Println(res.InsertedID)

	_ = db.Collection("product").FindOne(ctx, bson.M{"_id": res.InsertedID}).Decode(&product)

		c.JSON(200, gin.H{
		"message": "successfully added product",
		"data": product,
		})
	}
