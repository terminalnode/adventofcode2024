package main

import (
	"fmt"
	"github.com/terminalnode/adventofcode2024/common"
)

func main()                     { common.Setup(13, part1, part2) }
func part1(input string) string { return solve(input, false) }
func part2(input string) string { return solve(input, true) }

func solve(
	input string,
	part2 bool,
) string {
	problems, err := parseProblems(input, part2)
	if err != nil {
		return fmt.Sprintf("Failed to parse input: %v", err)
	}

	sum := int64(0)
	for _, p := range problems {
		// Apparently this is an application of some Cramer's rule. ¯\_(ツ)_/¯
		det := p.aX*p.bY - p.aY*p.bX
		a := (p.goalX*p.bY - p.goalY*p.bX) / det
		b := (p.aX*p.goalY - p.aY*p.goalX) / det

		x := a*p.aX + b*p.bX
		y := a*p.aY + b*p.bY
		if x == p.goalX && y == p.goalY {
			sum += 3*a + b
		}
	}

	return fmt.Sprintf("Cost is %d tokens to solve these %d claw machines", sum, len(problems))
}
