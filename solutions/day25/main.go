package main

import (
	"fmt"
	"github.com/terminalnode/adventofcode2024/common"
)

func main() {
	common.Setup(25, part1, nil)
}

func part1(
	input string,
) string {
	keys, locks := parse(input)
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

	return fmt.Sprintf("Number of ok key-lock combos: %d", count)
}
