package orchestrator

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"homepage-server/src/formatter"
	"homepage-server/src/showrepository"
	"log"
	"net/http"
)

func Orchestrate(showRunner *showrepository.TableRunner) func(*gin.Context) {
	return func(c *gin.Context) {
		shows, err := showrepository.GetShows(showRunner)
		if err != nil {
			log.Fatalf("Error retrieving shows: %v", err)
		}
		fmt.Println(shows)

		var formattedShows []*formatter.FormattedShow
		for i := 0; i < len(shows); i++ {
			formattedShows = append(formattedShows, formatter.BuildFormattedShow(shows[i]))
		}
		fmt.Println(formattedShows)

		c.IndentedJSON(http.StatusOK, formattedShows)
	}
}
