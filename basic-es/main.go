package main

import "github.com/gin-gonic/gin"

func PingController(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "pong",
	})
}

func main() {
	router := gin.Default()

	router.SetTrustedProxies([]string{"127.0.0.1"})
	router.Use(gin.Logger())
	router.Use(gin.Recovery())

	router.GET("/ping", PingController)

	router.Run(":8000")
}
