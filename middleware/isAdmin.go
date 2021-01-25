package middleware

import (
	"net/http"
	"strings"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)

func IsAdmin() gin.HandlerFunc {

	return func(c *gin.Context) {

		authHeader := c.Request.Header.Get("x-auth-token")
		bearerToken := strings.Split(authHeader, " ")[1]

		claims := jwt.MapClaims{}
		_, err := jwt.ParseWithClaims(bearerToken, claims, func(*jwt.Token) (interface{}, error){
			return []byte("secret"), nil //******get secret from env must
		})
		if err != nil {
			 c.String(http.StatusInternalServerError, "unable to parse token")
		}

		if !claims["authorized"].(bool) {
	//	c.String(http.StatusForbidden, "not authorized")
		//	c.String(http.StatusNotFound, "Not Found")
		c.JSON(403, gin.H{
			"error":   true,
			"message": "not authorized",
		})
		c.Abort()
		return	
	}

		c.Next()

	}

}