package main

import (
	"errors"
	"fmt"
	"regexp"
	"strconv"
)

type problem struct {
	aX    int64
	aY    int64
	bX    int64
	bY    int64
	goalX int64
	goalY int64
}

var r = regexp.MustCompile(`Button A: X\+(\d+), Y\+(\d+)\nButton B: X\+(\d+), Y\+(\d+)\nPrize: X=(\d+), Y=(\d+)`)

func parseProblems(
	input string,
	part2 bool,
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

		digits := make([]int64, 6)
		for j := 0; j < 6; j++ {
			d, err := strconv.ParseInt(string(match[j+1]), 10, 64)
			if err != nil {
				return problems, err
			}
			digits[j] = d
		}

		goalX := digits[4]
		goalY := digits[5]
		if part2 {
			goalX += 10000000000000
			goalY += 10000000000000
		}

		problems[i] = problem{
			aX:    digits[0],
			aY:    digits[1],
			bX:    digits[2],
			bY:    digits[3],
			goalX: goalX,
			goalY: goalY,
		}
	}

	return problems, nil
}
