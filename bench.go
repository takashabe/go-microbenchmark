package microbenchmark

import "time"

// Benchmark object
type Benchmark struct {
	// get begin/end time in used benchmark
	beginFunc func() time.Time
	endFunc   func() time.Time

	// benchmark start time
	start time.Time
}

var defaultFunc = func() time.Time {
	return time.Now()
}

// Create default benhcmark object
func NewBenchmark() *Benchmark {
	return &Benchmark{
		beginFunc: defaultFunc,
		endFunc:   defaultFunc,
	}
}

// measure args functions. return all functions sum elapsed time
func (b *Benchmark) Benchmark(args ...func()) (result time.Duration) {
	b.start = b.beginFunc()
	defer func() {
		result = b.endFunc().Sub(b.start)
	}()
	for _, fn := range args {
		fn()
	}
	return result
}

// begin benchmark
func (b *Benchmark) Begin() {
	b.start = b.beginFunc()
}

// end benchmark, return since Begin() to End() elapsed time
func (b *Benchmark) End() time.Duration {
	return b.endFunc().Sub(b.start)
}
