package main

import (
	"embed"
	"html/template"
	"io/fs"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/mrinjamul/go-secret/api/routes"

	"github.com/gin-gonic/gin"
)

//go:embed views/*
var files embed.FS

var (
	// startTime is the time when the server starts
	startTime time.Time = time.Now()
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
	templ := template.Must(template.New("").ParseFS(files, "views/base/*.html", "views/pages/*.html"))
	server.SetHTMLTemplate(templ)
	static, err := fs.Sub(files, "views/static")
	if err != nil {
		panic(err)
	}
	media, err := fs.Sub(files, "views/media")
	if err != nil {
		panic(err)
	}
	server.StaticFS("/static", http.FS(static))
	server.StaticFS("/media", http.FS(media))

	// Initialize the routes
	routes.StartTime = startTime
	routes.InitRoutes(server)
	routes.BootTime = time.Since(startTime)
	// Start and run the server
	log.Fatal(server.Run(port))
}
