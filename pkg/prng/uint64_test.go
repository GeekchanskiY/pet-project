package prng_test

import (
	"fmt"
	"testing"

	"github.com/GeekchanskiY/pet-project/pkg/prng"
)

// Used for manual check
func Benchmark_Uint64(b *testing.B) {
	generator := prng.NewUint64("sample seed")

	values := make(map[uint64]uint64, b.N)
	for i := 0; i <= 255; i++ {
		values[uint64(i)] = 0
	}

	for i := uint64(0); i < uint64(b.N); i++ {
		value := generator.Generate(i)
		_, ok := values[value]
		if !ok {
			values[value] = 0
		}

		values[value] = values[value] + 1
	}

	fmt.Println(b.N)
	fmt.Println("First: ", generator.Generate(0))
	fmt.Println("Second: ", generator.Generate(1))
	fmt.Println(values)
}
