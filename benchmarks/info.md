see info.md in /tests folder for `go test -v`
For benchmarks:
```go test -bench="."```

For memory bench:
```go test -bench="." -benchmem```
<!-- ------------------------------------------------------------ -->

One bench test:
```go test -bench=BenchmarkFields```

Bench tests by mask:
```go test -bench="Fields|Split"```

Only benchmarks without tests:
```go test -run="!" -bench="."```
