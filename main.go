package main

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"homepage-server/src/orchestrator"
	"homepage-server/src/showrepository"
)

func main() {

	showRunner := showrepository.BuildRunner("SHOWS")

	router := gin.Default()
	router.Use(cors.Default())
	router.GET("/shows", orchestrator.Orchestrate(showRunner))

	router.Run("localhost:8080")
}
