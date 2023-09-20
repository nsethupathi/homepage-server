package formatter

import (
	"homepage-server/src/showrepository"
	"strings"
)

type FormattedShow struct {
	Date    string `json:"date"`
	Time    string `json:"time"`
	Venue   string `json:"venue"`
	Address string `json:"address"`
}

func BuildFormattedShow(show *showrepository.ShowDO) *FormattedShow {
	splitDate := strings.Split(show.DATE, "T")
	isoDate := splitDate[0]
	isoTime := splitDate[1]
	return &FormattedShow{
		Date:    formatDate(isoDate),
		Time:    formatTime(isoTime),
		Venue:   show.VENUE,
		Address: show.ADDRESS,
	}
}
