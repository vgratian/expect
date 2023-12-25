package expect

import (
	"github.com/vgratian/go-testutil/internal"
	"testing"
)

// Equal passes the test if expected and actual are equal, otherwise fails current test-function
func Equal[T comparable](t *testing.T, expected, actual T, labels ...string) {
	t.Helper()
	internal.Equal(t, false, expected, actual, labels...)
}

// NotEqual fails the test if notExpected and actual are equal, otherwise fails current test-function
func NotEqual[T comparable](t *testing.T, expected, actual T, labels ...string) {
	t.Helper()
	internal.NotEqual(t, false, expected, actual, labels...)
}

// EqualSlice passes the test if the two slices have exactly same elements, otherwise fails current test-function
func EqualSlice[T comparable](t *testing.T, expected, actual []T, labels ...string) {
	t.Helper()
	internal.EqualSlice(t, false, expected, actual, labels...)
}

// Greater passes the test if a is greater than b, otherwise fails current test-function
func Greater[T internal.Numeric](t *testing.T, a, b T, labels ...string) {
	t.Helper()
	internal.Greater(t, false, a, b, labels...)
}

// Less passes the test if a is less than b, otherwise fails current test-function
func Less[T internal.Numeric](t *testing.T, a, b T, labels ...string) {
	t.Helper()
	internal.Less(t, false, a, b, labels...)
}

// GreaterOrEqual passes the test if a is greater than or equal to b, otherwise fails current test-function
func GreaterOrEqual[T internal.Numeric](t *testing.T, a, b T, labels ...string) {
	t.Helper()
	internal.GreaterOrEqual(t, false, a, b, labels...)
}

// LessOrEqual passes the test if a is less than b, otherwise fails current test-function
func LessOrEqual[T internal.Numeric](t *testing.T, a, b T, labels ...string) {
	t.Helper()
	internal.LessOrEqual(t, false, a, b, labels...)
}

// Error passes the test if err is not nil, otherwise fails current test-function
func Error(t *testing.T, err error, labels ...string) {
	t.Helper()
	internal.Error(t, false, err, labels...)
}

// NoError passes the test if err is nil, otherwise fails current test-function
func NoError(t *testing.T, err error, labels ...string) {
	t.Helper()
	internal.NoError(t, false, err, labels...)
}

// Nil passes the test if ptr is nil, otherwise fails current test-function
func Nil(t *testing.T, ptr interface{}, labels ...string) {
	t.Helper()
	internal.Nil(t, false, ptr, labels...)
}

// NotNil passes the test if ptr is nil, otherwise fails current test-function
func NotNil(t *testing.T, ptr interface{}, labels ...string) {
	t.Helper()
	internal.NotNil(t, false, ptr, labels...)
}
