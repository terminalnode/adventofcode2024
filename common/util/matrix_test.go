package util

import (
	"fmt"
	"testing"
)

func c(x int, y int) Coordinate {
	return Coordinate{X: x, Y: y}
}

var rawMatrix = [][]Coordinate{
	{c(0, 0), c(1, 0), c(2, 0), c(3, 0)},
	{c(0, 1), c(1, 1), c(2, 1), c(3, 1)},
	{c(0, 2), c(1, 2), c(2, 2), c(3, 2)},
	{c(0, 3), c(1, 3), c(2, 3), c(3, 3)},
	{c(0, 4), c(1, 4), c(2, 4), c(3, 4)},
}
var matrix = Matrix[Coordinate]{
	RawMatrix: rawMatrix,
	MaxX:      3,
	MaxY:      4,
}

func TestNewMatrixFromRows(t *testing.T) {
	t.Run("With good source matrix", func(t *testing.T) {
		testMatrix, err := NewMatrixFromRows(rawMatrix)

		if err != nil {
			t.Errorf("Method threw unexpected error: %v", err)
		}

		if testMatrix.MaxX != matrix.MaxX {
			t.Errorf("Expected MaxX=%d, but was %d", matrix.MaxX, testMatrix.MaxX)
		}

		if testMatrix.MaxY != matrix.MaxY {
			t.Errorf("Expected MaxY=%d, but was %d", matrix.MaxY, testMatrix.MaxY)
		}

		for y, tmRow := range testMatrix.RawMatrix {
			for x, actual := range tmRow {
				expected := rawMatrix[y][x]
				if actual != expected {
					t.Errorf("Expected (X=%d,Y=%d) to be %v, but was %v", x, y, expected, actual)
				}
			}
		}
	})

	t.Run("With long second row matrix", func(t *testing.T) {
		raw := [][]int{{1, 2}, {1, 2, 3}}
		testMatrix, err := NewMatrixFromRows(raw)
		errorMsg := "matrix[1] length (3) differs from matrix[0] (2)"

		if err != nil && err.Error() != errorMsg {
			t.Errorf("Expected error '%s' but got %v", errorMsg, err)
		}

		if testMatrix.MaxX != 0 && testMatrix.MaxY != 0 && len(testMatrix.RawMatrix) != 0 {
			t.Errorf("Expected zero-matrix in response, but was %q", testMatrix)
		}
	})

	t.Run("With short third row matrix", func(t *testing.T) {
		raw := [][]int{{1, 2, 3}, {1, 2, 3}, {1, 2}}
		testMatrix, err := NewMatrixFromRows(raw)
		errorMsg := "matrix[2] length (2) differs from matrix[0] (3)"

		if err != nil && err.Error() != errorMsg {
			t.Errorf("Expected error '%s' but got %v", errorMsg, err)
		}

		if testMatrix.MaxX != 0 && testMatrix.MaxY != 0 && len(testMatrix.RawMatrix) != 0 {
			t.Errorf("Expected zero-matrix in response, but was %q", testMatrix)
		}
	})
}

func TestNewCharMatrix(t *testing.T) {
	t.Run("With good source matrix", func(t *testing.T) {
		s := "1234\n5678\n9abc"
		sl := [][]uint8{
			{'1', '2', '3', '4'},
			{'5', '6', '7', '8'},
			{'9', 'a', 'b', 'c'},
		}
		testMatrix, err := NewCharMatrix(s)

		if err != nil {
			t.Errorf("Method threw unexpected error: %v", err)
		}

		if testMatrix.MaxX != 3 {
			t.Errorf("Expected MaxX=3, but was %d", testMatrix.MaxX)
		}

		if testMatrix.MaxY != 2 {
			t.Errorf("Expected MaxY=2, but was %d", testMatrix.MaxY)
		}

		for y, tmRow := range testMatrix.RawMatrix {
			for x, actual := range tmRow {
				expected := sl[y][x]
				if actual != expected {
					t.Errorf("Expected (X=%d,Y=%d) to be %v, but was %v", x, y, expected, actual)
				}
			}
		}
	})

	t.Run("With long second row matrix", func(t *testing.T) {
		s := "12\n123"
		testMatrix, err := NewCharMatrix(s)
		errorMsg := "matrix[1] length (3) differs from matrix[0] (2)"

		if err != nil && err.Error() != errorMsg {
			t.Errorf("Expected error '%s' but got %v", errorMsg, err)
		}

		if testMatrix.MaxX != 0 && testMatrix.MaxY != 0 && len(testMatrix.RawMatrix) != 0 {
			t.Errorf("Expected zero-matrix in response, but was %q", testMatrix)
		}
	})

	t.Run("With short third row matrix", func(t *testing.T) {
		s := "123\n123\n12"
		testMatrix, err := NewCharMatrix(s)
		errorMsg := "matrix[2] length (2) differs from matrix[0] (3)"

		if err != nil && err.Error() != errorMsg {
			t.Errorf("Expected error '%s' but got %v", errorMsg, err)
		}

		if testMatrix.MaxX != 0 && testMatrix.MaxY != 0 && len(testMatrix.RawMatrix) != 0 {
			t.Errorf("Expected zero-matrix in response, but was %q", testMatrix)
		}
	})
}

func TestMatrix_IsInMatrix(t *testing.T) {
	for x := -1; x <= matrix.MaxX+1; x++ {
		for y := -1; y <= matrix.MaxY+1; y++ {
			isIn := matrix.IsInMatrix(x, y)
			invalid := x < 0 || y < 0 || x > matrix.MaxX || y > matrix.MaxY

			if !isIn && !invalid {
				t.Errorf("Expected (%d,%d) to yield true, but it didn't", x, y)
			}

			if isIn && (x < 0 || x > matrix.MaxX) {
				t.Errorf("X=%d should not be inside the matrix", x)
			}

			if isIn && (y < 0 || y > matrix.MaxY) {
				t.Errorf("Y=%d should not be inside the matrix", y)
			}
		}
	}
}

func TestMatrix_Get(t *testing.T) {
	coZero := Coordinate{}

	for x := -1; x <= matrix.MaxX+1; x++ {
		for y := -1; y <= matrix.MaxY+1; y++ {
			co := Coordinate{X: x, Y: y}
			out, err := matrix.Get(x, y)
			isIn := matrix.IsInMatrix(x, y)

			if isIn && err != nil {
				t.Errorf("Expected to get %q but was: %v", co, err)
			}

			if !isIn && err == nil {
				t.Errorf("Expected an error for %v", co)
			}

			if isIn && out != co {
				t.Errorf("Got output %v but expected %v", out, co)
			}

			if err != nil && out != coZero {
				t.Errorf("Got output %v but expected %v because err != nil", out, co)
			}
		}
	}
}

func TestMatrix_GetOrDefault(t *testing.T) {
	defaultValue := Coordinate{1337, 420}

	for x := -1; x <= matrix.MaxX+1; x++ {
		for y := -1; y <= matrix.MaxY+1; y++ {
			expected := Coordinate{X: x, Y: y}
			actual, err := matrix.GetOrDefault(x, y, defaultValue)
			isIn := matrix.IsInMatrix(x, y)

			if isIn {
				if actual != expected {
					t.Errorf("Expected %v, but was %v", expected, actual)
				}

				if err != nil {
					t.Errorf("Expected an error, but got nil")
				}
			} else {
				if actual != defaultValue {
					t.Errorf("Expected default value fallback (%v), but got %v", defaultValue, actual)
				}

				if err == nil {
					t.Errorf("Expected err to be nil, but was %v", err)
				}
			}
		}
	}
}

func TestMatrix_Set(t *testing.T) {
	raw := [][]int{{1, 1}, {1, 1}}
	m, creationErr := NewMatrixFromRows(raw)
	coords := []Coordinate{c(0, 0), c(0, 1), c(1, 0), c(1, 1)}

	for _, point := range coords {
		name := fmt.Sprintf("Changing at point %v", point)
		t.Run(name, func(t *testing.T) {
			if creationErr != nil {
				t.Error("Initialization of matrix failed")
			}

			curV, err := m.Get(point.X, point.Y)
			if err != nil {
				t.Error("Failed to get current value")
			}

			expectedNewV := curV * 10
			err = m.Set(point.X, point.Y, expectedNewV)
			if err != nil {
				t.Error("Failed to set value")
			}

			actualNewV, err := m.Get(point.X, point.Y)
			if err != nil {
				t.Error("Failed to get new value")
			}

			if actualNewV != expectedNewV {
				t.Errorf("Expected new value %d, but was %d", expectedNewV, actualNewV)
			}
		})

		t.Run("Setting point outside of matrix", func(t *testing.T) {
			err := m.Set(2, 1, 1337)
			if err == nil {
				t.Error("Expected error from setting (2,1), but got nil")
			}
		})
	}
}

func TestCountInMatrix(t *testing.T) {
	m, err := NewMatrixFromRows([][]int{{1, 2, 3, 1}, {2, 3, 3, 1}, {3, 1, 2, 1}})
	count1 := 5
	count2 := 3
	count3 := 4
	count4 := 0

	if err != nil {
		t.Error("Failed to construct matrix")
	}
	var count int

	count = CountInMatrix(m, 1)
	if count != count1 {
		t.Errorf("Expected %d number of ones, got %d", count1, count)
	}

	count = CountInMatrix(m, 2)
	if count != count2 {
		t.Errorf("Expected %d number of ones, got %d", count2, count)
	}

	count = CountInMatrix(m, 3)
	if count != count3 {
		t.Errorf("Expected %d number of ones, got %d", count3, count)
	}

	count = CountInMatrix(m, 4)
	if count != count4 {
		t.Errorf("Expected %d number of ones, got %d", count4, count)
	}
}
