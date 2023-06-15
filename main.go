package main

import (
	"github.com/gin-gonic/gin"
	"homepage-server/src/aws/service"
)

func main() {

	router := gin.Default()
	router.GET("/song/:songName", service.GetMusicUrl)

	router.Run("localhost:8080")
}
