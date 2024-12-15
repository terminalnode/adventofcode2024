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
	WALL   = '#'
	BOX    = 'O'
	ROBOT  = '@'
	GROUND = '.'
)

func parse(
	input string,
) (problem, error) {
	whRaw, mRaw, err := splitInput(input)
	if err != nil {
		return problem{}, err
	}

	wh, r, err := parseWarehouse(whRaw)

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
) (warehouseMatrix, util.Coordinate, error) {
	lines := strings.Split(wh, "\n")
	out := make(warehouseMatrix, len(lines))
	robot := util.Coordinate{}
	foundRobot := false

	for y, line := range lines {
		out[y] = []int32(line)
		for x, ch := range line {
			if ch != WALL && ch != ROBOT && ch != BOX && ch != GROUND {
				return out, robot, fmt.Errorf("invalid character in warehouse matrix: %c", ch)
			}

			if ch != ROBOT {
				continue
			}
			foundRobot = true
			robot = util.Coordinate{X: x, Y: y}
			out[y][x] = GROUND
		}
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

		newMove, err := charToDirection(ch)
		if err != nil {
			return ms, err
		}
		ms = append(ms, newMove)
	}
	return ms, nil
}

func charToDirection(
	ch int32,
) (util.Direction, error) {
	switch ch {
	case '^':
		return util.Coordinate.North, nil
	case 'v':
		return util.Coordinate.South, nil
	case '>':
		return util.Coordinate.East, nil
	case '<':
		return util.Coordinate.West, nil
	}
	return nil, fmt.Errorf("invalid direction '%c'", ch)
}
