package internal

import "testing"

func Nil(t *testing.T, fatal bool, ptr interface{}, labels ...string) {
	if ptr == nil {
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
	if ptr != nil {
		logPass(t, "expected not nil", labels)
		return
	}
	logFail(t, "expected not nil", labels)
	if fatal {
		t.FailNow()
	}
}
