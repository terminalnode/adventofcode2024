package main

import (
	"github.com/terminalnode/adventofcode2024/common/util"
)

var numStart = util.Coordinate{X: 2, Y: 3}
var dirStart = util.Coordinate{X: 2, Y: 0}

// Char to Pos on numeric keypad
func charToPosNum(
	c int,
) util.Coordinate {
	switch c {
	case '0':
		return util.Coordinate{X: 1, Y: 3}
	case 'A':
		return numStart
	default:
		y := 2 - ((c - '0' - 1) / 3)
		x := (c - '0' - 1) % 3
		return util.Coordinate{X: int(x), Y: int(y)}
	}
}

// Char to Pos on directional keypad
func charToPosDir(
	c int,
) util.Coordinate {
	switch c {
	case '^':
		return util.Coordinate{X: 1, Y: 0}
	case '<':
		return util.Coordinate{X: 0, Y: 1}
	case 'v':
		return util.Coordinate{X: 1, Y: 1}
	case '>':
		return util.Coordinate{X: 2, Y: 1}
	default:
		return util.Coordinate{X: 2, Y: 0}
	}
}

func shortestPath(
	start, end util.Coordinate,
	useNumPad bool,
) []int {
	var path []int
	dx := end.X - start.X
	dy := end.Y - start.Y
	pVert, pHor := buildPath(dx, dy)

	left := end.X < start.X
	var gap bool
	if useNumPad {
		gap = (start.Y == 3 && end.X == 0) || (start.X == 0 && end.Y == 3)
	} else {
		gap = (start.X == 0 && end.Y == 0) || (start.Y == 0 && end.X == 0)
	}

	// The gap is on the left side in both pads, so if we're moving left and would cross
	// the gap we have to avoid it. If not, apparently it's always best to move vertical first.
	if left != gap {
		path = append(path, pHor...)
		path = append(path, pVert...)
	} else {
		path = append(path, pVert...)
		path = append(path, pHor...)
	}
	path = append(path, 'A')
	return path
}

func buildPath(
	dx int,
	dy int,
) ([]int, []int) {
	adx := util.AbsInt(dx)
	ady := util.AbsInt(dy)
	pVert := make([]int, ady)
	pHor := make([]int, adx)

	for i := range adx {
		if dx < 0 {
			pHor[i] = '<'
		} else {
			pHor[i] = '>'
		}
	}

	for i := range ady {
		if dy < 0 {
			pVert[i] = '^'
		} else {
			pVert[i] = 'v'
		}
	}

	return pVert, pHor
}
