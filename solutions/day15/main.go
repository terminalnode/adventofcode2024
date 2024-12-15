package main

import (
	"fmt"
	"github.com/terminalnode/adventofcode2024/common"
	"github.com/terminalnode/adventofcode2024/common/util"
)

func main() {
	common.Setup(15, part1, nil)
}

func part1(
	input string,
) string {
	p, err := parse(input)
	if err != nil {
		return fmt.Sprintf("Failed to parse input: %v", err)
	}

	for _, move := range p.moves {
		endPosition, err := findEndPosition(p.warehouse, p.robot, move)
		if err != nil {
			continue
		}

		newRobot := move(p.robot)
		if !newRobot.Equals(endPosition) {
			// End position is more than one step, meaning we need to move boxes
			p.warehouse[newRobot.Y][newRobot.X] = GROUND
			p.warehouse[endPosition.Y][endPosition.X] = BOX
		}
		p.robot = newRobot
	}

	return fmt.Sprintf("Sum of all GPS coordinates: %d", score(p.warehouse))
}

func findEndPosition(
	w warehouseMatrix,
	s util.Coordinate,
	d util.Direction,
) (util.Coordinate, error) {
	np := d(s)
	ch := w[np.Y][np.X]
	switch ch {
	case GROUND:
		return np, nil
	case BOX:
		return findEndPosition(w, np, d)
	case WALL:
		return np, fmt.Errorf("wall hit")
	}

	panic(fmt.Sprintf("Invalid character: %c", ch))
}

func score(
	wh warehouseMatrix,
) int {
	sum := 0
	for y, row := range wh {
		for x, ch := range row {
			if ch != BOX {
				continue
			}
			sum += x + 100*y
		}
	}
	return sum
}
