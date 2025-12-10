package humans_test

import (
	"testing"
	"time"

	"github.com/GeekchanskiY/pet-project/pkg/clock"
	"github.com/GeekchanskiY/pet-project/pkg/humans"
	"github.com/stretchr/testify/assert"
)

func TestStat(t *testing.T) {
	c := clock.New(time.Date(2020, 1, 1, 0, 0, 0, 0, time.UTC))

	t.Run("default usage", func(t *testing.T) {
		stat := humans.NewStat(c.Now(), "hunger", 100)
		assert.Equal(t, "hunger", stat.Name())
		assert.Equal(t, 100, stat.Value())

		assert.Equal(t, c.Now(), stat.GetLastChange())
		c.Tick()
		assert.NotEqual(t, c.Now(), stat.GetLastChange())

		stat.Change(c.Now(), 200)
		assert.Equal(t, c.Now(), stat.GetLastChange())
		assert.Equal(t, 200, stat.Value())
	})
}
