package main

import (
	"context"
	"log"

	"github.com/gin-gonic/gin"
)

func HealthController(c *gin.Context) {
	res, err := ES.Ping().Do(context.Background())
	if err != nil {
		log.Fatalf("Error getting response: %s", err)
	}

	c.JSON(200, gin.H{
		"es": res,
	})
}
