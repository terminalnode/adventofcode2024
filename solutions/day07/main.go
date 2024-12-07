package main

import (
	"fmt"
	"github.com/terminalnode/adventofcode2024/common"
	"strconv"
	"strings"
)

type equation = struct {
	Test int
	Ops  []int
}

func main() {
	common.Setup(7, part1, nil)
}

func part1(
	input string,
) string {
	equations, err := parseEquations(input)
	if err != nil {
		return fmt.Sprintf("Failed to parse input: %v", err)
	}

	sum := 0
	for _, eq := range equations {
		if validateEquation(eq.Ops[0], eq.Test, eq.Ops[1:]) {
			sum += eq.Test
		}
	}

	return fmt.Sprintf("Sum of all OK tests: %d", sum)
}

func parseEquations(
	input string,
) ([]equation, error) {
	lines := strings.Split(input, "\n")
	equations := make([]equation, len(lines))
	for i, rawEq := range lines {
		testSplitOps := strings.Split(rawEq, ": ")
		if len(testSplitOps) != 2 {
			return equations, fmt.Errorf("splitting %v by ': ' resulted in %d pieces (should be two)", rawEq, len(testSplitOps))
		}

		test, err := strconv.ParseInt(testSplitOps[0], 10, 0)
		if err != nil {
			return equations, err
		}

		splitOps := strings.Split(testSplitOps[1], " ")
		ops := make([]int, len(splitOps))
		for opI, rawOp := range splitOps {
			op, err := strconv.ParseInt(rawOp, 10, 0)
			if err != nil {
				return equations, err
			}
			ops[opI] = int(op)
		}

		equations[i] = equation{int(test), ops}
	}

	return equations, nil
}

func validateEquation(
	currentValue int,
	targetValue int,
	operators []int,
) bool {
	if currentValue > targetValue {
		return false
	} else if len(operators) == 0 {
		return currentValue == targetValue
	}
	next := operators[0]
	mul := validateEquation(currentValue*next, targetValue, operators[1:])
	add := validateEquation(currentValue+next, targetValue, operators[1:])
	return mul || add
}
