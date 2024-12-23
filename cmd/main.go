package main

import (
	"github.com/gin-gonic/gin"
	"golang-hexagonal-architechture-example/internal/adapter/http"
	"golang-hexagonal-architechture-example/pkg/db"
)

func main() {
	// Database initialization
	database := db.NewPostgresDatabase()

	// Initialize Gin router
	router := gin.Default()

	// Register HTTP routes
	http.RegisterRoutes(router, database)

	// Start the server
	router.Run(":8080")
}
