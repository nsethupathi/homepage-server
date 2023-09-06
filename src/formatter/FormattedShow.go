package formatter

import (
	"homepage-server/src/showrepository"
	"strings"
)

type FormattedShow struct {
	date    string
	time    string
	venue   string
	address string
}

func BuildFormattedShow(show *showrepository.ShowDO) *FormattedShow {
	splitDate := strings.Split(show.DATE, "T")
	isoDate := splitDate[0]
	isoTime := splitDate[1]
	return &FormattedShow{
		date:    formatDate(isoDate),
		time:    formatTime(isoTime),
		venue:   show.VENUE,
		address: show.ADDRESS,
	}
}
