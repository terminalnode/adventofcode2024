package main

import (
	"github.com/terminalnode/adventofcode2024/common"
	"github.com/terminalnode/adventofcode2024/common/util"
)

func main()                                                       { common.Setup(13, part1, part2) }
func part1(input util.AocInput) (util.AocSolution, util.AocError) { return solve(input, false) }
func part2(input util.AocInput) (util.AocSolution, util.AocError) { return solve(input, true) }

func solve(
	input util.AocInput,
	part2 bool,
) (util.AocSolution, util.AocError) {
	problems, err := parseProblems(input.Input, part2)
	if err != nil {
		return util.NewAocError(err.Error(), util.InputParsingError)
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

	return util.FormatAocSolution("Cost is %d tokens to solve these %d claw machines", sum, len(problems))
}
