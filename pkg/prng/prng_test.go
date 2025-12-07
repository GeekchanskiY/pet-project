package prng_test

import (
	"fmt"
	"testing"

	"github.com/GeekchanskiY/pet-project/pkg/prng"
	"github.com/stretchr/testify/assert"
)

func Test_Generate(t *testing.T) {
	t.Run("default usage", func(t *testing.T) {
		generator, err := prng.NewGenerator("sample seed")
		if err != nil {
			t.Fatal(err)
		}

		generated := generator.Generate(0)
		assert.Equal(t, uint8(66), generated)

		generated = generator.Generate(1)
		assert.Equal(t, uint8(71), generated)

		generated = generator.Generate(2)
		assert.Equal(t, uint8(71), generated)

		generated = generator.Generate(3)
		assert.Equal(t, uint8(82), generated)

		generated = generator.Generate(2)
		assert.Equal(t, uint8(71), generated)

		// duplicate generator with same seed
		generator, err = prng.NewGenerator("sample seed")
		if err != nil {
			t.Fatal(err)
		}

		generated = generator.Generate(0)
		assert.Equal(t, uint8(66), generated)

		generated = generator.Generate(1)
		assert.Equal(t, uint8(71), generated)

		generated = generator.Generate(2)
		assert.Equal(t, uint8(71), generated)

		generated = generator.Generate(3)
		assert.Equal(t, uint8(82), generated)

		generated = generator.Generate(2)
		assert.Equal(t, uint8(71), generated)
	})
}

func Benchmark_Generate(b *testing.B) {
	generator, err := prng.NewGenerator("sample seed xcsfafdsgasd")
	if err != nil {
		b.Fatal(err)
	}

	values := make(map[uint8]uint64, b.N)
	for i := 0; i <= 255; i++ {
		values[uint8(i)] = 0
	}

	for i := int64(0); i < int64(b.N); i++ {
		value := generator.Generate(i)
		values[value] = values[value] + 1
	}

	fmt.Println(b.N)
	fmt.Println("First: ", generator.Generate(0))
	fmt.Println("Second: ", generator.Generate(1))
	fmt.Println(values)
}
