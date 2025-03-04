package errx

import (
	"fmt"
	"reflect"
	"runtime"
	"strings"
)

type shadow struct{}

var (
	maxDepth    = 32
	packageName = reflect.TypeOf(shadow{}).PkgPath()
)

type Frame struct {
	FileName     string
	FileLine     int
	FunctionName string
}

func (f Frame) String() string {
	return fmt.Sprintf("%s:%d %s", f.FileName, f.FileLine, f.FunctionName)
}

func stack() []Frame {
	frames := make([]Frame, 0, maxDepth)
	for i := 0; i < maxDepth; i++ {
		pc, file, line, ok := runtime.Caller(i)
		if !ok {
			break
		}
		if strings.Contains(file, packageName) {
			continue
		}
		name := runtime.FuncForPC(pc).Name()
		frames = append(frames, Frame{file, line, name})
	}
	return frames
}
