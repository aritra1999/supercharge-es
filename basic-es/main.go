package main

import "github.com/gin-gonic/gin"

func main() {
	router := gin.Default()

	ConnectES()

	router.SetTrustedProxies([]string{"127.0.0.1"})
	router.Use(gin.Logger())
	router.Use(gin.Recovery())

	router.GET("/health", HealthController)

	router.Run(":8000")
}
