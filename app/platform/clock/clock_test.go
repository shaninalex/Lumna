package clock

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestSystemClock(t *testing.T) {
	clk := System()
	before := time.Now()
	now := clk.Now()
	after := time.Now()

	assert.False(t, now.Before(before))
	assert.False(t, now.After(after))
}

func TestMockClock(t *testing.T) {
	initial := time.Date(2026, time.September, 12, 12, 0, 0, 0, time.UTC)
	mock := NewMock(initial)

	assert.Equal(t, initial, mock.Now())

	mock.Advance(15 * time.Minute)
	assert.Equal(t, initial.Add(15*time.Minute), mock.Now())

	newTime := time.Date(2027, time.January, 1, 0, 0, 0, 0, time.UTC)
	mock.Set(newTime)
	assert.Equal(t, newTime, mock.Now())
}
