package clock

import "time"

type Clock struct{}

func (c Clock) Now() time.Time {
	return time.Now()
}

var Instance Clock
