package util

type Coordinate struct {
	X int
	Y int
}

type Direction = func(c Coordinate) Coordinate

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
