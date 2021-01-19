package handler

import (
	"fmt"
	"go-mongo/config"
	"log"
	"os"
	"time"

	"github.com/dgrijalva/jwt-go"
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

func createToken(userName string) (string,error) {
	jwtSecret := os.Getenv("JWT_SECRET")
	fmt.Println(jwtSecret)

claims := jwt.MapClaims{}
claims["authorized"] = true
claims["user_id"] = userName
claims["exp"] = time.Now().Add(time.Minute * 15).Unix()
token:= jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

tokenString,err := token.SignedString([]byte(jwtSecret))

if err != nil {
	log.Fatal("unable to generate the token")
	return "", err
}

return tokenString, nil
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
	token, tokenerr :=createToken(login.Email)
	if tokenerr != nil {
		c.JSON(402, gin.H{"error": "unable to generate the token"})
		return
	}

	//c.Response().Header.Set("auth-token", token)

	// c.Request.Response().Header.Set("auth-token", token)
	c.Header("auth-token", token)

	login.Password = ""
	c.JSON(200, gin.H{
	"message": "successfully logged in user",
	"data": login,
	})


	}
