package world

import (
	"fmt"
	"time"

	"github.com/GeekchanskiY/pet-project/pkg/clock"
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

	clock clock.Clock

	people []humans.Human
}

func NewWorld(seed string) World {
	w := &world{
		seed:  seed,
		prng:  prng.NewUint64(seed),
		clock: clock.New(time.Date(2000, 1, 1, 0, 0, 0, 0, time.UTC)),
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

// Live simulates world's activities. Tick interval = 1 hour
func (w *world) Live() {
	// Do all world's actions here
	for _, h := range w.people {
		fmt.Println(h.GetName(), h.GetSurname(), h.GetAge(), h.GetGender(), h.IsAlive())
	}

	w.clock.Tick()
}

func (w *world) Destroy() {
	for _, h := range w.people {
		h.Die(w.clock.Now())
	}

	w.clock.Tick()
}
