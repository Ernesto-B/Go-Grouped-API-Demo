package main

import (
	"github.com/gin-gonic/gin"
	handlers "go-server-demo/handlers"
)

func Routes(c *gin.Engine) {
		albumsEndpoint := c.Group("/albums")
		{
			albumsEndpoint.GET("", handlers.GetAlbums)
			albumsEndpoint.POST("", handlers.PostAlbum)
			albumsEndpoint.GET("/:id", handlers.GetAlbumbyId)
		}
}