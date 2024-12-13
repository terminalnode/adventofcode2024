package main

import (
	"fmt"
	"github.com/terminalnode/adventofcode2024/common"
	"github.com/terminalnode/adventofcode2024/common/util"
	"log"
	"math"
)

func main() {
	common.Setup(13, part1, nil)
}

func part1(
	input string,
) string {
	problems, err := parseProblems(input)
	if err != nil {
		return fmt.Sprintf("Failed to parse input: %v", err)
	}

	sum := 0
	for _, p := range problems {
		minTokenCost := findMinTokenCost(p)
		if minTokenCost != -1 {
			sum += minTokenCost
		}
	}

	return fmt.Sprintf("At least %d tokens to solve these %d claw machines", sum, len(problems))
}

func findMinTokenCost(
	p problem,
) int {
	smallest := math.MaxInt

	for timesA := 0; true; timesA++ {
		c := util.Coordinate{
			X: timesA * p.a.X,
			Y: timesA * p.a.Y,
		}
		if c.X > p.goal.X || c.Y > p.goal.Y {
			break
		}

		minMoves := findMinTokenCostAfterA(c, timesA, p)
		if minMoves != -1 && minMoves < smallest {
			log.Printf("New smallest! %d", minMoves)
			smallest = minMoves
		}
	}

	if smallest == math.MaxInt {
		return -1
	}
	return smallest
}

func findMinTokenCostAfterA(
	c util.Coordinate,
	timesA int,
	p problem,
) int {
	for timesB := 0; true; timesB++ {
		newC := util.Coordinate{
			X: timesB*p.b.X + c.X,
			Y: timesB*p.b.Y + c.Y,
		}

		if newC.X > p.goal.X || newC.Y > p.goal.Y {
			break
		} else if newC.X == p.goal.X && newC.Y == p.goal.Y {
			return timesA*3 + timesB
		}
	}

	return -1
}
