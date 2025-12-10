package main

import (
	"github.com/GeekchanskiY/pet-project/pkg/world"
)

func main() {
	newWorld := world.NewWorld("geekchanskiy")
	newWorld.Live()
}
