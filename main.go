package main

import (
	"github.com/gin-gonic/gin"
	"homepage-server/src/orchestrator"
)

func main() {

	router := gin.Default()
	router.GET("/shows", orchestrator.Orchestrate)

	router.Run("localhost:8080")
}
