package main

import (
	"fmt"
	"github.com/terminalnode/adventofcode2024/common"
	"strconv"
	"strings"
)

func main() {
	common.Setup(11, part1, nil)
}

func part1(
	input string,
) string {
	stones, err := parseStones(input)
	lenStart := len(stones)

	if err != nil {
		return fmt.Sprintf("Failed to parse stones: %v", err)
	}

	for i := 0; i < 25; i++ {
		stones, err = blink(stones)
		if err != nil {
			return fmt.Sprintf("Run failed on blink #%d: %v", i+1, err)
		}
	}

	return fmt.Sprintf("%d stones have turned into %d after 25 blinks", lenStart, len(stones))
}

func parseStones(
	input string,
) ([]int, error) {
	numbers := strings.Split(input, " ")

	out := make([]int, 0, len(numbers))

	for _, sNumber := range numbers {
		digit, err := strconv.ParseInt(sNumber, 10, 0)
		if err != nil {
			return out, err
		}
		out = append(out, int(digit))
	}

	return out, nil
}

func blink(
	stones []int,
) ([]int, error) {
	// Initialize the out-array with twice the size to ensure that we never need to resize it
	out := make([]int, 0, 2*len(stones))

	for _, stone := range stones {
		if stone == 0 {
			out = append(out, 1)
		} else if s := strconv.Itoa(stone); len(s)%2 == 0 {
			left, right, err := splitNumber(s)
			if err != nil {
				return out, err
			}
			out = append(out, left, right)
		} else {
			out = append(out, stone*2024)
		}
	}

	return out, nil
}

func splitNumber(
	s string,
) (int, int, error) {
	l := len(s)
	left, err := strconv.ParseInt(s[:l/2], 10, 0)
	if err != nil {
		return 0, 0, err
	}

	right, err := strconv.ParseInt(s[l/2:], 10, 0)
	if err != nil {
		return 0, 0, err
	}

	return int(left), int(right), nil
}
