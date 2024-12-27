package main

import (
	"github.com/terminalnode/adventofcode2024/common/util"
	"testing"
)

func TestParseMachine(t *testing.T) {
	testInput := `Register A: -1337
Register B: 0
Register C: 666

Program: 0,-1,5,-4,3,0`
	testProgram := []int64{0, -1, 5, -4, 3, 0}

	m, err := parseMachine(testInput)

	if err != nil {
		t.Errorf("Failed to parse machine: %v", err)
	}

	if m.a != -1337 {
		t.Errorf("expected m.a == -1337, but was %d", m.c)
	}

	if m.b != 0 {
		t.Errorf("expected m.b == 0, but was %d", m.c)
	}

	if m.c != 666 {
		t.Errorf("expected m.c == 666, but was %d", m.c)
	}

	if len(m.seq) != len(testProgram) {
		t.Errorf("Expected len(m.seq) == %d, but was %d", len(testProgram), len(m.seq))
	}

	for i, e := range testProgram {
		a := m.seq[i]
		if e != a {
			t.Errorf("Expected m.seq[%d] == %d, but was %d", i, e, a)
		}
	}
}

func TestPart1_Example(t *testing.T) {
	const ex = `Register A: 729
Register B: 0
Register C: 0

Program: 0,1,5,4,3,0`
	const exOut = "4,6,3,5,6,3,5,2,1,0"

	out, _ := part1(util.AocInput{Input: ex})
	if out.Solution != exOut {
		t.Errorf("expected '%s' but got '%s'", exOut, out)
	}
}

func TestPart2_Example(t *testing.T) {
	ex := `Register A: 2024
Register B: 0
Register C: 0

Program: 0,3,5,4,3,0`
	exOut := "Registry A should be 117440"

	out, _ := part2(util.AocInput{Input: ex})
	if out.Solution != exOut {
		t.Errorf("expected '%s' but got '%s'", exOut, out)
	}
}
