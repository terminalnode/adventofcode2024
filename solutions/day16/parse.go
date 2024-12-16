package main

import (
	"errors"
	"github.com/terminalnode/adventofcode2024/common/util"
	"strings"
)

type raceMap = [][]int32
type parsedInput struct {
	m raceMap
	s util.Coordinate
	e util.Coordinate
}

const (
	Start  = 'S'
	End    = 'E'
	Ground = '.'
	Wall   = '#'
)

func parse(
	input string,
) (parsedInput, error) {
	lines := strings.Split(input, "\n")
	m := make(raceMap, len(lines))
	s := util.Coordinate{}
	e := util.Coordinate{}

	for y, line := range lines {
		m[y] = []int32(line)
		if s.IsOrigin() || e.IsOrigin() {
			for x, ch := range line {
				if ch == Start {
					s = util.Coordinate{X: x, Y: y}
					m[y][x] = Ground
				} else if ch == End {
					e = util.Coordinate{X: x, Y: y}
					m[y][x] = Ground
				}
			}
		}
	}

	var err error
	switch {
	case s.IsOrigin() && e.IsOrigin():
		err = errors.New("neither start nor end was found")
	case s.IsOrigin():
		err = errors.New("start not found")
	case e.IsOrigin():
		err = errors.New("end not found")
	}
	return parsedInput{m: m, s: s, e: e}, err
}
