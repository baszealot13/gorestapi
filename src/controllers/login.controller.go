package controllers

import (
	"gorestapi/src/entities"
	"gorestapi/src/services"
	"net/http"

	"github.com/gin-gonic/gin"
)

// LogIn controller ... Login
func LogIn(c *gin.Context) {

	var user entities.User
	username := c.PostForm("username")
	password := c.PostForm("password")

	err := services.LogIn(&user, username, password)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
			"message": "Unauthorize",
		})
	}

	c.JSON(http.StatusOK, gin.H{
		"token": services.GenerateToken(user.UserID, user.Username, user.UserIsAdmin),
	})
}
