package buildings

import (
	"slices"

	"github.com/GeekchanskiY/pet-project/pkg/humans"
)

type Building interface {
	Citizens() []humans.Human
	Populate(h humans.Human)
	Evict(h humans.Human) error
}

type BuildingFactory interface {
	New() Building
}

type building struct {
	humans []humans.Human
}

type buildingFactory struct {
}

func NewFactory() BuildingFactory {
	return &buildingFactory{}
}

func (b *buildingFactory) New() Building {
	return &building{}
}

func (b *building) Citizens() []humans.Human {
	return b.humans
}

func (b *building) Populate(h humans.Human) {
	b.humans = append(b.humans, h)
}

func (b *building) Evict(h humans.Human) error {
	pos := slices.Index(b.humans, h)
	if pos == -1 {
		return ErrHumanNotFound
	}

	b.humans = append(b.humans[:pos], b.humans[pos+1:]...)

	return nil
}
