package main

import (
	"net/http"
	"github.com/gin-gonic/gin"
)

type album struct {
	ID string 	`json:"id"`
	Title string `json:"title"`
	Artist string `json:"artist"`
	Price float32 `json:"price"`
}

var albums = []album{
	{ID: "1", Title: "Blue Train", Artist: "John Coltrane", Price: 56.99},
    {ID: "2", Title: "Jeru", Artist: "Gerry Mulligan", Price: 17.99},
    {ID: "3", Title: "Sarah Vaughan and Clifford Brown", Artist: "Sarah Vaughan", Price: 39.99},
}

func getAlbums(c *gin.Context) {
	// IndentedJSON auto sends the response
	c.IndentedJSON(http.StatusOK, albums)
}

func postAlbum(c *gin.Context) {
	var newAlbum album
	// BindJSON auto parses JSON object from request
	if err := c.BindJSON(&newAlbum); err != nil {
		return
	}

	albums = append(albums, newAlbum)
	c.IndentedJSON(http.StatusCreated, newAlbum)
}

func getAlbumbyId(c *gin.Context) {
	id := c.Param("id")

	// range returns an index and value for each iteration
	for _, val := range albums { 	
		if val.ID == id {
			c.IndentedJSON(http.StatusOK, val)
			return
		}
	}
	c.IndentedJSON(http.StatusNotFound, gin.H{"message": "album not found"})
}

func main() {
	router := gin.Default()
	router.GET("/albums", getAlbums)
	router.POST("/albums", postAlbum)
	router.GET("/albums/:id", getAlbumbyId)

	router.Run("localhost:3000")
}