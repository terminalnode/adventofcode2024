package main

import (
	"fmt"
	"github.com/terminalnode/adventofcode2024/common/util"
	"strconv"
	"strings"
)

func parse(
	input string,
) ([]util.Coordinate, error) {
	split := strings.Split(input, "\n")
	out := make([]util.Coordinate, len(split))
	for i, s := range split {
		xy := strings.Split(s, ",")
		if len(xy) != 2 {
			return out, fmt.Errorf("expected '%s' split by ',' to become a size 2 list", s)
		}

		x, err := strconv.ParseInt(xy[0], 10, 0)
		if err != nil {
			return out, err
		}

		y, err := strconv.ParseInt(xy[1], 10, 0)
		if err != nil {
			return out, err
		}

		out[i] = util.Coordinate{X: int(x), Y: int(y)}
	}

	return out, nil
}
