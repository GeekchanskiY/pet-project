package humans

import "time"

type Stat interface {
	Name() string
	Value() any
	Change(now time.Time, value any)

	GetLastChange() time.Time
}

type statHistoryLog struct {
	value     any
	timestamp time.Time
}

type stat struct {
	name  string
	value any

	history []*statHistoryLog
}

func NewStat(now time.Time, name string, value any) Stat {
	return &stat{
		name:    name,
		value:   value,
		history: []*statHistoryLog{{timestamp: now, value: value}},
	}
}

func (s *stat) Name() string {
	return s.name
}

func (s *stat) Value() any {
	return s.value
}

func (s *stat) Change(now time.Time, value any) {
	s.history = append(s.history, &statHistoryLog{
		value:     s.value,
		timestamp: now,
	})

	s.value = value
}

func (s *stat) GetLastChange() time.Time {
	return s.history[len(s.history)-1].timestamp
}
