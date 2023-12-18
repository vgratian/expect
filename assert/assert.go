package assert

import (
	"github.com/vgratian/expect/internal"
	"testing"
)

// Equal passes the test if expected and actual are equal, otherwise aborts current test-function
func Equal[T comparable](t *testing.T, expected, actual T, labels ...string) {
	internal.Equal(t, true, expected, actual, labels...)
}

// NotEqual fails the test if notExpected and actual are equal, otherwise aborts current test-function
func NotEqual[T comparable](t *testing.T, expected, actual T, labels ...string) {
	internal.NotEqual(t, true, expected, actual, labels...)
}

// EqualSlice passes the test if the two slices have exactly same elements, otherwise aborts current test-function
func EqualSlice[T comparable](t *testing.T, expected, actual []T, labels ...string) {
	internal.EqualSlice(t, true, expected, actual, labels...)
}

// Greater passes the test if a is greater than b, otherwise aborts current test-function
func Greater[T internal.Numeric](t *testing.T, a, b T, labels ...string) {
	internal.Greater(t, true, a, b, labels...)
}

// Less passes the test if a is less than b, otherwise aborts current test-function
func Less[T internal.Numeric](t *testing.T, a, b T, labels ...string) {
	internal.Less(t, true, a, b, labels...)
}

// GreaterOrEqual passes the test if a is greater than or equal to b, otherwise aborts current test-function
func GreaterOrEqual[T internal.Numeric](t *testing.T, a, b T, labels ...string) {
	internal.GreaterOrEqual(t, true, a, b, labels...)
}

// LessOrEqual passes the test if a is less than b, otherwise aborts current test-function
func LessOrEqual[T internal.Numeric](t *testing.T, a, b T, labels ...string) {
	internal.LessOrEqual(t, true, a, b, labels...)
}

// Error passes the test if err is not nil, otherwise aborts current test-function
func Error(t *testing.T, err error, labels ...string) {
	internal.Error(t, true, err, labels...)
}

// NoError passes the test if err is nil, otherwise aborts current test-function
func NoError(t *testing.T, err error, labels ...string) {
	internal.NoError(t, true, err, labels...)
}

// Nil passes the test if ptr is nil, otherwise aborts current test-function
func Nil(t *testing.T, ptr interface{}, labels ...string) {
	internal.Nil(t, true, ptr, labels...)
}

// NotNil passes the test if ptr is nil, otherwise aborts current test-function
func NotNil(t *testing.T, ptr interface{}, labels ...string) {
	internal.NotNil(t, true, ptr, labels...)
}
