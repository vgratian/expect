package internal

import (
	"testing"
)

type T struct{}

func Test_isNil(t *testing.T) {
	var ptr *T
	// expect nil
	Equal(t, false, true, isNil(ptr))

	ptr = &T{}
	// expect not nil
	Equal(t, false, false, isNil(ptr))
}
