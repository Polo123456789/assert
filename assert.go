package assert

import (
	"cmp"
	"fmt"
	"io"
	"os"
	"reflect"
)

var Enabled = false
var Writer io.Writer = os.Stderr

// ContextWindow represents the number of source code lines to display above and
// below the line that caused the assertion failure.
var ContextWindow int = 5

// MaxTraceDepth represents the maximum number of stack frames to display in the
// stack trace.
var MaxTraceDepth int = 20

// ReturnValue represents the exit code to use when an assertion fails.
var ReturnValue int = 1 // Not using 42069 requires a lot of self control

func Equals[T comparable](a, b T) {
	if Enabled && a != b {
		reportFailure("Equals", stacktrace(2), a, b)
	}
}

func NotEquals[T comparable](a, b T) {
	if Enabled && a == b {
		reportFailure("NotEquals", stacktrace(2), a, b)
	}
}

func LessThan[T cmp.Ordered](a, b T) {
	if Enabled && a >= b {
		reportFailure("LessThan", stacktrace(2), a, b)
	}
}

func MoreThan[T cmp.Ordered](a, b T) {
	if Enabled && a <= b {
		reportFailure("MoreThan", stacktrace(2), a, b)
	}
}

func LessOrEquals[T cmp.Ordered](a, b T) {
	if Enabled && a > b {
		reportFailure("LessOrEquals", stacktrace(2), a, b)
	}
}

func MoreOrEquals[T cmp.Ordered](a, b T) {
	if Enabled && a < b {
		reportFailure("MoreOrEquals", stacktrace(2), a, b)
	}
}

func Nil[T any](v any) {
	if Enabled && !isNil(v) {
		reportFailure("Nil", stacktrace(2), v)
	}
}

func NotNil[T any](v T) {
	if Enabled && isNil(v) {
		reportFailure("NotNil", stacktrace(2), v)
	}
}

func isNil(v any) bool {
	if v == nil {
		return true
	}

	switch reflect.TypeOf(v).Kind() {
	case reflect.Pointer, reflect.Interface, reflect.Map,
		reflect.Slice, reflect.Chan, reflect.Func:
		return reflect.ValueOf(v).IsNil()
	}

	return false
}

func Always(v bool) bool {
	if Enabled && !v {
		reportFailure("Always", stacktrace(2), v)
	}
	return v
}

func Never(v bool) bool {
	if Enabled && v {
		reportFailure("Never", stacktrace(2), v)
	}
	return v
}

func reportFailure(name string, stack []frame, args ...any) {
	fmt.Fprintln(
		Writer,
		applyAnsi(SectionTitleColor, name+" Assertion Failed:\n"),
	)

	fmt.Fprintln(Writer, get_context(stack[0]))

	if len(args) != 0 {
		fmt.Fprintln(Writer, applyAnsi(SectionTitleColor, "Values:\n"))
		for i, arg := range args {
			fmt.Fprintf(
				Writer,
				"\t%s: %v\n",
				applyAnsi(ArgumentColor, fmt.Sprintf("Argument %d", i+1)),
				arg,
			)
		}
		fmt.Fprintln(Writer)
	}

	fmt.Fprintln(Writer, applyAnsi(SectionTitleColor, "Stacktrace:\n"))
	for _, frame := range stack {
		fmt.Fprintf(
			Writer,
			"\t%s\n",
			applyAnsi(StackColor, frame.String()),
		)
	}
	fmt.Fprintln(Writer)
	os.Exit(ReturnValue)
}
