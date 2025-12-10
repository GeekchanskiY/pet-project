package humans

import (
	"time"

	"github.com/GeekchanskiY/pet-project/pkg/prng"
)

var (
	surnames    = []string{"Smith", "Johnson", "Williams", "Kowalski", "Rodriguez", "Ivanov"}
	maleNames   = []string{"Smith", "John", "William", "Igor", "Dmitry"}
	femaleNames = []string{"Anna", "Anastasia", "Victoria", "Olga"}
)

type Generator interface {
	New(uint64) Human
}

type humanGenerator struct {
	generator prng.Uint64
}

func NewGenerator(generator prng.Uint64) Generator {
	return &humanGenerator{
		generator: generator,
	}
}

func (g *humanGenerator) New(number uint64) Human {
	humanCode := g.generator.Generate(number)

	indexNum := int(humanCode % 1000)

	surname := surnames[indexNum%len(surnames)]

	gender := Male
	if indexNum%10 > 4 {
		gender = Female
	}

	nameSet := maleNames
	if gender == Female {
		nameSet = femaleNames
	}

	name := nameSet[indexNum%len(nameSet)]

	age := indexNum % 100

	bornAt := time.Now().Add(time.Duration(age)*time.Hour*24*365*-1 +
		(time.Duration(indexNum%24) * time.Hour) + // make birth hour random
		(time.Duration(indexNum%365) * 24 * time.Hour)) // make date of birth random

	return &human{
		Name:    name,
		Surname: surname,
		Age:     uint8(age),
		Gender:  gender,
		Alive:   false,
		BornAt:  bornAt,
		DiedAt:  time.Time{},
		Hunger:  100,
	}
}
