package basic

import "testing"

func TestAddOne(t *testing.T) {
	var (
		input  = 1
		output = 2
	)

	actual := AddOne(1)

	if actual != output {
		t.Errorf("AddOne(%d), output = %d, actual = %d", input, output, actual)
	}
}
