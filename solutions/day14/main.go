package main

import (
	"fmt"
	"github.com/terminalnode/adventofcode2024/common"
)

func main() {
	common.Setup(14, part1, nil)
}

func part1(
	input string,
) string {
	return solve(input, 100, 100, 102)
}

func solve(
	input string,
	seconds int,
	maxX int,
	maxY int,
) string {
	// Lines assume that the board is an uneven number of tiles
	horizontalLine := maxY / 2
	verticalLine := maxX / 2

	robots, err := parseRobots(input)
	if err != nil {
		return fmt.Sprintf("Failed to parse robots: %v", err)
	}

	q1, q2, q3, q4 := 0, 0, 0, 0
	for _, r := range robots {
		move := r.move.Multiply(seconds, seconds)
		newPos := r.init.Add(move.X, move.Y).PositiveModulo(maxX+1, maxY+1)

		if newPos.X < verticalLine {
			if newPos.Y < horizontalLine {
				q1++
			} else if newPos.Y > horizontalLine {
				q2++
			}
		} else if newPos.X > verticalLine {
			if newPos.Y < horizontalLine {
				q3++
			} else if newPos.Y > horizontalLine {
				q4++
			}
		}
	}

	safetyFactor := q1 * q2 * q3 * q4
	return fmt.Sprintf("After 100 seconds, area has a safety factor of %d", safetyFactor)
}
