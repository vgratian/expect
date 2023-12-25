package internal

import "testing"

func Error(t *testing.T, fatal bool, err error, labels ...string) {
	t.Helper()
	if err != nil {
		logPass(t, "expected error - '%s'", labels, err.Error())
		return
	}
	logFail(t, "expected error - got nil", labels)
	if fatal {
		t.FailNow()
	}
}

func NoError(t *testing.T, fatal bool, err error, labels ...string) {
	t.Helper()
	if err == nil {
		logPass(t, "expected no error", labels)
		return
	}
	logFail(t, "expected no error - got '%s'", labels, err.Error())
	if fatal {
		t.FailNow()
	}
}
