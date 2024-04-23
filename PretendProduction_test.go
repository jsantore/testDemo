package main

import (
	"testing"
)

func TestCalcDistance(t *testing.T) {
	test_data := []struct {
		input_x1 float64
		inputy1  float64
		input_x2 float64
		input_y2 float64
		expected float64
	}{
		{0, 0, 3, 4, 5},
		{1, 1, 1, 1, 0},
		{-1, -1, -1, -1, 0},
		{},
	}
	for _, datum := range test_data {
		result := calcDistance(datum.input_x1, datum.inputy1,
			datum.input_x2, datum.input_y2)
		if result != datum.expected {
			t.Error("Expected ", datum.expected, " got ", result)
		}
	}

}
