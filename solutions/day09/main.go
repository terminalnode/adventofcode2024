package main

import (
	"fmt"
	"github.com/terminalnode/adventofcode2024/common"
)

func main() {
	common.Setup(9, part1, nil)
}

func part1(
	input string,
) string {
	sum := 0
	disk := parse(input)
	backIndex := len(disk)

	for i, id := range disk {
		if i >= backIndex {
			break
		}
		if id == -1 {
			backIndex--
			for disk[backIndex] == -1 {
				backIndex--
			}

			if i < backIndex {
				sum += i * disk[backIndex]
			}
		} else {
			sum += i * id
		}
	}

	return fmt.Sprintf("Sum: %d", sum)
}
