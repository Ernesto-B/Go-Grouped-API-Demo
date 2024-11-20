package main

import (
	"github.com/gin-gonic/gin"
	handlers "go-server-demo/handlers"
	"go-server-demo/middleware"
)

func Routes(c *gin.Engine) {
		albumsEndpoint := c.Group("/albums", middleware.LoggerMiddleware())
		{
			albumsEndpoint.GET("", handlers.GetAlbums)
			albumsEndpoint.POST("", handlers.PostAlbum)
			albumsEndpoint.GET("/:id", handlers.GetAlbumbyId)
		}

		nonAlbumEndpoint := c.Group("/other", middleware.FileLoggerMiddleware())
		{
			nonAlbumEndpoint.GET("", handlers.GetHome)
		}
}