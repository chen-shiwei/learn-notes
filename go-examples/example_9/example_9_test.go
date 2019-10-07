package main

import (
	"reflect"
	"testing"
	"time"
)

func TestCountDown(t *testing.T) {
	// t.Run("Print 5 to go !", func(t *testing.T) {
	// 	buffer := new(bytes.Buffer)
	// 	cdo := new(CountdownOperactionSpy)
	// 	CountDown(buffer, cdo)
	// 	got := buffer.String()
	// 	want := "3\n2\n1\nGo!"
	// 	if got != want {
	// 		t.Errorf("got %s want %s", got, want)
	// 	}
	// })

	t.Run("sleep after every print", func(t *testing.T) {
		sysSleepPrinter := new(CountdownOperactionSpy)
		CountDown(sysSleepPrinter, sysSleepPrinter)
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
		t.Log(sysSleepPrinter)
		if !reflect.DeepEqual(want, sysSleepPrinter.Calls) {
			t.Errorf("want calls %v got %v ", want, sysSleepPrinter)
		}
	})
}

func TestConfigurableSleeper(t *testing.T) {
	sleepTime := time.Second * 5
	spyTime := new(SpyTime)
	sleeper := ConfigurableSleeper{
		sleepTime,
		spyTime.Sleep,
	}
	sleeper.Sleep()
	if spyTime.durationSlept != sleepTime {
		t.Errorf("should have slept for %v but slept for %v",
			sleepTime,
			spyTime.durationSlept)
	}
	// sleeper.
}

const write = "writre"
const sleep = "sleep"

type CountdownOperactionSpy struct {
	Calls []string
}

func (s *CountdownOperactionSpy) Sleep() {
	s.Calls = append(s.Calls, sleep)
}

func (s *CountdownOperactionSpy) Write(p []byte) (n int, err error) {
	s.Calls = append(s.Calls, write)
	return
}

type SpyTime struct {
	durationSlept time.Duration
}

func (s *SpyTime) Sleep(duration time.Duration) {
	s.durationSlept = duration
}
