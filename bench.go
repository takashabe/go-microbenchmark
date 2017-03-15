package microbenchmark

import "time"

type Benchmark struct{}

func NewBenchmark() *Benchmark {
	return &Benchmark{}
}

func (b *Benchmark) Benchmark(args ...func()) (result time.Duration) {
	start := time.Now()
	defer func() {
		result = time.Since(start)
	}()
	for _, fn := range args {
		fn()
	}
	return result
}
