package main

import (
	"log"
	"os"
	"time"

	"github.com/mrinjamul/go-secret/api/routes"

	"github.com/gin-gonic/gin"
)

func main() {
	// Get port from env
	port := ":3000"
	_, present := os.LookupEnv("PORT")
	if present {
		port = ":" + os.Getenv("PORT")

	}

	// Set the router as the default one shipped with Gin
	server := gin.Default()
	// Initialize the routes

	routes.StartTime = time.Now()
	routes.InitRoutes(server)
	// Start and run the server
	log.Fatal(server.Run(port))
}
