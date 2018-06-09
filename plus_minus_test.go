package main

import "testing"

func TestStringConversion(t *testing.T) {
	input := []string{"1", "2", "3"}
	var output = transformInputToIntegers(input)
	expected := []int{1, 2, 3}
	if output[0] != expected[0] || output[1] != expected[1] || output[2] != expected[2] {
		t.Errorf("Transformed int slice was incorrect, got: %d, want: %d", output, expected)
	}
}
