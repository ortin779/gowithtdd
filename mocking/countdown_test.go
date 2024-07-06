package main

import (
	"bytes"
	"slices"
	"testing"
	"time"
)

const (
	sleep = "sleep"
	write = "write"
)

type SpyCountDownOperations struct {
	Calls []string
}

func (s *SpyCountDownOperations) Sleep() {
	s.Calls = append(s.Calls, sleep)
}

func (s *SpyCountDownOperations) Write(p []byte) (n int, err error) {
	s.Calls = append(s.Calls, write)
	return
}

func TestCountdown(t *testing.T) {
	t.Run("prints 3 to Go!", func(t *testing.T) {
		buffer := &bytes.Buffer{}
		sleeper := &SpyCountDownOperations{}

		Countdown(buffer, sleeper)

		got := buffer.String()
		expected := `3
2
1
Go!`

		if got != expected {
			t.Errorf("expected %q, but got %q", expected, got)
		}

		if len(sleeper.Calls) != 3 {
			t.Errorf("not enough call to sleeper, expected 3 but got %d", len(sleeper.Calls))
		}
	})

	t.Run("sleep before every write call", func(t *testing.T) {
		spySleeper := &SpyCountDownOperations{}

		Countdown(spySleeper, spySleeper)

		expectedCalls := []string{
			write,
			sleep,
			write,
			sleep,
			write,
			sleep,
			write,
		}

		if !slices.Equal(spySleeper.Calls, expectedCalls) {
			t.Errorf("wanted calls %v, but got %v", expectedCalls, spySleeper.Calls)
		}
	})
}

type SpyTime struct {
	durationSlept time.Duration
}

func (s *SpyTime) Sleep(duration time.Duration) {
	s.durationSlept = duration
}

func TestConfigurableSleeper(t *testing.T) {
	sleepTime := 5 * time.Second

	spyTime := &SpyTime{}

	sleeper := ConfigurableSleeper{sleepTime, spyTime.Sleep}

	sleeper.Sleep()

	if sleepTime != spyTime.durationSlept {
		t.Errorf("should have slept for %v but slept for %v", sleepTime, spyTime.durationSlept)
	}

}
