package booking

import (
	"fmt"
	"time"
)

// Schedule returns a time.Time from a string containing a date.
func Schedule(date string) time.Time {
	const layout = "1/02/2006 15:04:05"
	parsedTime, _ := time.Parse(layout, date)

	return parsedTime
}

// HasPassed returns whether a date has passed.
func HasPassed(date string) bool {
	const layout = "January _2, 2006 15:04:05"
	parsedTime, err := time.Parse(layout, date)

	if err != nil {
		fmt.Println("Error: ", err, " - ParsedTime: ", parsedTime)
		return false
	}

	return time.Now().After(parsedTime)
}

// IsAfternoonAppointment returns whether a time is in the afternoon.
func IsAfternoonAppointment(date string) bool {
	const layout = "Monday, January _2, 2006 15:04:05"
	parsedTime, err := time.Parse(layout, date)

	if err != nil {
		fmt.Println("Error: ", err, " - ParsedTime: ", parsedTime)
		return false
	}

	hour := parsedTime.Hour()

	return hour >= 12 && hour < 18
}

// Description returns a formatted string of the appointment time.
func Description(date string) string {
	const layout = "1/2/2006 15:04:05"

	parsedTime, err := time.Parse(layout, date)

	if err != nil {
		fmt.Println("Error: ", err, " - ParsedTime: ", parsedTime)
		return ""
	}

	return fmt.Sprintf("You have an appointment on %s, %s %d, %d, at %d:%d.", parsedTime.Weekday(), parsedTime.Month(), parsedTime.Day(), parsedTime.Year(), parsedTime.Hour(), parsedTime.Minute())
}

// AnniversaryDate returns a Time with this year's anniversary.
func AnniversaryDate() time.Time {
	return time.Date(time.Now().Year(), time.September, 15, 0, 0, 0, 0, time.UTC)
}
