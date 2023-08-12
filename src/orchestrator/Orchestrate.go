package orchestrator

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"homepage-server/src/service"
)

func Orchestrate(c *gin.Context) {

	runner := service.Build()

	fmt.Println(result)
}
