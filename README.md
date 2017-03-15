# go-microbenchmark

micro benchmark library

## feature

1. measure running time some functions

```
bench := NewBenchmark()
elapsed := bench.Benchmark(
  func() { ... },
  func() { ... },
  func() { ... },
)
```
