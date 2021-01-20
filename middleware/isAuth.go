package middleware

import (
	"fmt"
	"strings"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)


func TokenVerify() gin.HandlerFunc {

	return func(c *gin.Context) {
		authHeader := c.Request.Header.Get("x-auth-token")  // 	Header["x-auth-token"][0] // for [0] it will return the string otherwise it will return the array
		fmt.Println(authHeader)
		bearerToken := strings.Split(authHeader, " ")
	//	bearerToken :=authHeader.Split(" ")

		if len(bearerToken) == 2{

			authToken := bearerToken[1]

		token,error :=	jwt.Parse(authToken, func(token *jwt.Token) (interface{}, error){
			//	spew.Dump(token) // get the deatils of the token in seperate props and got the 
			if _,ok :=token.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil,fmt.Errorf("there is an error")
			}
			return []byte("secret"), nil

		})
		if error != nil {
			c.JSON(400, gin.H{
				"error":   true,
				"message": "please provide an valid email or pass or valid token",
			})
			return
		}
	//	spew.Dump(token) //get info auth or not here token.valid = true

	if _, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		c.Next()
		
	} else {
	c.JSON(400, gin.H{
		"error":   true,
		"message": "please try again",
	})
	return
}
			
		}	

	}
}