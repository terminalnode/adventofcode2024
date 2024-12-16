package main

import (
	"fmt"
	"github.com/terminalnode/adventofcode2024/common/util"
)

type intX = int
type intY = int
type intDirection = int
type intScore = int
type visitedSet = map[intY]map[intX]intScore
type winningSet = map[intY]map[intX]bool

const (
	North = iota
	East
	South
	West
)

type reindeer struct {
	p         util.Coordinate
	score     int
	direction int
}

func initialReindeer(
	start util.Coordinate,
) reindeer {
	return reindeer{
		p:         start,
		score:     0,
		direction: East,
	}
}

func (r reindeer) forward(m raceMap) (reindeer, error) {
	var newP util.Coordinate
	switch r.direction {
	case North:
		newP = r.p.North()
	case East:
		newP = r.p.East()
	case South:
		newP = r.p.South()
	case West:
		newP = r.p.West()
	default:
		return r, fmt.Errorf("illegal direction %d", r.direction)
	}

	if m[r.p.Y][r.p.X] == '#' {
		return r, fmt.Errorf("wall hit")
	}

	return reindeer{
		p:         newP,
		score:     1 + r.score,
		direction: r.direction,
	}, nil
}

func (r reindeer) turnClockwise() reindeer {
	return reindeer{
		p:         r.p,
		score:     1000 + r.score,
		direction: (r.direction + 1) % 4,
	}
}

func (r reindeer) turnCounterClockwise() reindeer {
	return reindeer{
		p:         r.p,
		score:     1000 + r.score,
		direction: (r.direction - 1 + 4) % 4,
	}
}

func (r reindeer) visitAndCheckIfDead(
	set visitedSet,
	leeway bool,
) bool {
	v := set[r.p.Y][r.p.X]
	if v == 0 || r.score < v {
		set[r.p.Y][r.p.X] = r.score
		return false
	}

	// This leeway allows to find "temporarily suboptimal paths" when doing part 2
	if leeway && r.score < (v+2000) {
		return false
	}

	return true
}

func (r reindeer) String() string {
	return fmt.Sprintf("Reindeer{p:%s, score:%d, direction:%d}", r.p, r.score, r.direction)
}
