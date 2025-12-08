package world

import "github.com/GeekchanskiY/pet-project/pkg/humans"

type World interface {
	Destroy()
}

type world struct {
	people []humans.Human
}

func NewWorld() World {
	return &world{}
}

func (w *world) Destroy() {
	for _, h := range w.people {
		h.Die()
	}
}
