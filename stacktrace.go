package assert

import (
	"fmt"
	"runtime"
)

type frame struct {
	Name string
	File string
	Line int
}

func stacktrace(skip int) []frame {
	frames := make([]frame, 0)
	for {
		pc, file, line, ok := runtime.Caller(skip)
		if !ok {
			break
		}
		frames = append(frames, frame{
			Name: runtime.FuncForPC(pc).Name(),
			File: file,
			Line: line,
		})
		skip++
		if skip > MaxTraceDepth {
			break
		}
	}
	return frames
}

func (f frame) String() string {
	return f.Name + " (" + f.File + ":" + fmt.Sprint(f.Line) + ")"
}
