package main

import (
	"sync"
	"testing"
)

func TestCounter(t *testing.T) {
	t.Run("increment counter", func(t *testing.T) {
		counter := Counter{}
		counter.Inc()
		counter.Inc()
		counter.Inc()
		assertCounter(t, &counter, 3)
	})

	t.Run("it runs safely concurrently", func(t *testing.T) {
		wantedCount := 1000
		counter := Counter{}

		var wg sync.WaitGroup
		wg.Add(wantedCount)

		for i := 0; i < wantedCount; i++ {
			go func() {
				counter.Inc()
				wg.Done()
			}()
		}
		wg.Wait()

		assertCounter(t, &counter, wantedCount)
	})
}

func assertCounter(t testing.TB, counter *Counter, want int) {
	t.Helper()
	got := counter.Value()
	if got != want {
		t.Errorf("got %d, want %d", got, want)
	}
}
