package greeting

import "testing"

func Test_If_FirstNumber_is_Greater_Than_the_Second(t *testing.T) {
	var a int = 20
	var b int = 10

	result, err := SubTraction(a, b)
	if err != nil {
		t.Errorf("Error: %v", err)
	}

	expected := a - b
	if result != expected {
		t.Errorf("Expected result %d, but got %d", expected, result)
	}
}
