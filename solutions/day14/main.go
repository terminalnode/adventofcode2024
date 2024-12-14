package main

import (
	"fmt"
	"github.com/terminalnode/adventofcode2024/common"
	"github.com/terminalnode/adventofcode2024/common/util"
)

const maxX = 100
const maxY = 102
const horizontalLine = maxY / 2
const verticalLine = maxX / 2

func main() {
	common.Setup(14, part1, nil)
}

func part1(
	input string,
) string {
	robots, err := parseRobots(input)
	if err != nil {
		return fmt.Sprintf("Failed to parse robots: %v", err)
	}

	q1, q2, q3, q4 := 0, 0, 0, 0
	for _, p := range getNewPositions(robots, 100) {
		if p.X < verticalLine {
			if p.Y < horizontalLine {
				q1++
			} else if p.Y > horizontalLine {
				q2++
			}
		} else if p.X > verticalLine {
			if p.Y < horizontalLine {
				q3++
			} else if p.Y > horizontalLine {
				q4++
			}
		}
	}

	return fmt.Sprintf("After 100 seconds, area has a safety factor of %d", q1*q2*q3*q4)
}

func getNewPositions(
	robots []robot,
	seconds int,
) []util.Coordinate {
	out := make([]util.Coordinate, len(robots))
	for i, r := range robots {
		move := r.move.Multiply(seconds, seconds)
		out[i] = r.init.Add(move.X, move.Y).PositiveModulo(maxX+1, maxY+1)

	}
	return out
}
