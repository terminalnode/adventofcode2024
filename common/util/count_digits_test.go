package util

import (
	"fmt"
	"testing"
)

func TestC(t *testing.T) {
	tests := []struct {
		input    int
		expected int
	}{
		// 0 - 9
		{0, 1}, {1, 1}, {2, 1}, {3, 1}, {4, 1},
		{5, 1}, {6, 1}, {7, 1}, {8, 1}, {9, 1},

		// 10-19
		{10, 2}, {11, 2}, {12, 2}, {13, 2}, {14, 2},
		{15, 2}, {16, 2}, {17, 2}, {18, 2}, {19, 2},

		// 100-109
		{100, 3}, {100, 3}, {100, 3}, {100, 3}, {100, 3},
		{100, 3}, {100, 3}, {100, 3}, {100, 3}, {100, 3},
	}

	for _, test := range tests {
		name := fmt.Sprintf("CountDigits(%d) should return %d", test.input, test.expected)
		t.Run(name, func(t *testing.T) {
			actual := CountDigits(test.input)
			if test.expected != actual {
				t.Errorf("expected %v but got %v", test.expected, actual)
			}
		})
	}
}
