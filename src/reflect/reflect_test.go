package reflect

import (
	"testing"

	"github.com/goltsev/image/src/resize"
)

func TestFNName(t *testing.T) {
	cs := []struct {
		name     string
		fn       interface{}
		expected string
	}{
		{
			"nil",
			nil,
			"",
		},
		{
			"resizeNaive",
			resize.Naive,
			"github.com/goltsev/image/src/resize.Naive",
		},
		{
			"anonymous",
			func() {},
			"github.com/goltsev/image/src/reflect.TestFNName.func1",
		},
	}

	for _, c := range cs {
		t.Run(c.name, func(t *testing.T) {
			if got := FNName(c.fn); got != c.expected {
				t.Errorf("got: %v; expected: %v;\n", got, c.expected)
			}
		})
	}
}

func TestTrimFN(t *testing.T) {
	name := "github.com/goltsev/image/src.resizeNaive"
	expected := "resizeNaive"
	got := TrimFN(name)
	if got != expected {
		t.Errorf("got: %v; expected: %v;\n", got, expected)
	}
}
