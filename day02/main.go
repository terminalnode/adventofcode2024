package main

import (
	"fmt"
	"github.com/terminalnode/adventofcode2024/common"
	"strconv"
	"strings"
)

type level = int

type report struct {
	levels []level
	isSafe bool
}

func main() {
	common.Setup(2, part1, part2)
}

func parseAllReports(input string) ([]report, error) {
	lines := strings.Split(input, "\n")
	reports := make([]report, len(lines))

	for i, line := range lines {
		levels, err := parseReportLevels(line)
		if err != nil {
			return nil, err
		}
		reports[i] = report{levels, isReportSafe(levels)}
	}

	return reports, nil
}

func parseReportLevels(input string) ([]level, error) {
	fields := strings.Fields(input)
	levels := make([]int, len(fields))

	for i, field := range fields {
		parsed, err := strconv.ParseInt(field, 10, 0)
		if err != nil {
			return nil, err
		}
		levels[i] = int(parsed)
	}

	return levels, nil
}

func isReportSafe(levels []level) bool {
	if len(levels) < 2 {
		return true
	}

	increasing := false
	if levels[0] < levels[1] {
		increasing = true
	}

	previous := levels[0]
	for _, current := range levels[1:] {
		var diff int
		if increasing {
			diff = current - previous
		} else {
			diff = previous - current
		}

		if diff < 1 || diff > 3 {
			return false
		}
		previous = current
	}

	return true
}

func part1(input string) string {
	reports, err := parseAllReports(input)
	if err != nil {
		return fmt.Sprintf("Failed to parse reports: %v", err)
	}

	count := 0
	for _, r := range reports {
		if r.isSafe {
			count += 1
		}
	}

	return fmt.Sprintf("Number of safe reports: %d", count)
}

func part2(input string) string {
	return fmt.Sprintf("Not implemented yet")
}
