package synchronization

import (
	"sync"
	"testing"
)

func TestCounter(t *testing.T) {
	t.Run("running in non-concurrent environment", func(t *testing.T) {
		expected := 3

		counter := NewCounter()

		counter.Inc()
		counter.Inc()
		counter.Inc()

		assertCount(t, counter, expected)
	})

	t.Run("running in concurrent environment", func(t *testing.T) {
		expectedCount := 1000
		counter := NewCounter()

		var wg sync.WaitGroup
		wg.Add(expectedCount)

		for i := 0; i < expectedCount; i++ {
			go func() {
				counter.Inc()
				wg.Done()
			}()
		}
		wg.Wait()

		assertCount(t, counter, expectedCount)
	})

}

func assertCount(t testing.TB, c *Counter, expected int) {
	t.Helper()

	if expected != c.Value() {
		t.Errorf("expected %d, but got %d", expected, c.Value())
	}
}
