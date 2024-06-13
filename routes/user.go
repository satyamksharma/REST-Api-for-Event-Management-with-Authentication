package routes

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/satyamksharma/REST-Api-for-Event-Management-with-Authentication.git/models"
)

func signup(context *gin.Context) {
	var user models.User
	err := context.ShouldBindJSON(&user)

	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": "Could not parse request data!"})
		return
	}

	err = user.Save()
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "Could not create user, Please try later!"})
		return
	}

	context.JSON(http.StatusCreated, gin.H{"message": "User Created!", "user": user})
}