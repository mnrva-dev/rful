package rful

import (
	"fmt"
	"runtime"
	"strings"
)

// "main.go:32"
func LocShortFile() string {
	frame := getFrame()

	path := strings.Split(frame.File, "/")
	file := path[len(path)-1]

	return fmt.Sprintf("%s:%d ", file, frame.Line)
}

// "/path/to/src/main.go:32"
func LocLongFile() string {
	frame := getFrame()

	file := frame.File

	return fmt.Sprintf("%s:%d ", file, frame.Line)
}

// "MyFunction:32"
func LocShortFunc() string {
	frame := getFrame()

	fn := shortFunc(frame)

	return fmt.Sprintf("%s:%d ", fn, frame.Line)
}

// "my/module/MyFunction.go:32"
func LocLongFunc() string {
	frame := getFrame()

	fn := frame.Function

	return fmt.Sprintf("%s:%d ", fn, frame.Line)
}

// main.go at MyFunction():12
func LocShortFileShortFunc() string {
	frame := getFrame()

	file := shortFile(frame)
	fn := shortFunc(frame)

	return fmt.Sprintf("%s at %s:%d ", file, fn, frame.Line)
}

// main.go at my/module.MyFuncion:12
func LocShortFileLongFunc() string {
	frame := getFrame()

	file := shortFile(frame)
	fn := frame.Function

	return fmt.Sprintf("%s at %s:%d ", file, fn, frame.Line)
}

// /path/to/src/main.go at my/module.MyFuncion:12
func LocLongFileLongFunc() string {
	frame := getFrame()

	file := frame.File
	fn := frame.Function

	return fmt.Sprintf("%s at %s:%d ", file, fn, frame.Line)
}

// /path/to/src/main.go at MyFuncion:12
func LocLongFileShortFunc() string {
	frame := getFrame()

	file := frame.File
	fn := shortFunc(frame)

	return fmt.Sprintf("%s at %s:%d ", file, fn, frame.Line)
}

func getFrame() runtime.Frame {
	pcs := make([]uintptr, 15)
	// need to skip 4 to skip getFrame(), rful.Loc, rful.w.Write(), and rful.Log()
	n := runtime.Callers(4, pcs)
	frames := runtime.CallersFrames(pcs[:n])
	frame, _ := frames.Next()
	return frame
}

func shortFunc(frame runtime.Frame) string {
	fullfn := strings.Split(frame.Function, ".")
	fn := fullfn[len(fullfn)-1]
	return fn
}

func shortFile(frame runtime.Frame) string {
	path := strings.Split(frame.File, "/")
	file := path[len(path)-1]
	return file
}
