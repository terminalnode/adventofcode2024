package main

import (
	"fmt"
	"regexp"
	"strconv"
	"strings"
)

var r = regexp.MustCompile(`Register A: (-?\d+)\nRegister B: (-?\d+)\nRegister C: (-?\d+)\n\nProgram: ((?:-?\d,)+)`)

func parseMachine(
	input string,
) (machine, error) {
	out := machine{}
	matches := r.FindStringSubmatch(input)
	if len(matches) != 5 {
		return out, fmt.Errorf("expected four fields but got %d", len(matches))
	}

	digits, err := parseRegisters(matches)
	if err != nil {
		return out, err
	}
	out.a = digits[0]
	out.b = digits[1]
	out.c = digits[2]

	seq, err := parseProgram(matches[4])
	out.seq = seq

	return out, nil
}

func parseRegisters(
	matches []string,
) ([]int64, error) {
	out := make([]int64, 3)
	for i, m := range matches[1:4] {
		mAsI64, err := strconv.ParseInt(m, 10, 64)
		if err != nil {
			return out, err
		}
		out[i] = mAsI64
	}
	return out, nil
}

func parseProgram(
	program string,
) ([]int64, error) {
	parts := strings.Split(program, ",")
	out := make([]int64, len(parts))
	for i, part := range parts {
		partAsI64, err := strconv.ParseInt(part, 10, 64)
		if err != nil {
			return out, err
		}
		out[i] = partAsI64
	}
	return out, nil
}
