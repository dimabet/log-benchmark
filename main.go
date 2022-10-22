package main

import (
	"runtime"

	"go.uber.org/zap"
)

func main() {}

// whateverLogger logs n times with provided logger
func whateverLogger(logger *zap.Logger, n int) {
	for i := 0; i < n; i++ {
		logger.Info("whateverLogger", zap.Int("i", i))
	}
}

// whateverFrameLogger logs n times with provided logger and adds caller function from runtime
func whateverFrameLogger(logger *zap.Logger, n int) {
	for i := 0; i < n; i++ {
		logger.Info("whateverLogger", zap.Int("i", i), zap.String("func", MyCaller()))
	}
}

// MyCaller returns the caller of the function that called it
func MyCaller() string {
	// Skip GetCallerFunctionName and the function to get the caller of
	return getFrame(2).Function
}

func getFrame(skipFrames int) runtime.Frame {
	// We need the frame at index skipFrames+2, since we never want runtime.Callers and getFrame
	targetFrameIndex := skipFrames + 2

	// Set size to targetFrameIndex+2 to ensure we have room for one more caller than we need
	programCounters := make([]uintptr, targetFrameIndex+2)
	n := runtime.Callers(0, programCounters)

	frame := runtime.Frame{Function: "unknown"}
	if n > 0 {
		frames := runtime.CallersFrames(programCounters[:n])
		for more, frameIndex := true, 0; more && frameIndex <= targetFrameIndex; frameIndex++ {
			var frameCandidate runtime.Frame
			frameCandidate, more = frames.Next()
			if frameIndex == targetFrameIndex {
				frame = frameCandidate
			}
		}
	}

	return frame
}
