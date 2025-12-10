package prng

import (
	"crypto/sha256"
	"sync"
)

type Uint64 interface {
	Generate(uint64) uint64
}

type uint64Generator struct {
	mu             *sync.Mutex
	computedValues map[uint64]uint64
	seed           []byte
}

func NewUint64(seed string) Uint64 {
	h := sha256.New()
	h.Write([]byte(seed))
	seedBytes := h.Sum(nil)

	computedValues := make(map[uint64]uint64)

	mu := &sync.Mutex{}

	zeroValue := uint64(seedBytes[0])

	for i, v := range seedBytes[1:] {
		zeroValue += uint64(v) * uint64(i)
	}

	computedValues[0] = zeroValue

	fistValue := uint64(seedBytes[0])

	for _, v := range seedBytes[1:] {
		fistValue += zeroValue * uint64(v)
	}

	computedValues[1] = fistValue

	return &uint64Generator{seed: seedBytes, computedValues: computedValues, mu: mu}
}

func (g *uint64Generator) Generate(pos uint64) uint64 {
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

	for i := uint64(2); i <= pos; i++ {
		increment := g.seed[xPrev%32]
		multiply := g.seed[(xPrev+xCurr+1)%32]
		mod := uint64(18446744073709551615) / xPrev * xCurr

		xNext := (uint64(multiply)*xCurr + uint64(increment) + i) % mod

		g.mu.Lock()
		g.computedValues[i] = xNext
		g.mu.Unlock()

		xPrev, xCurr = xCurr, xNext
	}

	return xCurr
}
