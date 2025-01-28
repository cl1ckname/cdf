package mock

import "time"

type Clock struct {
	Time time.Time
}

func (c Clock) Now() time.Time {
	return c.Time
}
