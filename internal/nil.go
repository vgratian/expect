package internal

import (
	"reflect"
	"testing"
)

func Nil(t *testing.T, fatal bool, ptr interface{}, labels ...string) {
	t.Helper()
	if isNil(ptr) {
		logPass(t, "expected nil", labels)
		return
	}
	logFail(t, "expected nil", labels)
	if fatal {
		t.FailNow()
	}
}

// NotNil fails the test if a pointer is nil
func NotNil(t *testing.T, fatal bool, ptr interface{}, labels ...string) {
	t.Helper()
	if !isNil(ptr) {
		logPass(t, "expected not nil", labels)
		return
	}
	logFail(t, "expected not nil", labels)
	if fatal {
		t.FailNow()
	}
}

func isNil(ptr interface{}) bool {
	return reflect.ValueOf(ptr).IsZero()
}
