package main

import (
	"fmt"
	"github.com/terminalnode/adventofcode2024/common"
	"slices"
	"strconv"
	"strings"
)

func main() {
	common.Setup(1, part1, part2)
}

func createLists(input string) ([]int, []int, error) {
	lines := strings.Split(input, "\n")
	left := make([]int, len(lines))
	right := make([]int, len(lines))

	// Parse the input
	for i, line := range lines {
		parts := strings.Fields(line)
		if len(parts) != 2 {
			return nil, nil, fmt.Errorf("unable to extract parts from row #%d '%s'", i, line)
		}

		leftNum, leftErr := strconv.ParseInt(parts[0], 10, 32)
		if leftErr != nil {
			return nil, nil, fmt.Errorf("unable to parse left number %d on line %d", leftNum, i)
		}
		left[i] = int(leftNum)

		rightNum, rightErr := strconv.ParseInt(parts[1], 10, 32)
		if rightErr != nil {
			return nil, nil, fmt.Errorf("unable to parse right number %d on line %d", rightNum, i)
		}
		right[i] = int(rightNum)
	}

	return left, right, nil
}

func part1(input string) string {
	left, right, err := createLists(input)
	if err != nil {
		return err.Error()
	}

	// Sort the lists
	slices.Sort(left)
	slices.Sort(right)

	// Calculate result
	sum := 0
	for i, leftNum := range left {
		diff := leftNum - right[i]
		if diff < 0 {
			diff *= -1
		}
		sum += diff
	}

	return fmt.Sprintf("Result for part 1: %d", sum)
}

func part2(input string) string {
	left, right, err := createLists(input)
	if err != nil {
		return err.Error()
	}

	rightMap := make(map[int]int)
	for _, r := range right {
		rightMap[r] += 1
	}

	sum := 0
	for _, l := range left {
		sum += l * rightMap[l]
	}

	return fmt.Sprintf("Result for part 2: %d", sum)
}
