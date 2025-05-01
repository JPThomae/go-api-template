package routes

import (
	"example.com/rest-api/api-test/middlewares"
	"github.com/gin-gonic/gin"
)

func RegisterRoutes(server *gin.Engine) {
	// : allows you to pass the id as a parameter in the URL
	server.GET("/events/:id", getEventById)
	server.GET("/events", getEvents)
	server.POST("/signup", createNewUser)
	server.POST("/login", login)

	authenticated := server.Group("/")
	authenticated.Use(middlewares.Authenticate)

	authenticated.POST("/events", createEvent)
	authenticated.PUT("/events/:id", updateEvent)
	authenticated.DELETE("/events/:id", deleteEvent)
	authenticated.POST("/events/:id/register", registerForEvent)
	authenticated.DELETE("/events/:id/unregister", unregisterFromEvent)
}