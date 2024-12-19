//go:build !assert
// +build !assert

package assert

import (
	"cmp"
)

func Equals[T comparable](a, b T)        {}
func NotEquals[T comparable](a, b T)     {}
func LessThan[T cmp.Ordered](a, b T)     {}
func MoreThan[T cmp.Ordered](a, b T)     {}
func LessOrEquals[T cmp.Ordered](a, b T) {}
func MoreOrEquals[T cmp.Ordered](a, b T) {}
func Nil(v any)                          {}
func NotNil(v any)                       {}
func Always(v bool) bool                 { return v }
func Never(v bool) bool                  { return v }
