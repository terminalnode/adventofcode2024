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

	newR, err := r.forward(i.m)
	if err == nil {
		loop(newR, i, set)
	}

	newR, err = r.turnClockwise().forward(i.m)
	if err == nil {
		loop(newR, i, set)
	}

	newR, err = r.turnClockwise().turnClockwise().forward(i.m)
	if err == nil {
		loop(newR, i, set)
	}

	newR, err = r.turnCounterClockwise().forward(i.m)
	if err == nil {
		loop(newR, i, set)
	}

	return
}

func lowestSoFar(
	i parsedInput,
	set visitedSet,
) int {
	return set[i.e.Y][i.e.X]
}

func initializeVisitedSet(
	i parsedInput,
) visitedSet {
	set := make(visitedSet)
	for y, line := range i.m {
		set[y] = make(map[intX]intScore)
		for x, _ := range line {
			set[y][x] = math.MaxInt
		}
	}
	return set
}
