package expect

import (
	"fmt"
	"testing"
)

// Equal passes the test if expected and actual are equal
func Equal[T comparable](t *testing.T, label string, expected, actual T) bool {
	if actual == expected {
		logPass(t, "expected %s: %v", label, expected)
		return true
	}
	logFail(t, "expected %s: %v, actual: %v", label, expected, actual)
	return false
}

// NotEqual fails the test if notExpected and actual are equal
func NotEqual[T comparable](t *testing.T, label string, notExpected, actual T) bool {
	if actual != notExpected {
		logPass(t, "actual: %v, expected not: %v", label, actual, notExpected)
		return true
	}
	logFail(t, "actual: %v", label, actual)
	return false
}

type numericType interface {
	int | uint | int8 | uint8 | int16 | uint16 | int32 | uint32 | int64 | uint64 | float32 | float64
}

// Greater passes the test if a is greater than b
func Greater[T numericType](t *testing.T, label string, a, b T) bool {
	if a > b {
		logPass(t, "%s: %v vs %v", label, a, b)
		return true
	}
	logFail(t, "%s: %v vs %v", label, a, b)
	return false
}

// Less passes the test if a is less than b
func Less[T numericType](t *testing.T, label string, a, b T) bool {
	if a < b {
		logPass(t, "%s: %v vs %v", label, a, b)
		return true
	}
	logFail(t, "%s: %v vs %v", label, a, b)
	return false
}

// EqualSlice passes the test if the two slices have exactly same elements
func EqualSlice[T comparable](t *testing.T, label string, expected, actual []T) bool {
	if areEqualSlices(expected, actual) {
		logPass(t, "expected %s: %v", label, expected)
		return true
	}
	logFail(t, "expected %s: %v, actual: %v", label, expected, actual)
	return false
}

// Error passes the test if err is not nil, panics otherwise
func Error(t *testing.T, label string, err error) bool {
	if err != nil {
		logPass(t, "%s error: %s", label, err.Error())
		return true
	}
	logFail(t, "%s: expected error, but got nil", label)
	return false
}

// NoError passes the test if err is nil, panics otherwise
func NoError(t *testing.T, label string, err error) bool {
	if err == nil {
		logPass(t, "%s: no error", label)
		return true
	}
	logFail(t, "%s: %s", label, err.Error())
	return false
}

// Nil passes the test if a pointer is nil
func Nil(t *testing.T, label string, ptr interface{}) bool {
	if ptr == nil {
		logPass(t, "%s: not nil", label)
		return true
	}
	logFail(t, "%s: nil", label)
	return false
}

// NotNil fails the test if a pointer is nil
func NotNil(t *testing.T, label string, ptr interface{}) bool {
	if ptr != nil {
		logPass(t, "%s: not nil", label)
		return true
	}
	logFail(t, "%s: nil", label)
	return false
}

func areEqualSlices[T comparable](a, b []T) bool {
	if len(a) != len(b) {
		return false
	}
	for i := range a {
		if a[i] != b[i] {
			return false
		}
	}
	return true
}

const (
	labelPass = "\033[32m" + "PASS" + "\033[0m"
	labelFail = "\033[31m" + "FAIL" + "\033[0m"
)

func logPass(t *testing.T, f string, a ...interface{}) {
	args := make([]interface{}, 1, len(a)+1)
	args[0] = labelPass
	t.Log(fmt.Sprintf("%s - "+f, append(args, a...)...))
}

func logFail(t *testing.T, f string, a ...interface{}) {
	args := make([]interface{}, 1, len(a)+1)
	args[0] = labelFail
	t.Error(fmt.Sprintf("%s - "+f, append(args, a...)...))
}
