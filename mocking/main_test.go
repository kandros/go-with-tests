package main

import (
	"bytes"
	"reflect"
	"testing"
	"time"
)

type CountdownOperationsSpy struct {
	Calls []string
}

func (s *CountdownOperationsSpy) Sleep() {
	s.Calls = append(s.Calls, sleep)
}

func (s *CountdownOperationsSpy) Write(p []byte) (n int, err error) {
	s.Calls = append(s.Calls, write)
	return
}

const (
	write = "write"
	sleep = "sleep"
)

func TestCountDown(t *testing.T) {
	t.Run("sleep before every print", func(t *testing.T) {
		spySleepPrinter := &CountdownOperationsSpy{}
		Countdown(spySleepPrinter, spySleepPrinter)

		got := spySleepPrinter.Calls
		want := []string{
			sleep,
			write,
			sleep,
			write,
			sleep,
			write,
			sleep,
			write,
		}

		if !reflect.DeepEqual(want, got) {
			t.Errorf("got '%v' want '%v'", got, want)
		}
	})

	t.Run("prints 3 to Go!", func(t *testing.T) {
		buffer := bytes.Buffer{}
		spySleepPrinter := &CountdownOperationsSpy{}

		Countdown(&buffer, spySleepPrinter)

		got := buffer.String()
		want := `3
2
1
Go!`

		if got != want {
			t.Errorf("got '%v' want '%v'", got, want)
		}
	})
}

type SpyTime struct {
	durationSlept time.Duration
}

func (s *SpyTime) Sleep(duration time.Duration) {
	s.durationSlept = duration
}

func TestConfiguarableSleeper(t *testing.T) {
	sleepTime := 5 * time.Second
	spyTime := SpyTime{}
	sleeper := ConfigurableSleeper{duration: sleepTime, sleep: spyTime.Sleep}
	sleeper.Sleep()

	if sleepTime != spyTime.durationSlept {
		t.Errorf("should have slept %v but slept for %v", sleepTime, spyTime.durationSlept)
	}
}
