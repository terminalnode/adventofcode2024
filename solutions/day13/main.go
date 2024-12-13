package main

import (
	"fmt"
	"github.com/terminalnode/adventofcode2024/common"
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

	sum := int64(0)
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
) int64 {
	smallest := int64(math.MaxInt64)

	for timesA := int64(0); true; timesA++ {
		cX := timesA * p.aX
		cY := timesA * p.aY
		if cX > p.goalX || cY > p.goalY {
			break
		}

		minMoves := findMinTokenCostAfterA(cX, cY, timesA, p)
		if minMoves != -1 && minMoves < smallest {
			log.Printf("New smallest! %d", minMoves)
			smallest = minMoves
		}
	}

	if smallest == math.MaxInt64 {
		return -1
	}
	return smallest
}

func findMinTokenCostAfterA(
	cX int64,
	cY int64,
	timesA int64,
	p problem,
) int64 {
	for timesB := int64(0); true; timesB++ {
		newX := timesB*p.bX + cX
		newY := timesB*p.bY + cY

		if newX > p.goalX || newY > p.goalY {
			break
		} else if newX == p.goalX && newY == p.goalY {
			return timesA*3 + timesB
		}
	}

	return -1
}
