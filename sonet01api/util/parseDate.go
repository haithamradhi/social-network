package util

import "time"

func ParseDate(dateStr string) time.Time {
	layout := "2006-01-02" // Adjust the layout according to your date format
	date, err := time.Parse(layout, dateStr)
	if err != nil {
		return time.Time{} // Return zero value of time.Time in case of error
	}
	return date
}
