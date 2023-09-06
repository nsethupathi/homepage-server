package formatter

import (
	"fmt"
	"strconv"
	"strings"
)

func formatDate(isoDate string) string {
	dateMap := map[string]string{
		"01": "January",
		"02": "February",
		"03": "March",
		"04": "April",
		"05": "May",
		"06": "June",
		"07": "July",
		"08": "August",
		"09": "September",
		"10": "October",
		"11": "November",
		"12": "December",
	}

	yearMonthDay := strings.Split(isoDate, "-")
	month := dateMap[yearMonthDay[1]]
	return month + " " + yearMonthDay[2] + ", " + yearMonthDay[0]
}

func formatTime(isoTime string) string {
	timeMap := map[string]string{
		"00": "12",
		"01": "1",
		"02": "2",
		"03": "3",
		"04": "4",
		"05": "5",
		"06": "6",
		"07": "7",
		"08": "8",
		"09": "9",
		"10": "10",
		"11": "11",
		"12": "12",
		"13": "1",
		"14": "2",
		"15": "3",
		"16": "4",
		"17": "5",
		"18": "6",
		"19": "7",
		"20": "8",
		"21": "9",
		"22": "10",
		"23": "11",
	}

	hourMinSec := strings.Split(isoTime, ":")
	hour := hourMinSec[0]
	minutes := hourMinSec[1]
	formattedHour := timeMap[hour]

	var meridian string
	hourNum, err := strconv.Atoi(hour)

	if err != nil {
		fmt.Println("Error during conversion")
		return ""
	}

	if hourNum < 12 {
		meridian = "AM"
	} else {
		meridian = "PM"
	}

	return formattedHour + ":" + minutes + " " + meridian
}
