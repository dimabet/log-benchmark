# log-benchmark

### Benchmarks for 3 approaches:
- calling default Zap logger with `hardcoded` message
- calling Zap logger with `configured` field for showing caller function
- calling default Zap logger and on each call `dynamically` get the caller function

### Results
suffix in the benchmark name `100` or `1000` means that in a single function call `.Info()` was call 100 and 1000 times respectively

```
BenchmarkHardcodeLogger100-8   	  113385	     12194 ns/op --- hardcoded
BenchmarkZapFuncConfig100-8   	  113414	     12713 ns/op --- configured
BenchmarkFuncFrameLogger100-8     9692	             104599 ns/op --- dynamic
```

```
BenchmarkHardcodeLogger1000-8     10570	            113106 ns/op --- hardcoded
BenchmarkZapFuncConfig1000-8   	  9571	            123616 ns/op --- configured
BenchmarkFuncFrameLogger1000-8    1081	            1049376 ns/op --- dynamic
```