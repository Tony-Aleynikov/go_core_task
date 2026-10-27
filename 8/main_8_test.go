package main

import (
	"math"
	"runtime"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestCustomWaitGroup_Add(t *testing.T) {
	t.Run("increments count", func(t *testing.T) {
		var wg CustomWaitGroup
		wg.Add(2)

		assert.Equal(t, 2, wg.count)
	})

	t.Run("allows decrement without reaching zero", func(t *testing.T) {
		var wg CustomWaitGroup
		wg.Add(2)
		wg.Add(-1)

		assert.Equal(t, 1, wg.count)
	})

	t.Run("zero delta does not change count", func(t *testing.T) {
		var wg CustomWaitGroup
		wg.Add(2)
		wg.Add(0)

		assert.Equal(t, 2, wg.count)
	})

	t.Run("rejects negative count", func(t *testing.T) {
		var wg CustomWaitGroup

		assert.Panics(t, func() { wg.Add(-1) })
	})

	t.Run("rejects counter overflow", func(t *testing.T) {
		var wg CustomWaitGroup
		wg.Add(math.MaxInt)

		assert.Panics(t, func() { wg.Add(1) })
		assert.Equal(t, math.MaxInt, wg.count)
	})
}

func TestCustomWaitGroup_Done(t *testing.T) {
	t.Run("decrements count", func(t *testing.T) {
		var wg CustomWaitGroup
		wg.Add(2)
		wg.Done()

		assert.Equal(t, 1, wg.count)
	})

	t.Run("rejects negative count", func(t *testing.T) {
		var wg CustomWaitGroup

		assert.Panics(t, wg.Done)
	})
}

func TestCustomWaitGroup_Wait(t *testing.T) {
	t.Run("returns for zero count", func(t *testing.T) {
		var wg CustomWaitGroup
		done := make(chan struct{})
		go func() {
			wg.Wait()
			close(done)
		}()

		select {
		case <-done:
		case <-time.After(time.Second):
			t.Fatal("Wait did not return for a zero-value wait group")
		}
	})

	t.Run("waits until count reaches zero", func(t *testing.T) {
		var wg CustomWaitGroup
		wg.Add(2)
		done := make(chan struct{})
		go func() {
			wg.Wait()
			close(done)
		}()
		waitForWaiterCount(t, &wg, 1)

		wg.Done()
		select {
		case <-done:
			t.Fatal("Wait returned before all tasks completed")
		case <-time.After(20 * time.Millisecond):
		}

		wg.Done()
		select {
		case <-done:
		case <-time.After(time.Second):
			t.Fatal("Wait did not return after all tasks completed")
		}
	})

	t.Run("wakes all waiters and supports reuse", func(t *testing.T) {
		var wg CustomWaitGroup
		for generation := range 2 {
			wg.Add(1)
			const waiterCount = 3
			done := make(chan struct{}, waiterCount)
			for range waiterCount {
				go func() {
					wg.Wait()
					done <- struct{}{}
				}()
			}
			waitForWaiterCount(t, &wg, waiterCount)

			wg.Done()
			for range waiterCount {
				select {
				case <-done:
				case <-time.After(time.Second):
					t.Fatalf("generation %d: not all waiters returned", generation)
				}
			}
		}
	})
}

// waitForWaiterCount помогает нам дождаться,
// пока количество ожидающих Wait() вызовов не станет равным count
func waitForWaiterCount(t *testing.T, wg *CustomWaitGroup, count int) {
	t.Helper()
	deadline := time.Now().Add(time.Second)
	for time.Now().Before(deadline) {
		wg.mu.Lock()
		waiterCount := len(wg.waiters)
		wg.mu.Unlock()
		if waiterCount == count {
			return
		}
		runtime.Gosched()
	}
	t.Fatalf("timed out waiting for %d registered waiters", count)
}
