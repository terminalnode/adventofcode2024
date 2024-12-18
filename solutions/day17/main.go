package main

import (
	"fmt"
	"github.com/terminalnode/adventofcode2024/common"
)

func main() {
	common.Setup(17, part1, part2)
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

func part2(
	input string,
) string {
	m, err := parseMachine(input)
	if err != nil {
		return fmt.Sprintf("Failed to parse machine: %v", err)
	}

	lenSeq := len(m.seq)
	initA := int64(0)
	initB := m.b
	initC := m.c
	for initA = 0; true; initA++ {
		m.run(lenSeq)
		if compareSeqOut(m, lenSeq) {
			break
		}
		m.a = initA
		m.b = initB
		m.c = initC
		m.out = m.out[:0]
		m.pos = 0
	}

	return fmt.Sprintf("Registry A should be %d", initA-1)
}

func compareSeqOut(
	m machine,
	lenSeq int,
) bool {
	if len(m.out) != lenSeq {
		return false
	}

	for i, seq := range m.seq {
		if m.out[i] != seq {
			return false
		}
	}

	return true
}
