package handlers

import (
	"net/http"
	"github.com/gin-gonic/gin"
)

func GetAlbums(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, albums)
}

func PostAlbum(c *gin.Context) {
	var newAlbum album
	err := c.BindJSON(&newAlbum) ; if err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"error": "Invalid input"})
		return
	}
	albums = append(albums, newAlbum)
	c.IndentedJSON(http.StatusCreated, newAlbum)
}

func GetAlbumbyId(c *gin.Context) {
	param := c.Param("id")
	// Search for the id in albums
	for _, val := range albums {
		if val.ID != param {
			continue
		}
		c.IndentedJSON(http.StatusOK, val)
		return
	}
	c.IndentedJSON(http.StatusNotFound, gin.H{"message": "Album with that ID not found"})
}