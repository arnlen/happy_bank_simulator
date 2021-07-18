package helpers

import (
	"time"
)

func ParseStringToDate(stringDate string) time.Time {
	parsedDate, _ := time.Parse("01/2006", stringDate)
	return parsedDate
}

func AddMonthsToDate(startDate time.Time, monthsToAdd int) time.Time {
	newDate := startDate.AddDate(0, monthsToAdd, 0)
	return newDate
}
