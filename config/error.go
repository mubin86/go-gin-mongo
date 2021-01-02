package config


import (
	//"context"
	"github.com/gin-gonic/gin"

)

//error handler

func Error(c *gin.Context, err error) bool {
	if err != nil {
		c.Error(err)
		c.AbortWithStatusJSON(500, gin.H{"message": err.Error(), "data": ""})
		return true
	}
	return false
}