package routes

import (
	"embed"
	"io/fs"
	"log"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/mrinjamul/go-secret/api/services"
)

// ViewsFs for static files
var ViewsFs embed.FS
var StartTime time.Time

func InitRoutes(routes *gin.Engine) {
	svc := services.NewServices()
	// Serve the frontend
	// This will ensure that the files are served correctly
	fsRoot, err := fs.Sub(ViewsFs, "views")
	if err != nil {
		log.Println(err)
	}
	routes.NoRoute(gin.WrapH(http.FileServer(http.FS(fsRoot))))

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
		api.GET("/message", func(ctx *gin.Context) {
			svc.MessageService().GetAll(ctx)
		})
		api.GET("/message/:id", func(ctx *gin.Context) {
			svc.MessageService().Get(ctx)
		})
		api.POST("/message", func(ctx *gin.Context) {
			svc.MessageService().Add(ctx)
		})
		api.PUT("/message/:id", func(ctx *gin.Context) {
			svc.MessageService().Update(ctx)
		})
		api.DELETE("/message", func(ctx *gin.Context) {
			svc.MessageService().Delete(ctx)
		})
		api.GET("/:hash", func(ctx *gin.Context) {
			svc.MessageService().GetAndRead(ctx)
		})
	}
}
