package humans

import (
	_ "embed"
	"strings"
	"time"

	"github.com/GeekchanskiY/pet-project/pkg/prng"
)

var (
	//go:embed datasets/surnames.data
	surnamesRaw string
	//go:embed datasets/male_names.data
	maleNamesRaw string
	//go:embed datasets/female_names.data
	femaleNamesRaw string
)

type Generator interface {
	New(uint64) Human
}

type humanGenerator struct {
	surnames    []string
	maleNames   []string
	femaleNames []string

	generator prng.Uint64
}

func NewGenerator(generator prng.Uint64) Generator {
	return &humanGenerator{
		surnames:    strings.Split(surnamesRaw, "\n"),
		maleNames:   strings.Split(maleNamesRaw, "\n"),
		femaleNames: strings.Split(femaleNamesRaw, "\n"),

		generator: generator,
	}
}

func (g *humanGenerator) New(number uint64) Human {
	humanCode := g.generator.Generate(number)

	indexNum := int(humanCode % 1000)

	surname := g.surnames[indexNum%len(g.surnames)]

	gender := Male
	if indexNum%10 > 4 {
		gender = Female
	}

	nameSet := g.maleNames
	if gender == Female {
		nameSet = g.femaleNames
	}

	name := nameSet[indexNum%len(nameSet)]

	age := indexNum % 100

	bornAt := time.Now().Add(time.Duration(age)*time.Hour*24*365*-1 +
		(time.Duration(indexNum%24) * time.Hour) + // make birth hour random
		(time.Duration(indexNum%365) * 24 * time.Hour)) // make date of birth random

	return &human{
		name:    name,
		surname: surname,
		age:     uint8(age),
		gender:  gender,
		bornAt:  bornAt,
		diedAt:  time.Time{},
	}
}
