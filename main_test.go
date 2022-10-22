package main

import (
	"testing"

	"go.uber.org/zap"
)

// BenchmarkHardcodeLogger runs 100 whateverLogger() with default zap logger and function name is hardcoded
// BenchmarkHardcodeLogger100-8   	  113385	     12194 ns/op
func BenchmarkHardcodeLogger100(b *testing.B) {
	config := zap.NewProductionConfig()
	logger, _ := config.Build()
	for n := 0; n < b.N; n++ {
		whateverLogger(logger, 100)
	}
}

// BenchmarkHardcodeLogger runs 1000 whateverLogger() with default zap logger and function name is hardcoded
// BenchmarkHardcodeLogger1000-8   	   10570	    113106 ns/op
func BenchmarkHardcodeLogger1000(b *testing.B) {
	config := zap.NewProductionConfig()
	logger, _ := config.Build()
	for n := 0; n < b.N; n++ {
		whateverLogger(logger, 1000)
	}
}

// BenchmarkZapFuncConfig runs 100 whateverLogger() with zap logger configured to add "caller" name
// BenchmarkZapFuncConfig100-8   	  113414	     12713 ns/op
func BenchmarkZapFuncConfig100(b *testing.B) {
	config := zap.NewProductionConfig()
	config.EncoderConfig.FunctionKey = "func"
	logger, _ := config.Build()
	for n := 0; n < b.N; n++ {
		whateverLogger(logger, 100)
	}
}

// BenchmarkZapFuncConfig runs 1000 whateverLogger() with zap logger configured to add "caller" name
// BenchmarkZapFuncConfig1000-8   	    9571	    123616 ns/op
func BenchmarkZapFuncConfig1000(b *testing.B) {
	config := zap.NewProductionConfig()
	config.EncoderConfig.FunctionKey = "func"
	logger, _ := config.Build()
	for n := 0; n < b.N; n++ {
		whateverLogger(logger, 1000)
	}
}

// BenchmarkFuncFrameLogger runs 100 whateverFrameLogger() with default zap logger and figures out caller name dynamically
// BenchmarkFuncFrameLogger100-8   	    9692	    104599 ns/op
func BenchmarkFuncFrameLogger100(b *testing.B) {
	config := zap.NewProductionConfig()
	logger, _ := config.Build()
	for n := 0; n < b.N; n++ {
		whateverFrameLogger(logger, 100)
	}
}

// BenchmarkFuncFrameLogger runs 1000 whateverFrameLogger() with default zap logger and figures out caller name dynamically
// BenchmarkFuncFrameLogger1000-8   	    1081	   1049376 ns/op
func BenchmarkFuncFrameLogger1000(b *testing.B) {
	config := zap.NewProductionConfig()
	logger, _ := config.Build()
	for n := 0; n < b.N; n++ {
		whateverFrameLogger(logger, 1000)
	}
}
