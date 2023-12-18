package internal

import "testing"

type Numeric interface {
	int | uint | int8 | uint8 | int16 | uint16 | int32 | uint32 | int64 | uint64 | float32 | float64
}

func Greater[T Numeric](t *testing.T, fatal bool, a, b T, labels ...string) {
	if a > b {
		logPass(t, "expected greater - %v > %v", labels, a, b)
		return
	}
	logFail(t, "expected greater - %v ≯ %v", labels, a, b)
	if fatal {
		t.FailNow()
	}
}

func Less[T Numeric](t *testing.T, fatal bool, a, b T, labels ...string) {
	if a < b {
		logPass(t, "expected less - %v < %v", labels, a, b)
		return
	}
	logFail(t, "expected less - %v ≮ %v", labels, a, b)
	if fatal {
		t.FailNow()
	}
}

func GreaterOrEqual[T Numeric](t *testing.T, fatal bool, a, b T, labels ...string) {
	if a >= b {
		logPass(t, "expected greater or equal - %v ≥ %v", labels, a, b)
		return
	}
	logFail(t, "expected greater or equal - %v ≱ %v", labels, a, b)
	if fatal {
		t.FailNow()
	}
}

func LessOrEqual[T Numeric](t *testing.T, fatal bool, a, b T, labels ...string) {
	if a <= b {
		logPass(t, "expected less or equal - %v ≤ %v", labels, a, b)
		return
	}
	logFail(t, "expected less or equal - %v ≰ %v", labels, a, b)
	if fatal {
		t.FailNow()
	}
}
