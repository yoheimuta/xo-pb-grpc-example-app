package exptime

import "time"

// Clock represents a wrapper of time.Time functions using current time.
type Clock struct{}

// NewClock creates a new Clock.
func NewClock() *Clock {
	return &Clock{}
}

// Now returns the current utc time rounding down to second.
func (c *Clock) Now() time.Time {
	return time.Now().UTC().Truncate(time.Second)
}
