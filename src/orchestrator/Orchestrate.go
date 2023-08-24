package orchestrator

import (
	"github.com/gin-gonic/gin"
	"homepage-server/src/aws"
	"log"
	"net/http"
)

func Orchestrate(c *gin.Context) {
	shows, err := aws.GetShows()
	if err != nil {
		log.Fatalf("Error retrieving shows: %v", err)
	}

	c.IndentedJSON(http.StatusOK, shows)
}
