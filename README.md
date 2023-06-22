Simple, lightweight library for go tests.

## Examples

### Equal, Greater, Less

```go
func Add(a,b int) int {
	return a + b
}

func TestAdd(t *testing.T) {
	x := Add(1,2)
	expect.Equal(t, "Add", 3, x)
	expect.Greater(t, "Add", 2, x)
	
	y := Add(x, 1)
	expect.Less(t, "Add", x, y)
}
```

### EqualSlice

```go
func TestSlices(t *testing.T) {
	
	x := []float32{0.1, 0.001}
	y := []float32{0.1, 0.001}
	expect.EqualSlice(t, "floats", x, y) // will pass
	
	a := []string{"a", "b"}
	b := []string{"b", "a"}
	expect.EqualSlice(t, "strings", a, b) // will fail
}
```

### Error, NoError, Nil, NotNil
```go
func TestErrors(t *testing.T) {
	var err error
	expect.NoError(t, "error", err)
	expect.Nil(t, "error", err)
	
	err = fmt.Errorf("error")
	expect.Error(t, "error", err)
	expect.NotNil(t, "error", err)
}
```

### Return value
All functions return a bool indicating whether test was passed or not.
```go
func TestExample(t *testing.T) {
	err := foo()
	ok := expect.NoError(err, "error from foo", err)
	if !ok {
		t.Fatal()
	}
}
```