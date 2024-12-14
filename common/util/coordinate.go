package util

import "fmt"

type Coordinate struct {
	X int
	Y int
}

type Direction = func(c Coordinate) Coordinate

func In2DArray[T any](
	c Coordinate,
	m [][]T,
) bool {
	invalid := len(m) == 0 || len(m[0]) == 0 ||
		c.X < 0 || c.X >= len(m[0]) ||
		c.Y < 0 || c.Y >= len(m)
	return !invalid
}

func (c Coordinate) Add(x int, y int) Coordinate {
	return Coordinate{X: c.X + x, Y: c.Y + y}
}

func (c Coordinate) Multiply(x int, y int) Coordinate {
	return Coordinate{X: c.X * x, Y: c.Y * y}
}

func (c Coordinate) Modulo(x int, y int) Coordinate {
	return Coordinate{X: c.X % x, Y: c.Y % y}
}

func (c Coordinate) PositiveModulo(x int, y int) Coordinate {
	// This is similar to how modulo works in Python
	// See: https://stackoverflow.com/questions/13794171/how-to-make-the-mod-of-a-negative-number-to-be-positive/13794192

	if x < 0 {
		x = -x
	}
	if y < 0 {
		y = -y
	}
	return Coordinate{
		X: ((c.X % x) + x) % x,
		Y: ((c.Y % y) + y) % y,
	}
}

func (c Coordinate) North() Coordinate {
	return Coordinate{c.X, c.Y - 1}
}

func (c Coordinate) East() Coordinate {
	return Coordinate{c.X + 1, c.Y}
}

func (c Coordinate) South() Coordinate {
	return Coordinate{c.X, c.Y + 1}
}

func (c Coordinate) West() Coordinate {
	return Coordinate{c.X - 1, c.Y}
}

func (c Coordinate) NorthWest() Coordinate {
	return Coordinate{c.X - 1, c.Y - 1}
}

func (c Coordinate) NorthEast() Coordinate {
	return Coordinate{c.X + 1, c.Y - 1}
}

func (c Coordinate) SouthWest() Coordinate {
	return Coordinate{c.X - 1, c.Y + 1}
}

func (c Coordinate) SouthEast() Coordinate {
	return Coordinate{c.X + 1, c.Y + 1}
}

func (c Coordinate) String() string {
	return fmt.Sprintf("(%d,%d)", c.X, c.Y)
}
