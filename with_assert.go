//go:build assert
// +build assert

package assert

import (
	"cmp"
	"fmt"
	"os"
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
	if v != nil {
		reportFailure("Nil", stacktrace(2), v)
	}
}

func NotNil(v any) {
	if v == nil {
		reportFailure("NotNil", stacktrace(2), v)
	}
}

func Always(v bool) bool {
	if !v {
		reportFailure("Always", stacktrace(2), v)
	}
	return true
}

func Never(v bool) bool {
	if v {
		reportFailure("Never", stacktrace(2), v)
	}
	return false
}

func reportFailure(name string, stack []frame, args ...any) {
	fmt.Fprintln(
		Writter,
		applyAnsi(SectionTitleColor, name+" Assertion Failed:\n"),
	)

	fmt.Fprintln(Writter, get_context(stack[0]))

	if len(args) != 0 {
		fmt.Fprintln(Writter, applyAnsi(SectionTitleColor, "Values:\n"))
		for i, arg := range args {
			fmt.Fprintf(
				Writter,
				"\t%s: %v\n",
				applyAnsi(ArgumentColor, fmt.Sprintf("Argument %d", i+1)),
				arg,
			)
		}
		fmt.Fprintln(Writter)
	}

	fmt.Fprintln(Writter, applyAnsi(SectionTitleColor, "Stacktrace:\n"))
	for _, frame := range stack {
		fmt.Fprintf(
			Writter,
			"\t%s\n",
			applyAnsi(StackColor, frame.String()),
		)
	}
	fmt.Fprintln(Writter)
	os.Exit(ReturnValue)
}
