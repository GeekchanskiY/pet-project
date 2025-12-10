package clock_test

import (
	"testing"
	"time"

	"github.com/GeekchanskiY/pet-project/pkg/clock"
	"github.com/stretchr/testify/assert"
)

func TestClock(t *testing.T) {
	t.Run("default usage", func(t *testing.T) {
		c := clock.New(time.Date(2000, 1, 1, 0, 0, 0, 0, time.UTC))
		c.Tick()
		assert.Equal(t, time.Date(2000, 1, 1, 1, 0, 0, 0, time.UTC), c.Now())
		assert.Equal(t, time.Date(2000, 1, 1, 0, 0, 0, 0, time.UTC), c.StartTime())
	})
}
