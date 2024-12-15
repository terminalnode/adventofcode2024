package main

import (
	"fmt"
	"github.com/terminalnode/adventofcode2024/common/util"
	"strings"
)

type warehouseMatrix = [][]int32
type moveList = []util.Direction
type robotPosition = util.Coordinate

type problem struct {
	robot     robotPosition
	warehouse warehouseMatrix
	moves     moveList
}

const (
	Wall     = '#'
	Box      = 'O'
	Robot    = '@'
	Ground   = '.'
	LeftBox  = '['
	RightBox = ']'
)

func parse(
	input string,
	makeWide bool,
) (problem, error) {
	whRaw, mRaw, err := splitInput(input)
	if err != nil {
		return problem{}, err
	}

	wh, r, err := parseWarehouse(whRaw, makeWide)

	m, err := parseMoves(mRaw)
	if err != nil {
		return problem{}, err
	}

	return problem{
		robot:     r,
		warehouse: wh,
		moves:     m,
	}, nil
}

func splitInput(
	input string,
) (string, string, error) {
	split := strings.Split(input, "\n\n")
	if len(split) != 2 {
		return "", "", fmt.Errorf("expected input to have two parts, but was %d", len(split))
	}
	return split[0], split[1], nil
}

func parseWarehouse(
	wh string,
	makeWide bool,
) (warehouseMatrix, util.Coordinate, error) {
	lines := strings.Split(wh, "\n")
	out := make(warehouseMatrix, len(lines))
	robot := util.Coordinate{}
	foundRobot := false

	for y, line := range lines {
		var row []int32
		if makeWide {
			row = make([]int32, 0, len(line)*2)
		} else {
			row = make([]int32, 0, len(line))
		}

		for x, ch := range line {
			switch ch {
			case Wall:
				if makeWide {
					row = append(row, Wall)
				}
				row = append(row, Wall)
			case Box:
				if makeWide {
					row = append(row, LeftBox, RightBox)
				} else {
					row = append(row, Box)
				}
			case Ground:
				if makeWide {
					row = append(row, Ground)
				}
				row = append(row, Ground)
			case Robot:
				foundRobot = true
				if makeWide {
					row = append(row, Ground, Ground)
					robot = util.Coordinate{X: 2 * x, Y: y}
				} else {
					row = append(row, Ground)
					robot = util.Coordinate{X: x, Y: y}
				}
			}
		}
		out[y] = row
	}

	if !foundRobot {
		return out, robot, fmt.Errorf("no robot in input")
	}

	return out, robot, nil
}

func parseMoves(
	m string,
) (moveList, error) {
	ms := make(moveList, 0, len(m))
	for _, ch := range m {
		if ch == '\n' {
			// As per instructions, ignore new lines
			continue
		}

		switch ch {
		case '^':
			ms = append(ms, util.Coordinate.North)
		case 'v':
			ms = append(ms, util.Coordinate.South)
		case '>':
			ms = append(ms, util.Coordinate.East)
		case '<':
			ms = append(ms, util.Coordinate.West)
		default:
			return ms, fmt.Errorf("invalid direction '%c'", ch)
		}
	}
	return ms, nil
}
