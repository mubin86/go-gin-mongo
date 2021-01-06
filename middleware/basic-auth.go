package middleware

import "github.com/gin-gonic/gin"

func BasicAuth() gin.HandlerFunc{

	return gin.BasicAuth(gin.Accounts{
		"mubin" : "hello",
	})
}


func AdminAuth() gin.HandlerFunc{

	return gin.BasicAuth(gin.Accounts{
		"admin" : "hello new",
	})
}