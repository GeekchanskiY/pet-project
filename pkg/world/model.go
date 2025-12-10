package world

import (
	"fmt"

	"github.com/GeekchanskiY/pet-project/pkg/humans"
	"github.com/GeekchanskiY/pet-project/pkg/prng"
)

type World interface {
	Live()
	Destroy()
}

type world struct {
	seed string
	prng prng.Uint64

	clock uint64

	people []humans.Human
}

func NewWorld(seed string) World {
	w := &world{
		seed: seed,
		prng: prng.NewUint64(seed),
	}

	w.init()

	return w
}

func (w *world) init() {
	humanGenerator := humans.NewGenerator(w.prng)

	people := make([]humans.Human, 100)
	for i := range uint64(100) {
		people[i] = humanGenerator.New(i)
	}

	w.people = people
}

// Live simulates world's activities.
func (w *world) Live() {
	// Do all world's actions here
	for _, h := range w.people {
		fmt.Println(h.GetName(), h.GetSurname(), h.GetAge(), h.GetGender(), h.IsAlive())
	}

	w.clock++
}

func (w *world) Destroy() {
	for _, h := range w.people {
		h.Die()
	}

	w.clock++
}
