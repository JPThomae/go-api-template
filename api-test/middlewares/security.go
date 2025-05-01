package middlewares

import (
	"net/http"
	"github.com/gin-gonic/gin"
	"example.com/rest-api/utilities"
)

func Authenticate(context *gin.Context) {
	token := context.GetHeader("Authorization")

	if token == "" {
		context.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "authorization token required"})
		return
	}

	userId, err := utilities.VerifyToken(token)

	if err != nil {
		context.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "invalid or expired token"})
		return
	}

	context.Set("userId", userId)
	context.Next()
}

func IdentitfyUser(context *gin.Context) {
	userId, exists := context.Get("userId")
	if !exists {
		context.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "user not authenticated"})
		return
	}

	context.Set("userId", userId)
	context.Next()
}