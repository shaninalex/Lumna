package clock

import (
	"sync"
	"time"
)

// Clock provides an interface for getting the current time.
// It allows injecting a mocked or frozen clock during testing.
type Clock interface {
	Now() time.Time
}

type systemClock struct{}

func (systemClock) Now() time.Time {
	return time.Now()
}

// System returns a Clock that uses real system time.
func System() Clock {
	return systemClock{}
}

// Mock provides a thread-safe controllable clock for unit tests.
type Mock struct {
	mu  sync.RWMutex
	now time.Time
}

// NewMock creates a Mock clock initialized at the specified time.
func NewMock(t time.Time) *Mock {
	return &Mock{now: t}
}

// Frozen creates a Mock clock initialized at current system time.
func Frozen(t time.Time) *Mock {
	return NewMock(t)
}

// Now returns the current simulated time.
func (m *Mock) Now() time.Time {
	m.mu.RLock()
	defer m.mu.RUnlock()
	return m.now
}

// Set explicitly sets the simulated time.
func (m *Mock) Set(t time.Time) {
	m.mu.Lock()
	defer m.mu.Unlock()
	m.now = t
}

// Advance moves simulated time forward by the given duration.
func (m *Mock) Advance(d time.Duration) {
	m.mu.Lock()
	defer m.mu.Unlock()
	m.now = m.now.Add(d)
}
