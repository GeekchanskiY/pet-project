package prng_test

import (
	"testing"

	"github.com/GeekchanskiY/pet-project/pkg/prng"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
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
		assert.Equal(t, uint8(102), generated)

		generated = generator.Generate(2)
		assert.Equal(t, uint8(170), generated)

		generated = generator.Generate(3)
		assert.Equal(t, uint8(141), generated)

		generated = generator.Generate(2)
		assert.Equal(t, uint8(170), generated, "already computed value must remain same")

		// duplicate generator with same seed
		generator, err = prng.NewGenerator("sample seed")
		if err != nil {
			t.Fatal(err)
		}

		generated = generator.Generate(0)
		assert.Equal(t, uint8(66), generated,
			"generated value in different generator with same seed must remain same")

		generated = generator.Generate(1)
		assert.Equal(t, uint8(102), generated,
			"generated value in different generator with same seed must remain same")

		generated = generator.Generate(2)
		assert.Equal(t, uint8(170), generated,
			"generated value in different generator with same seed must remain same")

		generated = generator.Generate(3)
		assert.Equal(t, uint8(141), generated,
			"generated value in different generator with same seed must remain same")

		generated = generator.Generate(2)
		assert.Equal(t, uint8(170), generated, "already computed value must remain same")
	})

	t.Run("values 'randomness'", func(t *testing.T) {
		generator, err := prng.NewGenerator("sample seed")
		if err != nil {
			t.Fatal(err)
		}

		values := make(map[uint8]uint64, 1000)
		for i := 0; i <= 255; i++ {
			values[uint8(i)] = 0
		}

		for i := int64(0); i < 2000; i++ {
			res := generator.Generate(i)
			values[res] = values[res] + 1
		}

		zeroValues := make([]uint8, 0, 255)
		for num, amount := range values {
			if amount == 0 {
				zeroValues = append(zeroValues, num)
			}
		}

		require.Less(t, len(zeroValues), 100)
	})
}

func Benchmark_Generate(b *testing.B) {
	generator, err := prng.NewGenerator("sample seed xcsfafdsgasd")
	if err != nil {
		b.Fatal(err)
	}

	//values := make(map[uint8]uint64, b.N)
	//for i := 0; i <= 255; i++ {
	//	values[uint8(i)] = 0
	//}

	for i := int64(0); i < int64(b.N); i++ {
		generator.Generate(i)
		// values[value] = values[value] + 1
	}

	//fmt.Println(b.N)
	//fmt.Println("First: ", generator.Generate(0))
	//fmt.Println("Second: ", generator.Generate(1))
	//fmt.Println(values)
}
