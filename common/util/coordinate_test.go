package util

import (
	"fmt"
	"testing"
)

func TestCoordinateDirections(t *testing.T) {
	origin := Coordinate{X: 0, Y: 0}

	tests := []struct {
		name     string
		actual   Coordinate
		expected Coordinate
	}{
		{"North", origin.North(), Coordinate{X: 0, Y: -1}},
		{"NorthEast", origin.NorthEast(), Coordinate{X: 1, Y: -1}},
		{"NorthWest", origin.NorthWest(), Coordinate{X: -1, Y: -1}},
		{"South", origin.South(), Coordinate{X: 0, Y: 1}},
		{"SouthEast", origin.SouthEast(), Coordinate{X: 1, Y: 1}},
		{"SouthWest", origin.SouthWest(), Coordinate{X: -1, Y: 1}},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			if test.expected != test.actual {
				t.Errorf("expected %v but got %v", test.expected, test.actual)
			}
		})
	}
}

func TestCoordinate_String(t *testing.T) {
	tests := []struct {
		c        Coordinate
		expected string
	}{
		{c: Coordinate{X: 123, Y: 456}, expected: "(123,456)"},
		{c: Coordinate{X: -1741, Y: 917}, expected: "(-1741,917)"},
		{c: Coordinate{X: 184701, Y: -10471}, expected: "(184701,-10471)"},
	}

	for _, test := range tests {
		name := fmt.Sprintf("Coordinate{X:%d, Y:%d}.String()", test.c.X, test.c.Y)
		actual := test.c.String()

		t.Run(name, func(t *testing.T) {
			if test.expected != actual {
				t.Errorf("expected %v but got %v", test.expected, actual)
			}
		})

		name = fmt.Sprintf("Formatted print with %%s of Coordinate{X:%d, Y:%d}", test.c.X, test.c.Y)
		actual = fmt.Sprintf("%s", test.c)
		t.Run(name, func(t *testing.T) {
			if test.expected != test.c.String() {
				t.Errorf("expected %v but got %v", test.expected, actual)
			}
		})
	}
}

func TestIn2DArray(t *testing.T) {
	matrix := [][]int{{0, 0}, {0, 0}}
	tests := []struct {
		c        Coordinate
		expected bool
	}{
		// Top left corner
		{Coordinate{X: -1, Y: -1}, false},
		{Coordinate{X: +0, Y: -1}, false},
		{Coordinate{X: -1, Y: +0}, false},
		{Coordinate{X: +0, Y: +0}, true},

		// Top right corner
		{Coordinate{X: +1, Y: -1}, false},
		{Coordinate{X: +2, Y: -1}, false},
		{Coordinate{X: +1, Y: +0}, true},
		{Coordinate{X: +2, Y: +0}, false},

		// Bottom left corner
		{Coordinate{X: -1, Y: +1}, false},
		{Coordinate{X: +0, Y: +1}, true},
		{Coordinate{X: -1, Y: +2}, false},
		{Coordinate{X: +0, Y: +2}, false},

		// Bottom right corner
		{Coordinate{X: +1, Y: +1}, true},
		{Coordinate{X: +2, Y: +1}, false},
		{Coordinate{X: +1, Y: +2}, false},
		{Coordinate{X: +2, Y: +2}, false},
	}

	for _, test := range tests {
		name := fmt.Sprintf("%v in 2D matrix", test.c)
		t.Run(name, func(t *testing.T) {
			actual := In2DArray(test.c, matrix)
			if test.expected != actual {
				t.Errorf("expected In2DArray(test.c, matrix) == %v but got %v", test.expected, actual)
			}
		})
	}
}
