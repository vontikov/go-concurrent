package concurrent

import (
	"runtime"
	"time"
)

// IdleStrategy is used when there is no work to do
type IdleStrategy interface {
	// Idle performs idle action
	Idle()
}

// sleepingIdleStrategy pauses execution for the period d.
type sleepingIdleStrategy struct {
	d time.Duration
}

// yeildingIdleStrategy Yields the processor.
type yeildingIdleStrategy struct{}

// NewSleepingIdleStrategy returns a new instance of sleepingIdleStrategy.
// Duration d defines the execution pause.
func NewSleepingIdleStrategy(d time.Duration) IdleStrategy {
	return &sleepingIdleStrategy{d: d}
}

// NewYeildingIdleStrategy returns a new instance of yeildingIdleStrategy
func NewYeildingIdleStrategy() IdleStrategy {
	return &yeildingIdleStrategy{}
}

// Idle performs idle action
func (s *sleepingIdleStrategy) Idle() {
	time.Sleep(s.d)
}

// Idle performs idle action
func (yeildingIdleStrategy) Idle() {
	runtime.Gosched()
}
