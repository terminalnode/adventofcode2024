package main

import (
	"errors"
	"fmt"
	"github.com/terminalnode/adventofcode2024/common/util"
	"regexp"
	"strconv"
)

type problem struct {
	a    util.Coordinate
	b    util.Coordinate
	goal util.Coordinate
}

var r = regexp.MustCompile(`Button A: X\+(\d+), Y\+(\d+)\nButton B: X\+(\d+), Y\+(\d+)\nPrize: X=(\d+), Y=(\d+)`)

func parseProblems(
	input string,
) ([]problem, error) {
	matches := r.FindAllSubmatch([]byte(input), -1)
	problems := make([]problem, len(matches))

	if len(problems) == 0 {
		return problems, errors.New("input set contains no problems")
	}

	for i, match := range matches {
		if len(match) < 7 {
			return problems, fmt.Errorf("expected match to be at least 7 long, but was %d", len(match))
		}

		digits := make([]int, 6)
		for j := 0; j < 6; j++ {
			d, err := strconv.ParseInt(string(match[j+1]), 10, 0)
			if err != nil {
				return problems, err
			}
			digits[j] = int(d)
		}

		problems[i] = problem{
			a:    util.Coordinate{X: digits[0], Y: digits[1]},
			b:    util.Coordinate{X: digits[2], Y: digits[3]},
			goal: util.Coordinate{X: digits[4], Y: digits[5]},
		}
	}

	return problems, nil
}
