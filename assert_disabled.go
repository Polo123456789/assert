//go:build !assert
// +build !assert

package assert

import (
	"cmp"
)

func Equals[T comparable](a, b T, context ...any)        {}
func NotEquals[T comparable](a, b T, context ...any)     {}
func LessThan[T cmp.Ordered](a, b T, context ...any)     {}
func MoreThan[T cmp.Ordered](a, b T, context ...any)     {}
func LessOrEquals[T cmp.Ordered](a, b T, context ...any) {}
func MoreOrEquals[T cmp.Ordered](a, b T, context ...any) {}
func Nil(v any, context ...any)                          {}
func NotNil(v any, context ...any)                       {}

func Always(v bool) bool {
	return v
}

func Never(v bool) bool {
	return v
}
