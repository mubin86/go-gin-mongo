package handler

import (
	"net/http"
	"github.com/google/uuid"
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
		reqBody.ID = uuid.New().String()

		Users = append(Users, reqBody)

		c.JSON(200, gin.H{
		"message": "successfully added User",
		"data": reqBody,
		})
	}
}



func EditUser() gin.HandlerFunc {

return func(c *gin.Context){

	id :=c.Param("id")

	var reqBody User

	if err := c.ShouldBindJSON(&reqBody); err != nil {
		c.JSON(422, gin.H{
			"error":   true,
			"message": "invalid request body",
		})
		return
	}

	for i, u := range Users {
		if u.ID == id{	
			Users[i].Name = reqBody.Name
			Users[i].Age = reqBody.Age

			c.JSON(200, gin.H{
				"message": "updated User data",
				"data": Users[i],
			})
			return
		}
	}
	c.JSON(404, gin.H{
		"message": "Invalid ID",
		"error": true,
	})
}
}

	func DeleteUser() gin.HandlerFunc {

		return func(c *gin.Context) {
			id :=c.Param("id")
	
			for i, u := range Users {
				if u.ID == id{	
					Users = append(Users[:i], Users[i+1:] ... )
		
					c.JSON(200, gin.H{
						"message": "delete User data",
						"data": "finished",
					})
					return
				}
			}

	c.JSON(404, gin.H{
		"error": true,
		"message": "Invalid User ID",
	})
		
}

}