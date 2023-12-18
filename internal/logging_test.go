package internal

import (
	"fmt"
	"testing"
)

/* Some of these tests are meant to fail
   !!!!
*/

func Test_Equal(t *testing.T) {
	Equal(t, false, 'a', 'a')
	Equal(t, false, 'a', 'b')
	Equal(t, false, 42, 42, "calculated value")
	Equal(t, false, 42, 0, "calculated value")
	Equal(t, false, "AAA", "AAA", "password")
	Equal(t, false, "AAA", "BBB", "password")
	Equal(t, true, 4.5, 4.4)
}

func Test_NotEqual(t *testing.T) {
	NotEqual(t, false, 'a', 'a')
	NotEqual(t, false, 'a', 'b')
	NotEqual(t, false, 42, 42)
	NotEqual(t, false, 42, 0)
	NotEqual(t, false, "AAA", "AAA")
	NotEqual(t, false, "AAA", "BBB")
	NotEqual(t, true, 4.5, 4.4)
}

func Test_EqualSlice(t *testing.T) {
	EqualSlice(t, false, []int{1, 1}, []int{1, 1}, "slice")
	EqualSlice(t, false, []int{1, 1}, []int{1, 1, 1}, "slice")
	EqualSlice(t, false, []int{1, 1}, []int{1, 2}, "slice")
}

func Test_Error(t *testing.T) {
	var err error
	Error(t, false, err, "read file")
	err = fmt.Errorf("ERROR")
	Error(t, false, err, "read file")
}

func Test_NoError(t *testing.T) {
	var err error
	NoError(t, false, err)
	err = fmt.Errorf("ERROR")
	NoError(t, false, err)
}

func Test_Greater(t *testing.T) {
	Greater(t, false, 1, 0, "result")
	Greater(t, false, 0.0, 0.1, "result")
}

func Test_Less(t *testing.T) {
	Less(t, false, 1, 0)
	Less(t, false, 0.0, 0.1)
}

func Test_Nil(t *testing.T) {
	var err error
	Nil(t, false, err, "err")
	err = fmt.Errorf("ERROR")
	Nil(t, false, err, "err")
}

func Test_NotNil(t *testing.T) {
	var err error
	NotNil(t, false, err, "err")
	err = fmt.Errorf("ERROR")
	NotNil(t, false, err, "err")
}
