//go:build assert
// +build assert

package assert

import (
	"cmp"
	"fmt"
	"os"
	"reflect"
)

func Equals[T comparable](a, b T) {
	if a != b {
		reportFailure("Equals", stacktrace(2), a, b)
	}
}

func NotEquals[T comparable](a, b T) {
	if a == b {
		reportFailure("NotEquals", stacktrace(2), a, b)
	}
}

func LessThan[T cmp.Ordered](a, b T) {
	if a >= b {
		reportFailure("LessThan", stacktrace(2), a, b)
	}
}

func MoreThan[T cmp.Ordered](a, b T) {
	if a <= b {
		reportFailure("MoreThan", stacktrace(2), a, b)
	}
}

func LessOrEquals[T cmp.Ordered](a, b T) {
	if a > b {
		reportFailure("LessOrEquals", stacktrace(2), a, b)
	}
}

func MoreOrEquals[T cmp.Ordered](a, b T) {
	if a < b {
		reportFailure("MoreOrEquals", stacktrace(2), a, b)
	}
}

func Nil(v any) {
	if !isNil(v) {
		reportFailure("Nil", stacktrace(2), v)
	}
}

func NotNil(v any) {
	if isNil(v) {
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
	if !v {
		reportFailure("Always", stacktrace(2), v)
	}
	return v
}

func Never(v bool) bool {
	if v {
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
				// Could we possibly print the argument name here? We would need
				// to parse the source code to get the argument name.
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
