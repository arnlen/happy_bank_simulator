package helpers

import (
	"time"
)

var dateFormat = "01/2006"

func ParseStringToDate(stringDate string) time.Time {
	parsedDate, _ := time.Parse(dateFormat, stringDate)
	return parsedDate
}

func TimeDateToString(timeDate time.Time) string {
	return timeDate.Format(dateFormat)
}

func AddMonthsToDate(startDate time.Time, monthsToAdd int) time.Time {
	newDate := startDate.AddDate(0, monthsToAdd, 0)
	return newDate
}
