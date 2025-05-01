package routes

import (
	"net/http"

	"example.com/rest-api/models"
	"example.com/rest-api/utilities"
	"github.com/gin-gonic/gin"
)

func createNewUser(context *gin.Context) {
	var user models.User
	err := context.ShouldBindJSON(&user)

	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	err = user.Save()

	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"error": "could not create user"})
		return
	}
	context.JSON(http.StatusCreated, gin.H{"message": "user created"})
}

func login(context *gin.Context) {
	var user models.User

	err := context.ShouldBindJSON(&user)

	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Here you would typically check the user's credentials against the database
	err = user.ValidateCredentials()
	if err != nil {	
		context.JSON(http.StatusUnauthorized, gin.H{"message": err})
		return
	}

	token, err := utilities.GenerateToken(user.Email, int64(user.Id))

	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"error": "could not generate token"})
		return
	}

	context.JSON(http.StatusOK, gin.H{"message": "login successful", "token" : token})
}