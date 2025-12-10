package humans

import (
	"reflect"
	"time"
)

type Human interface {
	GetName() string
	GetSurname() string
	GetAge() uint8
	GetGender() Gender
	IsAlive() bool
	Die(now time.Time)

	Live()
}

type human struct {
	name    string
	surname string
	age     uint8
	gender  Gender
	bornAt  time.Time
	diedAt  time.Time

	stats []Stat
}

func New(name string, surname string, age uint8, gender Gender) Human {
	return &human{
		name:    name,
		surname: surname,
		age:     age,
		gender:  gender,
		stats:   make([]Stat, 0),
	}
}

func (h *human) GetName() string {
	return h.name
}

func (h *human) GetSurname() string {
	return h.surname
}

func (h *human) GetAge() uint8 {
	return h.age
}

func (h *human) GetGender() Gender {
	return h.gender
}

func (h *human) IsAlive() bool {
	return h.diedAt.IsZero()
}

func (h *human) Die(now time.Time) {
	h.diedAt = now

	for _, stat := range h.stats {
		stat.Change(now, reflect.Zero(reflect.TypeOf(stat.Value())))
	}
}
