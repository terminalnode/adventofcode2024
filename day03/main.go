package main

import (
	"fmt"
	"github.com/terminalnode/adventofcode2024/common"
	"regexp"
	"strconv"
)

var r = regexp.MustCompile("(mul|do|don't)\\(((\\d+),(\\d+))?\\)")

func main() {
	common.Setup(3, part1, part2)
}

func part1(
	input string,
) string {
	finds := r.FindAllSubmatch([]byte(input), -1)
	sum := 0
	for _, match := range finds {
		verb := string(match[1])
		if verb != "mul" {
			continue
		}

		multiplied, err := mul(string(match[3]), string(match[4]))
		if err != nil {
			return fmt.Sprintf("Failed to parse multiplication of %q:\n%v\n", match, err)
		}
		sum += multiplied
	}
	return fmt.Sprintf("Result: %d", sum)
}

func part2(
	input string,
) string {
	finds := r.FindAllSubmatch([]byte(input), -1)
	sum := 0

	enabled := true
	for _, match := range finds {
		verb := string(match[1])
		switch verb {
		case "mul":
			if enabled {
				multiplied, err := mul(string(match[3]), string(match[4]))
				if err != nil {
					return fmt.Sprintf("Failed to parse multiplication of %q:\n%v\n", match, err)
				}
				sum += multiplied
			}
		case "do":
			enabled = true
		case "don't":
			enabled = false
		}
	}

	return fmt.Sprintf("Result: %d", sum)
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
