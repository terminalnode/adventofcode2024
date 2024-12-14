package main

import (
	"errors"
	"fmt"
	"github.com/terminalnode/adventofcode2024/common/util"
	"regexp"
	"strconv"
)

type robot struct {
	init util.Coordinate
	move util.Coordinate
}

func (r robot) String() string {
	return fmt.Sprintf(
		"p=%d,%d v=%d,%d",
		r.init.X,
		r.init.Y,
		r.move.X,
		r.move.Y,
	)
}

var rx = regexp.MustCompile(`p=(-?\d+),(-?\d+) v=(-?\d+),(-?\d+)`)

func parseRobots(
	input string,
) ([]robot, error) {
	matches := rx.FindAllSubmatch([]byte(input), -1)
	robots := make([]robot, len(matches))

	if len(robots) == 0 {
		return robots, errors.New("input set contains no robots")
	}

	for i, match := range matches {
		if len(match) != 5 {
			return robots, fmt.Errorf("expected match to be 5 long, but was %d", len(match))
		}

		numbers := make([]int, 4)
		for j := 0; j < 4; j++ {
			n, err := strconv.ParseInt(string(match[j+1]), 10, 0)
			if err != nil {
				return robots, err
			}
			numbers[j] = int(n)
		}

		robots[i] = robot{
			init: util.Coordinate{X: numbers[0], Y: numbers[1]},
			move: util.Coordinate{X: numbers[2], Y: numbers[3]},
		}
	}

	return robots, nil
}
