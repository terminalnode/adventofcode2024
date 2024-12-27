package main

import (
	"github.com/terminalnode/adventofcode2024/common"
	"github.com/terminalnode/adventofcode2024/common/util"
	"math"
)

func main() {
	common.Setup(16, part1, part2)
}

func part1(
	input util.AocInput,
) (util.AocSolution, util.AocError) {
	i, err := parse(input.Input)
	if err != nil {
		return util.NewAocError(err.Error(), util.InputParsingError)
	}

	// Initialize visited set with end point
	set := initializeVisitedSet(i)
	loop(initialReindeer(i.s), i, set, initializeWinningSet(i), -1)

	return util.FormatAocSolution("Cheapest path: %v", lowestSoFar(i, set))
}

func part2(
	input util.AocInput,
) (util.AocSolution, util.AocError) {
	i, err := parse(input.Input)
	if err != nil {
		return util.NewAocError(err.Error(), util.InputParsingError)
	}

	// Initialize visited set with end point
	visited := initializeVisitedSet(i)
	loop(initialReindeer(i.s), i, visited, initializeWinningSet(i), -1)

	// Run again with the correct answer, which should weed out
	// all paths that were the best before better paths were found.
	winners := initializeWinningSet(i)
	answer := lowestSoFar(i, visited)
	loop(initialReindeer(i.s), i, initializeVisitedSet(i), winners, answer)
	winners[i.e.Y][i.e.X] = true

	count := 0
	for _, ys := range winners {
		count += len(ys)
	}

	return util.FormatAocSolution("Best seats: %d", count)
}

func loop(
	r reindeer,
	i parsedInput,
	set visitedSet,
	winners winningSet,
	answer int,
) bool {
	isWinner := false
	if r.visitAndCheckIfDead(set, answer != -1) {
		return false
	}

	lowest := lowestSoFar(i, set)
	if answer == -1 && lowest != 0 && r.score > lowest {
		winners[r.p.Y][r.p.X] = true
		return false
	}

	if r.p.Equals(i.e) {
		if answer == -1 {
			return true
		} else {
			return r.score == answer
		}
	}

	newR, err := r.forward(i.m)
	if err == nil && loop(newR, i, set, winners, answer) {
		isWinner = true
	}

	newR, err = r.turnClockwise().forward(i.m)
	if err == nil && loop(newR, i, set, winners, answer) {
		isWinner = true
	}

	newR, err = r.turnCounterClockwise().forward(i.m)
	if err == nil && loop(newR, i, set, winners, answer) {
		isWinner = true
	}

	if isWinner {
		winners[r.p.Y][r.p.X] = true
	}

	return isWinner
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

func initializeWinningSet(
	i parsedInput,
) winningSet {
	set := make(winningSet)
	for y, _ := range i.m {
		set[y] = make(map[intX]bool)
	}
	return set
}
