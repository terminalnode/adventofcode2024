package main

import (
	"errors"
	"github.com/terminalnode/adventofcode2024/common/util"
	"strings"
)

type intX = int
type intY = int
type raceMap map[intY]map[intX]bool
type distanceMap map[intY]map[intX]int

type parsedInput struct {
	m      raceMap
	length int
	s      util.Coordinate
	e      util.Coordinate
}

func parse(
	input string,
) (parsedInput, error) {
	lines := strings.Split(input, "\n")
	m := make(raceMap)
	length := 0
	s := util.Coordinate{}
	e := util.Coordinate{}
	sFound := false
	eFound := false

	for y, line := range lines {
		m[y] = make(map[intX]bool)
		for x, ch := range line {
			if ch == '#' {
				continue
			}
			length += 1

			if ch == 'S' {
				sFound = true
				s = util.Coordinate{X: x, Y: y}
			} else if ch == 'E' {
				eFound = true
				e = util.Coordinate{X: x, Y: y}
			}
			m[y][x] = true
		}
	}
	out := parsedInput{m: m, s: s, e: e, length: length}

	if !sFound && !eFound {
		return out, errors.New("neither start nor end found")
	} else if !sFound {
		return out, errors.New("no start found")
	} else if !eFound {
		return out, errors.New("no end found")
	}

	return out, nil
}
