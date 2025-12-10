package prng

import (
	"crypto/sha256"
	"sync"
)

type Uint8 interface {
	Generate(int64) uint8
}

type uint8Generator struct {
	mu             *sync.Mutex
	computedValues map[int64]uint8
	seed           []byte
}

func NewUint8Generator(seed string) Uint8 {
	h := sha256.New()
	h.Write([]byte(seed))
	seedBytes := h.Sum(nil)

	computedValues := make(map[int64]uint8)

	mu := &sync.Mutex{}

	zeroValue := seedBytes[0]

	for _, v := range seedBytes[1:] {
		zeroValue += v
	}

	computedValues[0] = zeroValue

	fistValue := seedBytes[0]

	for _, v := range seedBytes[1:] {
		fistValue += zeroValue % v
	}

	computedValues[1] = fistValue

	return &uint8Generator{seed: seedBytes, computedValues: computedValues, mu: mu}
}

func (g *uint8Generator) Generate(pos int64) uint8 {
	if pos < 0 {
		pos = -pos
	}

	g.mu.Lock()
	val, exists := g.computedValues[pos]
	g.mu.Unlock()

	if exists {
		return val
	}

	xPrev := g.computedValues[0]
	xCurr := g.computedValues[1]

	for i := int64(2); i <= pos; i++ {
		increment := g.seed[xPrev%32]
		multiply := g.seed[(xPrev+xCurr+1)%32]

		xNext := (uint16(multiply)*uint16(xCurr) + uint16(increment) + uint16(i)) % uint16(256)

		g.mu.Lock()
		g.computedValues[i] = uint8(xNext)
		g.mu.Unlock()

		xPrev, xCurr = xCurr, uint8(xNext)
	}

	return xCurr
}
