package main

import (
	"fmt"
	"github.com/terminalnode/adventofcode2024/common"
	"math"
)

func main() {
	common.Setup(16, part1, nil)
}

func part1(
	input string,
) string {
	i, err := parse(input)
	if err != nil {
		return fmt.Sprintf("Failed to parse input: %v", err)
	}

	// Initialize visited set with end point
	set := initializeVisitedSet(i)
	loop(initialReindeer(i.s), i, set)

	return fmt.Sprintf("Cheapest path: %v", lowestSoFar(i, set))
}

func loop(
	r reindeer,
	i parsedInput,
	set visitedSet,
) {
	if r.visitAndCheckIfDead(set) {
		return
	}

	lowest := lowestSoFar(i, set)
	if lowest != 0 && r.score > lowest {
		return
	}

	if r.p.Equals(i.e) {
		return
	}

	loop(r.turnClockwise(), i, set)
	loop(r.turnCounterClockwise(), i, set)
	fwd, err := r.forward(i.m)
	if err == nil {
		loop(fwd, i, set)
	}

	return
}

func lowestSoFar(
	i parsedInput,
	set visitedSet,
) int {
	lowest := 0
	for _, n := range set[i.e.Y][i.e.X] {
		if lowest == 0 || n < lowest {
			lowest = n
		}
	}
	return lowest
}

func initializeVisitedSet(
	i parsedInput,
) visitedSet {
	set := make(visitedSet)
	for y, line := range i.m {
		set[y] = make(map[intX]map[intDirection]intScore)
		for x, _ := range line {
			set[y][x] = make(map[intDirection]intScore)
			set[y][x][North] = math.MaxInt
			set[y][x][South] = math.MaxInt
			set[y][x][West] = math.MaxInt
			set[y][x][East] = math.MaxInt
		}
	}
	return set
}
