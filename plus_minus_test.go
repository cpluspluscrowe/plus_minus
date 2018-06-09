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

func TestGetSignCounts(t *testing.T) {
	input := []int{0, 0, 0, 1, -1, -2}
	var positive_count, zero_count, negative_count int = getSignCounts(input)
	if positive_count != 1 {
		t.Errorf("Incorrect positive_count, got: %d, want: %d", positive_count, 1)
	}
	if zero_count != 3 {
		t.Errorf("Incorrect zero_count, got: %d, want: %d", zero_count, 3)
	}
	if negative_count != 2 {
		t.Errorf("Incorrect negative_count, got: %d, want: %d", negative_count, 2)
	}
}
