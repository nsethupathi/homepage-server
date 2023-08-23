package orchestrator

import (
	"github.com/gin-gonic/gin"
	"homepage-server/src/AwsService"
	"log"
	"net/http"
)

func Orchestrate(c *gin.Context) {

	runner := AwsService.BuildRunner()
	shows, err := runner.GetShows()
	if err != nil {
		log.Fatalf("Error retrieving shows: %v", err)
	}

	c.IndentedJSON(http.StatusOK, shows)
}
