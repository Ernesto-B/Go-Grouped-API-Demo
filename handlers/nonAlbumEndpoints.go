package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)


func GetHome(c *gin.Context) {
	c.String(http.StatusOK, "You have accessed the homepage, after logging")
}