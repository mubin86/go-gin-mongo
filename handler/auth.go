package handler

import (
	"fmt"
	"go-mongo/config"

	//	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
	"golang.org/x/crypto/bcrypt"
)

type Login struct {
  Email string `json:"email,omitempty" bson:"email,omitempty"`
	Password string `json:"password,omitempty" bson:"password,omitempty"`
}

type Data struct {
  Email string `json:"email,omitempty" bson:"email,omitempty" binding:"required"`
	Password string `json:"password,omitempty" bson:"password,omitempty" binding:"required"`
}



func LoginUser(c *gin.Context)  {

	db, _ := config.Connect()

	data := new(Data) 
 login := new(Login)
	if err :=c.BindJSON(&data); err != nil {
			c.JSON(422, gin.H{
				"error":   true,
				"message": "invalid request body",
			})
			return
	}
		
		err := db.Collection("user").FindOne(ctx, bson.M{"email": data.Email}).Decode(&login)

		if err != nil {
			c.JSON(404, gin.H{
				"error":   true,
				"message": "please provide an valid email or pass",
			})
			return
		}

		
	  passerr := bcrypt.CompareHashAndPassword([]byte(login.Password), []byte(data.Password))
	if passerr != nil {
		c.JSON(402, gin.H{"error": "Email or password is invalid."})
		return
	}
	fmt.Println("success")
	



	}
