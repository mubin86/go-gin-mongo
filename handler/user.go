package handler

import (
	"fmt"
	"go-mongo/config"
	"log"
	"time"

	//	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"golang.org/x/crypto/bcrypt"
)

type User struct {
	ID  primitive.ObjectID `json:"_id,omitempty" bson:"_id,omitempty"`
	Name string `json:"name,omitempty" bson:"name,omitempty"`
	Email string ` idx:"{email},unique" json:"email" bson:"email" binding:"required"`
	Password string `json:"password" bson:"password" binding:"required"`
	Phone  int `idx:"{phone},unique" json:"phone" bson:"phone" binding:"required"`
	Address struct {
		ZipCode  int    `json:"zipcode" bson:"zipcode"`
		Country  string  `json:"country" bson:"country"`
	}
	CreatedAt time.Time   `json:"createdAt,omitempty" bson:"createdAt,omitempty"`
	UpdatedAt time.Time   `json:"updatedAt,omitempty" bson:"updatedAt,omitempty"`
	
}

//context
// var ctx = func() context.Context {
// 	return context.Background()
// //	return context.WithTimeout(context.Background(), 10*time.Second)
// }()


func CreateUser(c *gin.Context)  {

	db, _ := config.Connect()

	//var reqBody User 
 user := new(User)
	if err :=c.BindJSON(&user); err != nil {
			c.JSON(422, gin.H{
				"error":   true,
				"message": "invalid request body",
			})
			return
		}

		count, err := db.Collection("user").CountDocuments(ctx, bson.M{"email": user.Email})

		fmt.Println(count)
		if err != nil {
			log.Fatal(err)
			c.JSON(404, gin.H{
				"error":   true,
				"message": "something went wrong",
			})
			return
		}

		if count >= 1 {
			fmt.Println("Documents exist in this collection!")
			c.JSON(404, gin.H{
				"error":   true,
				"message": "email already exist",
			})
			return
	}
		
	  	//t := time.Now()
			//fmt.Println(t.Format(time.ANSIC))
			
		fmt.Println(user.Name)

		hash, err := bcrypt.GenerateFromPassword([]byte(user.Password), 10)
	if err != nil {
		log.Fatal(err)
		return
	}
	user.Password = string(hash)
	fmt.Println(user.Password)

		user.CreatedAt = time.Now()
		user.UpdatedAt = time.Now()
		
		res,err :=db.Collection("user").InsertOne(ctx,user)
		
		if config.Error(c, err) {
			return //exit
		}
		//fmt.Println(res.InsertedID)

	_ = db.Collection("user").FindOne(ctx, bson.M{"_id": res.InsertedID}).Decode(&user)

	user.Password = ""
		c.JSON(200, gin.H{
		"message": "successfully added user",
		"data": user,
		})
	}


	
	func UpdateUser(c *gin.Context)  {
		db, _ := config.Connect()
	
		user := new(User)
		if err :=c.BindJSON(&user); err != nil {
				c.JSON(422, gin.H{
					"error":   true,
					"message": "invalid request body",
				})
				return
			}
	
		id := c.Param("id")
		_id, _ := primitive.ObjectIDFromHex(id)
	
		filter := bson.M{"_id": _id}
	
	
	user.UpdatedAt = time.Now()
	fmt.Println(user.UpdatedAt)
	
		 _,err := db.Collection("user").UpdateOne(ctx, filter, bson.M{"$set": user})
		
		if err != nil {
			c.JSON(404, gin.H{
				"error":   true,
				"message": "something went wrong",
			})
			return
		}
		err2 := db.Collection("user").FindOne(ctx, filter).Decode(&user)
	
	if config.Error(c, err2) { //hndling with global error 
		return //exit
	}
	
	c.JSON(200, gin.H{
		"message": "succesfully updated",
		"data": user,
	})
	
	}