package main

import (
	"github.com/terminalnode/adventofcode2024/common"
	"github.com/terminalnode/adventofcode2024/common/util"
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

func parseAllReports(
	input string,
	problemDampener bool,
) ([]report, error) {
	lines := strings.Split(input, "\n")
	reports := make([]report, len(lines))

	for i, line := range lines {
		levels, err := parseReportLevels(line)
		if err != nil {
			return nil, err
		}

		isSafe := isReportSafe(levels, -1)
		if !isSafe && problemDampener {
			for skipIdx := range levels {
				isSafe = isReportSafe(levels, skipIdx)
				if isSafe {
					break
				}
			}
		}

		reports[i] = report{levels, isSafe}
	}

	return reports, nil
}

func parseReportLevels(
	input string,
) ([]level, error) {
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

func isReportSafe(
	levels []level,
	skipIndex int,
) bool {
	var increasing bool
	switch skipIndex {
	case 0:
		increasing = levels[1] < levels[2]
	case 1:
		increasing = levels[0] < levels[2]
	default:
		increasing = levels[0] < levels[1]
	}

	var firstIdx int
	if skipIndex == 0 {
		firstIdx = 1
	} else {
		firstIdx = 0
	}
	previous := levels[firstIdx]

	for idx, current := range levels[1:] {
		if skipIndex == idx+1 || firstIdx == idx+1 {
			continue
		}

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

func countSafe(
	reports []report,
) int {
	count := 0
	for _, r := range reports {
		if r.isSafe {
			count += 1
		}
	}
	return count
}

func part1(
	input util.AocInput,
) (util.AocSolution, util.AocError) {
	reports, err := parseAllReports(input.Input, false)
	if err != nil {
		return util.NewAocError(err.Error(), util.InputParsingError)
	}

	return util.FormatAocSolution("Number of safe reports: %d", countSafe(reports))
}

func part2(
	input util.AocInput,
) (util.AocSolution, util.AocError) {
	reports, err := parseAllReports(input.Input, true)
	if err != nil {
		return util.NewAocError(err.Error(), util.InputParsingError)
	}

	return util.FormatAocSolution("Number of safe reports: %d", countSafe(reports))
}
