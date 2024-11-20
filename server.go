package main

import (
	"github.com/gin-gonic/gin"
	// "go-server-demo/middleware"
)


func main() {
	router := gin.Default()
	// Apply LoggerMiddleware globally, or enable in the groups
	// router.Use(middleware.LoggerMiddleware()) 
	Routes(router)
	router.Run("localhost:3000")
}