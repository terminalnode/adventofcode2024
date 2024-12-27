package main

import (
	"github.com/terminalnode/adventofcode2024/common"
	"github.com/terminalnode/adventofcode2024/common/util"
)

func main() {
	common.Setup(25, part1, part2)
}

func part1(
	input util.AocInput,
) (util.AocSolution, util.AocError) {
	keys, locks := parse(input.Input)
	count := 0
	for _, key := range keys {
		for _, lock := range locks {
			allOk := true
			for i, keyN := range key.a {
				lockN := lock.a[i]
				if keyN+lockN > 7 {
					allOk = false
					break
				}
			}
			if allOk {
				count++
			}
		}
	}

	return util.FormatAocSolution("Number of ok key-lock combos: %d", count)
}

func part2(
	_ util.AocInput,
) (util.AocSolution, util.AocError) {
	return util.NewAocSolution("Maybe the true day 25 part 2 were the friends we made along the way?")
}
