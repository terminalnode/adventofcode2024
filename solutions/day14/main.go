package main

import (
	"bytes"
	"github.com/terminalnode/adventofcode2024/common"
	"github.com/terminalnode/adventofcode2024/common/util"
	"regexp"
)

const maxX = 100
const maxY = 102
const horizontalLine = maxY / 2
const verticalLine = maxX / 2

func main() {
	common.Setup(14, part1, part2)
}

func part1(
	input util.AocInput,
) (util.AocSolution, util.AocError) {
	robots, err := parseRobots(input.Input)
	if err != nil {
		return util.NewAocError(err.Error(), util.InputParsingError)
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

	return util.FormatAocSolution("After 100 seconds, area has a safety factor of %d", q1*q2*q3*q4)
}

func part2(
	input util.AocInput,
) (util.AocSolution, util.AocError) {
	robots, err := parseRobots(input.Input)
	if err != nil {
		return util.NewAocError(err.Error(), util.InputParsingError)
	}

	// The image we're looking for has a frame, and all non-image parts seem highly irregular.
	// So the assumption is that this shouldn't appear outside the sought frame.
	// The original image was found by saving logs to a file and searching for hashes too, just less specific.
	r := regexp.MustCompile(`.###############################.`)

	// We could very reasonably skip ~5000 steps it seems, as all answers on the sub is 5k+,
	// but this already runs in less than 20 seconds so no reason to cheat.
	found := false
	var step int
	for step = 1; step < 15_000; step++ {
		np := getNewPositions(robots, step)
		m := make([][]byte, maxY+1)
		for y := range m {
			m[y] = bytes.Repeat([]byte{'.'}, maxX+1)
		}
		for _, p := range np {
			m[p.Y][p.X] = '#'
		}

		for _, row := range m {
			if r.Match(row) {
				found = true
				break
			}
		}

		if found {
			break
		}
	}

	return util.FormatAocSolution("Answer is probably %d", step)
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
