package main

import (
	"fmt"
	"github.com/terminalnode/adventofcode2024/common"
	"strconv"
	"strings"
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

	maxIdx := len(m.seq) - 1
	octArr := make([]int, len(m.seq))
	octIdx := 0
	initB := m.b
	initC := m.c

	for {
		seqIdx := maxIdx - octIdx

		// Rig the machine
		octArr[octIdx] += 1
		m.b = initB
		m.c = initC
		m.out = m.out[:0]

		m.a, err = arrayToOct(octArr)
		if err != nil {
			return fmt.Sprintf("Failed to read %v as octal string: %v", octArr, err)
		}

		// Run the program and verify output
		m.run(maxIdx + 1)
		correct := m.out[seqIdx] == m.seq[seqIdx]

		if octIdx == maxIdx && correct {
			break
		} else if correct {
			octIdx++
			octArr[octIdx] = -1
		}

		for octArr[octIdx] == 7 {
			octArr[octIdx] = 0
			octIdx--
		}
	}

	final, err := arrayToOct(octArr)
	if err != nil {
		return fmt.Sprintf("Solved it, but failed to extract number: %v", err)
	}
	return fmt.Sprintf("Registry A should be %d", final)
}

func arrayToOct(
	arr []int,
) (int64, error) {
	strArr := make([]string, len(arr))
	for i, n := range arr {
		strArr[i] = strconv.Itoa(n)
	}

	oct, err := strconv.ParseInt(strings.Join(strArr, ""), 8, 64)
	if err != nil {
		return 0, err
	}
	return oct, nil
}
