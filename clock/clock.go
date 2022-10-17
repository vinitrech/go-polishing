package clock

import (
	"fmt"
	"strconv"
	"time"
)

// Define the Clock type here.
type Clock struct {
	hour    string
	minutes string
}

func New(h, m int) Clock {

	clock := time.Date(0, 0, 0, h, m, 0, 0, time.UTC)

	return Clock{
		hour:    fmt.Sprintf("%02d", clock.Hour()),
		minutes: fmt.Sprintf("%02d", clock.Minute()),
	}
}

func (c Clock) Add(m int) Clock {
	hours, _ := strconv.Atoi(c.hour)
	minutes, _ := strconv.Atoi(c.minutes)

	clock := time.Date(0, 0, 0, hours, minutes, 0, 0, time.UTC)
	finalTime := clock.Add(time.Minute * (time.Duration(m)))

	c.hour = fmt.Sprintf("%02d", finalTime.Hour())
	c.minutes = fmt.Sprintf("%02d", finalTime.Minute())

	return c
}

func (c Clock) Subtract(m int) Clock {
	hours, _ := strconv.Atoi(c.hour)
	minutes, _ := strconv.Atoi(c.minutes)

	clock := time.Date(0, 0, 0, hours, minutes, 0, 0, time.UTC)
	finalTime := clock.Add(-time.Minute * (time.Duration(m)))

	c.hour = fmt.Sprintf("%02d", finalTime.Hour())
	c.minutes = fmt.Sprintf("%02d", finalTime.Minute())

	return c
}

func (c Clock) String() string {
	return fmt.Sprintf("%s:%s", c.hour, c.minutes)
}
