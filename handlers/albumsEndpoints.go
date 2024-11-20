package handlers

import (
	models "go-server-demo/models"
	"net/http"
	"sort"

	"github.com/gin-gonic/gin"
)

func GetAlbums(c *gin.Context) {
	// Make a copy so we dont modify the original models.Albums
	albumsCopy := make([]models.Album, len(models.Albums))
	copy(albumsCopy, models.Albums)

	sortby := c.Query("sort")

	switch sortby {
	case "title":
		sort.Slice(albumsCopy, func(i, j int) bool {
			return albumsCopy[i].Title < albumsCopy[j].Title
		})
	case "artist":
		sort.Slice(albumsCopy, func(i, j int) bool {
			return albumsCopy[i].Artist < albumsCopy[j].Artist
		})
	case "price":
		sort.Slice(albumsCopy, func(i, j int) bool {
			return albumsCopy[i].Price < albumsCopy[j].Price
		})
	default:
	}

	c.IndentedJSON(http.StatusOK, albumsCopy)
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