package main

// Imorted gin, install using 'go get github.com/gin-gonic/gin'
import (
	"example.com/rest-api/routes"
	"example.com/rest-api/db"	
	"github.com/gin-gonic/gin"
)

func main() {
	// Initialize the database connection
	db.InitDB()
	// Create a new Gin router
	// gin.Default() creates a router with default middleware: logger and recovery (crash-free) middleware
	httpServer := gin.Default()

	// Sending my server to all my routing functions
	routes.RegisterRoutes(httpServer)

	// listen and serve on localhost:8080
	httpServer.Run(":8080")
}
