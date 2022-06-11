package routes

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/mrinjamul/go-secret/api/services"
)

var StartTime time.Time

func InitRoutes(routes *gin.Engine) {
	svc := services.NewServices()
	// Serve the frontend
	// Process the templates at the start so that they don't have to be loaded
	// from the disk again. This makes serving HTML pages very fast.
	routes.LoadHTMLGlob("views/**/*")
	// routes.Static("/static", "static")
	routes.GET("/", func(c *gin.Context) {
		c.HTML(http.StatusOK, "index.html", gin.H{
			"title": "Secret — Home",
		})
	})
	// serve static pages under static folder
	routes.GET("/static/*filepath", func(c *gin.Context) {
		c.File("views/static/" + c.Param("filepath"))
	})
	routes.GET("/media/*filepath", func(c *gin.Context) {
		c.File("views/media/" + c.Param("filepath"))
	})
	routes.POST("/new", func(ctx *gin.Context) {
		svc.MessageService().AddMessage(ctx)
	})
	routes.GET("/new", func(c *gin.Context) {
		c.HTML(http.StatusOK, "404.html", gin.H{
			"title": "Secret — Error",
		})
	})

	routes.GET("/:hash", func(ctx *gin.Context) {
		svc.MessageService().ShowMessage(ctx)
	})

	// Add 404 page
	routes.NoRoute(func(c *gin.Context) {
		c.HTML(http.StatusNotFound, "404.html", gin.H{
			"title": "Secret — 404",
		})
	})

	// api routes group
	api := routes.Group("/api")
	// api.Use(middleware.CORSMiddleware())
	{
		// health check
		api.GET("/health", func(c *gin.Context) {
			c.Header("Cache-Control", "no-cache, no-store, must-revalidate")
			c.JSON(http.StatusOK, gin.H{
				"uptime": time.Since(StartTime).String(),
				"status": "ok",
				"time":   time.Now().Format(time.RFC3339),
			})

		})
		// Get messages by unique hash
		api.GET("/:hash", func(ctx *gin.Context) {
			svc.MessageService().GetAndRead(ctx)
		})
		// Add a new message
		api.POST("/message", func(ctx *gin.Context) {
			svc.MessageService().Add(ctx)
		})
		// {
		// 	// Get all messages
		// 	api.GET("/message", func(ctx *gin.Context) {
		// 		svc.MessageService().GetAll(ctx)
		// 	})
		// 	// Get a message
		// 	api.GET("/message/:id", func(ctx *gin.Context) {
		// 		svc.MessageService().Get(ctx)
		// 	})
		// 	// Update a message
		// 	api.PUT("/message/:id", func(ctx *gin.Context) {
		// 		svc.MessageService().Update(ctx)
		// 	})
		// 	// Delete a message
		// 	api.DELETE("/message", func(ctx *gin.Context) {
		// 		svc.MessageService().Delete(ctx)
		// 	})
		// }
	}
}
