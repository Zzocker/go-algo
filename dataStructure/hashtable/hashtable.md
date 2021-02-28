# Hash table

## Benchmark

```
With Chaining :
BenchmarkChaining-4      1000000              1671 ns/op             208 B/op          6 allocs/op

With Open Addressing :
BenchmarkOpenAddress-4           1000000              1616 ns/op             208 B/op          6 allocs/op

Go Map :
BenchmarkGoMap-4         2213106               494 ns/op             154 B/op          1 allocs/op
```