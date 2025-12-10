package humans

import "time"

type Human interface {
	GetName() string
	GetSurname() string
	GetAge() uint8
	GetGender() Gender
	IsAlive() bool
	Die()
}

type human struct {
	Name    string
	Surname string
	Age     uint8
	Gender  Gender
	Alive   bool
	BornAt  time.Time
	DiedAt  time.Time

	Hunger int
}

func New(name string, surname string, age uint8, gender Gender) Human {
	return &human{
		Name:    name,
		Surname: surname,
		Age:     age,
		Gender:  gender,
		Alive:   true,
		Hunger:  0,
	}
}

func (h *human) GetName() string {
	return h.Name
}

func (h *human) GetSurname() string {
	return h.Surname
}

func (h *human) GetAge() uint8 {
	return h.Age
}

func (h *human) GetGender() Gender {
	return h.Gender
}

func (h *human) IsAlive() bool {
	return h.Alive
}

func (h *human) Die() {
	h.Hunger = 0
	h.Alive = false
}
