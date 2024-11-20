package handlers

import (
	"net/http"
	"github.com/gin-gonic/gin"
	models "go-server-demo/models"
)

func GetAlbums(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, models.Albums)
}

func PostAlbum(c *gin.Context) {
	var newAlbum models.Album
	err := c.BindJSON(&newAlbum) ; if err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"error": "Invalid input"})
		return
	}
	models.Albums = append(models.Albums, newAlbum)
	c.IndentedJSON(http.StatusCreated, newAlbum)
}

func GetAlbumbyId(c *gin.Context) {
	param := c.Param("id")
	// Search for the id in models.Albums
	for _, val := range models.Albums {
		if val.ID != param {
			continue
		}
		c.IndentedJSON(http.StatusOK, val)
		return
	}
	c.IndentedJSON(http.StatusNotFound, gin.H{"message": "Album with that ID not found"})
}