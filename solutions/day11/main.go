package main

import (
	"fmt"
	"github.com/terminalnode/adventofcode2024/common"
	"strconv"
	"strings"
)

// Stone => blinks remaining => output
type blinkCache = map[int]map[int]int

func main() {
	common.Setup(11, part1, part2)
}

func part1(input string) string { return solve(input, 25) }

func part2(input string) string { return solve(input, 75) }

func solve(
	input string,
	blinks int,
) string {
	stones, err := parseStones(input)
	cache := make(blinkCache)
	if err != nil {
		return fmt.Sprintf("Failed to parse stones: %v", err)
	}
	lenStart := len(stones)

	out := 0
	for i, stone := range stones {
		result, err := blink(stone, blinks, cache)
		if err != nil {
			return fmt.Sprintf("Failed to blink stone #%d (%d): %v", i+1, stone, err)
		}

		out += result
		fmt.Printf("Stone with value %d resulted in %d (out = %d)\n", stone, result, out)
	}

	return fmt.Sprintf("%d stones have turned into %d after %d blinks", lenStart, out, blinks)
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
	stone int,
	blinksLeft int,
	cache blinkCache,
) (int, error) {
	var out int
	var err error

	// Early exit if this value has already been computed.
	out = getCacheValueOrBaseCase(cache, stone, blinksLeft)
	if out != 0 {
		return out, err
	}

	// Apply the three rules
	if stone == 0 {
		out, err = blink(1, blinksLeft-1, cache)
	} else if s := strconv.Itoa(stone); len(s)%2 == 0 {
		out, err = blinkSplit(s, blinksLeft-1, cache)
	} else {
		out, err = blink(stone*2024, blinksLeft-1, cache)
	}

	// Exit on error
	if err != nil {
		return 0, err
	}

	// Enter value into the cache and return
	cache[stone][blinksLeft] = out
	return out, nil
}

func getCacheValueOrBaseCase(
	cache blinkCache,
	stone int,
	blinksLeft int,
) int {
	if blinksLeft == 0 {
		return 1
	}

	if cache[stone] == nil {
		cache[stone] = make(map[int]int)
	}

	return cache[stone][blinksLeft]
}

func blinkSplit(
	stone string,
	blinks int,
	cache blinkCache,
) (int, error) {
	left, right, err := splitNumber(stone)
	if err != nil {
		return 0, err
	}

	leftValue, err := blink(left, blinks, cache)
	if err != nil {
		return 0, err
	}

	rightValue, err := blink(right, blinks, cache)
	if err != nil {
		return 0, err
	}

	return leftValue + rightValue, nil
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
