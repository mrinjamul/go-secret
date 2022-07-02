package routes

import (
	"time"

	"github.com/gin-gonic/gin"
	"github.com/mrinjamul/go-secret/api/services"
	"github.com/mrinjamul/go-secret/middleware"
)

var (
	StartTime time.Time
	BootTime  time.Duration
)

func InitRoutes(routes *gin.Engine) {
	svc := services.NewServices()
	// Add CORS middleware
	routes.Use(middleware.CORS())
	// Serve the frontend
	// Process the templates at the start so that they don't have to be loaded
	// from the disk again. This makes serving HTML pages very fast.

	// Home Page
	routes.GET("/", func(ctx *gin.Context) {
		svc.ViewService().Index(ctx)
	})
	// New Message Page
	routes.POST("/new", func(ctx *gin.Context) {
		svc.ViewService().AddMessage(ctx)
	})
	routes.GET("/new", func(ctx *gin.Context) {
		svc.ViewService().NotFound(ctx)
	})
	// Message Page
	routes.GET("/:hash", func(ctx *gin.Context) {
		svc.ViewService().ShowMessage(ctx)
	})

	// Add 404 page
	routes.NoRoute(func(ctx *gin.Context) {
		svc.ViewService().NotFound(ctx)
	})

	// api routes group
	api := routes.Group("/api")
	// api.Use(middleware.CORSMiddleware())
	{
		// health check
		api.GET("/health", func(ctx *gin.Context) {
			svc.HealthCheckService().HealthCheck(ctx, StartTime, BootTime)
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
