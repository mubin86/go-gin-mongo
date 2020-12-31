package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type User struct {
	ID   string `json:"id"`
	Name string `json:"name"`
	Age  int    `json:"age"`
}

var Users []User

func GetUsers() gin.HandlerFunc {

	return func(c *gin.Context) {
		c.JSON(http.StatusOK, Users)
	}
}

func CreateUser() gin.HandlerFunc {

	return func(c *gin.Context) {
		var reqBody User

		if err := c.ShouldBindJSON(&reqBody); err != nil {
			c.JSON(422, gin.H{
				"error":   true,
				"message": "invalid request body",
			})
			return
		}

		Users = append(Users, reqBody)
		c.JSON(200, gin.H{
			"message": "successfully added User",
		})
	}
}
