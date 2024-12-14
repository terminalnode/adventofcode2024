package main

import (
	"bytes"
	"fmt"
	"github.com/terminalnode/adventofcode2024/common"
	"github.com/terminalnode/adventofcode2024/common/util"
	"log"
)

const maxX = 100
const maxY = 102
const horizontalLine = maxY / 2
const verticalLine = maxX / 2

func main() {
	common.Setup(14, part1, part2)
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

func part2(
	input string,
) string {
	robots, err := parseRobots(input)
	if err != nil {
		return fmt.Sprintf("Failed to parse robots: %v", err)
	}

	// It's not going to be be very early, so lets save some time by skipping 5k
	for i := 5000; i < 15_000; i++ {
		printPositions(i, getNewPositions(robots, i))
	}

	return "kubectl logs -l day=14 -f | tee -a log.txt, then grep for a bunch of # :-)"
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

func printPositions(
	steps int,
	positions []util.Coordinate,
) {
	m := make([][]byte, maxY+1)
	for y := range m {
		m[y] = bytes.Repeat([]byte{'.'}, maxX+1)
	}

	for _, p := range positions {
		m[p.Y][p.X] = '#'
	}

	log.Println("Matrix after step ", steps)
	for _, r := range m {
		log.Println(string(r))
	}
	log.Println("------------------")
}
