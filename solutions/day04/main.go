package main

import (
	"github.com/terminalnode/adventofcode2024/common"
	"github.com/terminalnode/adventofcode2024/common/util"
)

type CharMatrix = util.CharMatrix
type Coordinate = util.Coordinate

func main() {
	common.Setup(4, part1, part2)
}

func part1(
	input util.AocInput,
) (util.AocSolution, util.AocError) {
	m, err := util.NewCharMatrix(input.Input)
	if err != nil {
		return util.NewAocError(err.Error(), util.InputParsingError)
	}

	count := 0
	for x := 0; x <= m.MaxX; x++ {
		for y := 0; y <= m.MaxY; y++ {
			start := Coordinate{X: x, Y: y}
			if !isThereAThisHere(m, start, 'X') {
				continue
			}

			count += searchXmas(m, start, Coordinate.North) +
				searchXmas(m, start, Coordinate.NorthEast) +
				searchXmas(m, start, Coordinate.NorthWest) +
				searchXmas(m, start, Coordinate.South) +
				searchXmas(m, start, Coordinate.SouthEast) +
				searchXmas(m, start, Coordinate.SouthWest) +
				searchXmas(m, start, Coordinate.East) +
				searchXmas(m, start, Coordinate.West)
		}
	}

	return util.FormatAocSolution("Number of XMAS: %d", count)
}

func part2(
	input util.AocInput,
) (util.AocSolution, util.AocError) {
	m, err := util.NewCharMatrix(input.Input)
	if err != nil {
		return util.NewAocError(err.Error(), util.InputParsingError)
	}
	count := 0
	for x := 1; x < m.MaxX; x++ {
		for y := 1; y < m.MaxY; y++ {
			start := Coordinate{X: x, Y: y}
			if !isThereAThisHere(m, start, 'A') {
				continue
			}
			nw := start.NorthWest()
			northWest, _ := m.Get(nw.X, nw.Y)
			if northWest == 'X' {
				continue
			}

			ne := start.NorthEast()
			northEast, _ := m.Get(ne.X, ne.Y)
			if northEast == 'X' {
				continue
			}

			sw := start.SouthWest()
			southWest, _ := m.Get(sw.X, sw.Y)
			if southWest == 'X' {
				continue
			}

			se := start.SouthEast()
			southEast, _ := m.Get(se.X, se.Y)
			if southEast == 'X' {
				continue
			}

			northM := northWest == 'M' && northEast == 'M' &&
				southWest == 'S' && southEast == 'S'
			southM := northWest == 'S' && northEast == 'S' &&
				southWest == 'M' && southEast == 'M'
			eastM := northWest == 'S' && southWest == 'S' &&
				northEast == 'M' && southEast == 'M'
			westM := northWest == 'M' && southWest == 'M' &&
				northEast == 'S' && southEast == 'S'

			if northM || southM || eastM || westM {
				count += 1
			}
		}
	}

	return util.FormatAocSolution("Number of X-shaped MAS: %d", count)
}

func searchXmas(
	m CharMatrix,
	start Coordinate,
	direction util.Direction,
) int {
	letterM := direction(start)
	letterA := direction(letterM)
	letterS := direction(letterA)

	itsXmas := isThereAThisHere(m, letterS, 'S') &&
		isThereAThisHere(m, letterA, 'A') &&
		isThereAThisHere(m, letterM, 'M')

	if itsXmas {
		return 1
	} else {
		return 0
	}
}

func isThereAThisHere(
	m CharMatrix,
	c Coordinate,
	expected uint8,
) bool {
	actual, err := m.Get(c.X, c.Y)
	return err == nil && actual == expected
}
