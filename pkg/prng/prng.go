package prng

import (
	"crypto/sha256"
	"sync"
)

type PseudoRandomNumberGenerator interface {
	Generate(int64) uint8
}

type generator struct {
	mu             *sync.Mutex
	computedValues map[int64]uint8
	seed           []byte
}

func NewGenerator(seed string) (PseudoRandomNumberGenerator, error) {
	h := sha256.New()
	h.Write([]byte(seed))
	seedBytes := h.Sum(nil)

	computedValues := make(map[int64]uint8)

	mu := &sync.Mutex{}

	return generator{seed: seedBytes, computedValues: computedValues, mu: mu}, nil
}

func (g generator) Generate(pos int64) uint8 {
	if pos < 0 {
		pos = -pos
	}

	g.mu.Lock()
	preComputedValue, exists := g.computedValues[pos]
	g.mu.Unlock()
	if exists {
		return preComputedValue
	}

	if pos == 0 {
		res := g.seed[0]

		for i := 1; i < len(g.seed); i++ {
			res += g.seed[i]
		}

		g.mu.Lock()
		g.computedValues[pos] = res
		g.mu.Unlock()

		return res
	}

	if pos == 1 {
		zeroValue := byte(g.Generate(0))
		res := g.seed[0]

		for i := 1; i < len(g.seed); i++ {
			res += (zeroValue >> i) % g.seed[i]
		}

		g.mu.Lock()
		g.computedValues[pos] = res
		g.mu.Unlock()

		return res
	}

	// Xn+1 = (aXn + c) mod m
	// m, 0 < m  - modulus
	// a, 0 < a < m  - multiplier
	// c, 0 ? c < m  - increment
	// x0, 0 ? x0 < m  - the seed or start value

	return 0
}
