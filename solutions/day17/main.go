package main

import (
	"fmt"
	"github.com/terminalnode/adventofcode2024/common"
)

func main() {
	common.Setup(17, part1, nil)
}

func part1(
	input string,
) string {
	m, err := parseMachine(input)
	if err != nil {
		return fmt.Sprintf("Failed to parse machine: %v", err)
	}
	m.run(-1)
	return m.strOut()
}
