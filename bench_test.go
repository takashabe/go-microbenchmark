package microbenchmark

import (
	"testing"
	"time"
)

func TestElapsedTime(t *testing.T) {
	sleep := func() {
		time.Sleep(time.Millisecond)
	}
	cases := []struct {
		input []func()
		// got between expect and expectLimit
		// WARNING: depends time.Now()
		expect      time.Duration
		expectLimit time.Duration
	}{
		{
			[]func(){sleep},
			time.Millisecond,
			time.Millisecond * 2,
		},
		{
			[]func(){sleep, sleep, sleep, sleep},
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
