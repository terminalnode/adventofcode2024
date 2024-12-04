package util

import (
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
