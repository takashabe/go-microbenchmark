package microbenchmark

import (
	"testing"
	"time"
)

var testSleep = func() {
	time.Sleep(time.Millisecond)
}

func TestElapsedTimeWithBenchmark(t *testing.T) {
	cases := []struct {
		input []func()
		// got between expect and expectLimit
		// WARNING: depends time.Now()
		expect      time.Duration
		expectLimit time.Duration
	}{
		{
			[]func(){testSleep},
			time.Millisecond,
			time.Millisecond * 2,
		},
		{
			[]func(){testSleep, testSleep, testSleep, testSleep},
			time.Millisecond * 4,
			time.Millisecond * 4 * 2,
		},
	}
	for i, c := range cases {
		bench := NewBenchmark()
		got := bench.Benchmark(c.input...)
		if !(c.expect < got && got < c.expectLimit) {
			t.Errorf("#%d: want %v ~ %v, got %v", i, c.expect, c.expectLimit, got)
		}
	}
}

func TestElapsedTimeWithBeginEnd(t *testing.T) {
	cases := []struct {
		fnList []func()
		// got between expect and expectLimit
		// WARNING: depends time.Now()
		expect      time.Duration
		expectLimit time.Duration
	}{
		{
			[]func(){testSleep},
			time.Millisecond,
			time.Millisecond * 2,
		},
		{
			[]func(){testSleep, testSleep, testSleep, testSleep},
			time.Millisecond * 4,
			time.Millisecond * 4 * 2,
		},
	}
	for i, c := range cases {
		bench := NewBenchmark()
		bench.Begin()
		for _, fn := range c.fnList {
			fn()
		}
		got := bench.End()
		if !(c.expect < got && got < c.expectLimit) {
			t.Errorf("#%d: want %v ~ %v, got %v", i, c.expect, c.expectLimit, got)
		}
	}
}
