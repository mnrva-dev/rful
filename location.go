package rful

import (
	"fmt"
	"runtime"
)

// "main.go:32"
func LocShortFile() string {
	_, filename, line, _ := runtime.Caller(2)

	return fmt.Sprintf("%s:%d", filename, line)
}

// main.go at myFunction():12
func LocShortFileFunc() string {
	pcs := make([]uintptr, 15)
	n := runtime.Callers(2, pcs)
	frames := runtime.CallersFrames(pcs[:n])
	frame, _ := frames.Next()

	return fmt.Sprintf("%s at %s:%d", frame.File, frame.Function, frame.Line)
}
