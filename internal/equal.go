package internal

import (
	"testing"
)

func Equal[T comparable](t *testing.T, fatal bool, expected, actual T, labels ...string) {
	if actual == expected {
		logPass(t, "expected '%v'", labels, expected)
		return
	}
	logFail(t, "expected '%v', actual '%v'", labels, expected, actual)
	if fatal {
		t.FailNow()
	}
}

func NotEqual[T comparable](t *testing.T, fatal bool, notExpected, actual T, labels ...string) {
	if actual != notExpected {
		logPass(t, "actual '%v', expected not '%v'", labels, actual, notExpected)
		return
	}

	logFail(t, "expected not '%v'", labels, notExpected)
	if fatal {
		t.FailNow()
	}
}

func EqualSlice[T comparable](t *testing.T, fatal bool, expected, actual []T, labels ...string) {
	if areEqualSlices(expected, actual) {
		logPass(t, "expected '%v'", labels, expected)
		return
	}
	logFail(t, "expected '%v', actual '%v'", labels, expected, actual)
	if fatal {
		t.FailNow()
	}
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
