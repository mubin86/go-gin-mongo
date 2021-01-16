package handler

import (
	"fmt"
	"go-mongo/config"
	"time"

	//	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type User struct {
	ID  primitive.ObjectID `json:"_id,omitempty" bson:"_id,omitempty"`
	Name string `json:"name,omitempty" bson:"name,omitempty"`
	Email string ` idx:"{email},unique" json:"email" bson:"email" binding:"required"`
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
	  	//t := time.Now()
			//fmt.Println(t.Format(time.ANSIC))
			
		fmt.Println(user.Name)

		user.CreatedAt = time.Now()
		user.UpdatedAt = time.Now()
		
		res,err :=db.Collection("user").InsertOne(ctx,user)
		
		if config.Error(c, err) {
			return //exit
		}
		//fmt.Println(res.InsertedID)

	_ = db.Collection("user").FindOne(ctx, bson.M{"_id": res.InsertedID}).Decode(&user)

		c.JSON(200, gin.H{
		"message": "successfully added user",
		"data": user,
		})
	}



// func EditUser() gin.HandlerFunc {

// return func(c *gin.Context){

// 	id :=c.Param("id")

// 	var reqBody User

// 	if err := c.ShouldBindJSON(&reqBody); err != nil {
// 		c.JSON(422, gin.H{
// 			"error":   true,
// 			"message": "invalid request body",
// 		})
// 		return
// 	}

// 	for i, u := range Users {
// 		if u.ID == id{	
// 			Users[i].Name = reqBody.Name
// 			Users[i].Age = reqBody.Age

// 			c.JSON(200, gin.H{
// 				"message": "updated User data",
// 				"data": Users[i],
// 			})
// 			return
// 		}
// 	}
// 	c.JSON(404, gin.H{
// 		"message": "Invalid ID",
// 		"error": true,
// 	})
// }
// }

// 	func DeleteUser() gin.HandlerFunc {

// 		return func(c *gin.Context) {
// 			id :=c.Param("id")
	
// 			for i, u := range Users {
// 				if u.ID == id{	
// 					Users = append(Users[:i], Users[i+1:] ... )
		
// 					c.JSON(200, gin.H{
// 						"message": "delete User data",
// 						"data": "finished",
// 					})
// 					return
// 				}
// 			}

// 	c.JSON(404, gin.H{
// 		"error": true,
// 		"message": "Invalid User ID",
// 	})
		
// }

// }