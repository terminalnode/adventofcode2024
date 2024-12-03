package main

import (
	"fmt"
	"github.com/terminalnode/adventofcode2024/common"
	"regexp"
	"strconv"
)

var findMulRegex = regexp.MustCompile("mul\\((\\d+),(\\d+)\\)")

func main() {
	common.Setup(2, part1, part2)
}

func part1(
	input string,
) string {
	finds := findMulRegex.FindAllSubmatch([]byte(input), -1)
	sum := 0
	for _, match := range finds {
		if len(match) != 3 {
			return fmt.Sprintf("Expected 3 results from regex, got %q\n", match)
		}

		multiplied, err := mul(string(match[1]), string(match[2]))
		if err != nil {
			return fmt.Sprintf("Failed to parse multiplication of %q:\n%v\n", match, err)
		}
		sum += multiplied
	}
	return fmt.Sprintf("Sum of all %d multiplications: %d", len(finds), sum)
}

func part2(
	input string,
) string {
	return "Not solved yet"
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
