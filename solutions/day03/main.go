package main

import (
	"fmt"
	"github.com/terminalnode/adventofcode2024/common"
	"github.com/terminalnode/adventofcode2024/common/util"
	"regexp"
	"strconv"
)

var r = regexp.MustCompile("(mul|do|don't)\\(((\\d+),(\\d+))?\\)")

func main() {
	common.Setup(3, part1, part2)
}

func part1(
	input util.AocInput,
) (util.AocSolution, util.AocError) {
	finds := r.FindAllSubmatch([]byte(input.Input), -1)
	sum := 0
	for _, match := range finds {
		verb := string(match[1])
		if verb != "mul" {
			continue
		}

		multiplied, err := mul(string(match[3]), string(match[4]))
		if err != nil {
			return util.NewAocError(fmt.Sprintf("Failed to parse multiplication of %q:\n%v\n", match, err), util.ParsingError)
		}
		sum += multiplied
	}
	return util.FormatAocSolution("Result: %d", sum)
}

func part2(
	input util.AocInput,
) (util.AocSolution, util.AocError) {
	finds := r.FindAllSubmatch([]byte(input.Input), -1)
	sum := 0

	enabled := true
	for _, match := range finds {
		verb := string(match[1])
		switch verb {
		case "mul":
			if enabled {
				multiplied, err := mul(string(match[3]), string(match[4]))
				if err != nil {
					return util.NewAocError(fmt.Sprintf("Failed to parse multiplication of %q:\n%v\n", match, err), util.ParsingError)
				}
				sum += multiplied
			}
		case "do":
			enabled = true
		case "don't":
			enabled = false
		}
	}

	return util.FormatAocSolution("Result: %d", sum)
}

func mul(sub1 string, sub2 string) (int, error) {
	first, err := strconv.ParseInt(sub1, 10, 0)
	if err != nil {
		return 0, err
	}

	second, err := strconv.ParseInt(sub2, 10, 0)
	if err != nil {
		return 0, err
	}

	return int(first * second), nil
}
