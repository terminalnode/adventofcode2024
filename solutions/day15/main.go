package main

import (
	"errors"
	"github.com/terminalnode/adventofcode2024/common"
	"github.com/terminalnode/adventofcode2024/common/util"
)

func main()                                                       { common.Setup(15, part1, part2) }
func part1(input util.AocInput) (util.AocSolution, util.AocError) { return solve(input, false) }
func part2(input util.AocInput) (util.AocSolution, util.AocError) { return solve(input, true) }

type visitedSet = map[int]map[int]bool

func solve(
	input util.AocInput,
	makeWide bool,
) (util.AocSolution, util.AocError) {
	p, err := parse(input.Input, makeWide)
	if err != nil {
		return util.NewAocError(err.Error(), util.InputParsingError)
	}

	for _, move := range p.moves {
		newPos := move(p.robot)
		boxMoves, err := getBoxMoves(p.warehouse, newPos, move, make(visitedSet))
		if err != nil {
			continue
		}

		// Zero out all moves
		for _, bm := range boxMoves {
			c := bm.curr
			p.warehouse[c.Y][c.X] = Ground
		}

		// Put all boxes in their new positions
		for _, bm := range boxMoves {
			c := bm.new
			p.warehouse[c.Y][c.X] = bm.ch
		}

		p.robot = newPos
	}

	if makeWide {
		return util.FormatAocSolution("Sum of all GPS coordinates in the wide area: %d", score(p.warehouse))
	}
	return util.FormatAocSolution("Sum of all GPS coordinates: %d", score(p.warehouse))
}

type boxMove struct {
	ch   int32
	curr util.Coordinate
	new  util.Coordinate
}

func getBoxMoves(
	w warehouseMatrix,
	start util.Coordinate,
	d util.Direction,
	set visitedSet,
) ([]boxMove, error) {
	if inVisitedSet(set, start) {
		return []boxMove{}, nil
	}
	set[start.X][start.Y] = true

	// Initializations, boxMoves capacity is just a guess
	boxMoves := make([]boxMove, 0, 200)
	ch := w[start.Y][start.X]

	// Early returns on ground and wall
	if ch == Ground {
		return boxMoves, nil
	} else if ch == Wall {
		return boxMoves, errors.New("wall hit")
	}

	// Whatever happens, add this move to the list
	newPos := d(start)
	boxMoves = append(boxMoves, boxMove{ch: ch, curr: start, new: newPos})

	// Get next moves and add to the list
	next, err := getBoxMoves(w, newPos, d, set)
	if err != nil {
		return boxMoves, err
	}
	boxMoves = append(boxMoves, next...)

	// Now do the same thing for the partner boxes, if any
	if ch != Box {
		if ch == LeftBox {
			next, err = getBoxMoves(w, start.East(), d, set)
		} else if ch == RightBox {
			next, err = getBoxMoves(w, start.West(), d, set)
		}
		if err != nil {
			return boxMoves, err
		}
		boxMoves = append(boxMoves, next...)
	}

	return boxMoves, err
}

func inVisitedSet(
	set visitedSet,
	c util.Coordinate,
) bool {
	if set[c.X] == nil {
		set[c.X] = make(map[int]bool)
	}
	return set[c.X][c.Y]
}

func score(
	wh warehouseMatrix,
) int {
	sum := 0
	for y, row := range wh {
		for x, ch := range row {
			if ch != Box && ch != LeftBox {
				continue
			}
			sum += x + 100*y
		}
	}
	return sum
}
