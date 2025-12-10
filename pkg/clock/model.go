package clock

import "time"

type Clock interface {
	Now() time.Time
	Tick()
	StartTime() time.Time
}

type clock struct {
	startTime    time.Time
	currentTime  time.Time
	tickInterval time.Duration
}

func New(startTime time.Time) Clock {
	return &clock{
		startTime:    startTime,
		currentTime:  startTime,
		tickInterval: time.Hour,
	}
}

func (c *clock) Now() time.Time {
	return c.currentTime
}
func (c *clock) Tick() {
	c.currentTime = c.currentTime.Add(c.tickInterval)
}

func (c *clock) StartTime() time.Time {
	return c.startTime
}
