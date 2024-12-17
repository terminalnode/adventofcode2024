package main

import (
	"fmt"
	"math"
	"testing"
)

func TestMachineRun(t *testing.T) {
	tests := []struct {
		n string
		m machine
		f func(machine) bool
	}{
		{
			n: "#1 If register C contains 9, the program 2,6 would set register B to 1.",
			m: machine{c: 9, seq: []int64{2, 6}},
			f: func(m machine) bool { return m.b == 1 },
		}, {
			n: "#2 If register A contains 10, the program 5,0,5,1,5,4 would output 0,1,2.",
			m: machine{a: 10, seq: []int64{5, 0, 5, 1, 5, 4}},
			f: func(m machine) bool { return m.strOut() == "0,1,2" },
		}, {
			n: "#3 If register A contains 2024, the program 0,1,5,4,3,0 would output 4,2,5,6,7,7,7,7,3,1,0 and leave 0 in register A.",
			m: machine{a: 2024, seq: []int64{0, 1, 5, 4, 3, 0}},
			f: func(m machine) bool {
				return m.a == 0 && m.strOut() == "4,2,5,6,7,7,7,7,3,1,0"
			},
		}, {
			n: "#4 If register B contains 29, the program 1,7 would set register B to 26.",
			m: machine{b: 29, seq: []int64{1, 7}},
			f: func(m machine) bool { return m.b == 26 },
		}, {
			n: "#5 If register B contains 2024 and register C contains 43690, the program 4,0 would set register B to 44354.",
			m: machine{b: 2024, c: 43690, seq: []int64{4, 0}},
			f: func(m machine) bool { return m.b == 44354 },
		},
	}

	for _, test := range tests {
		t.Run(test.n, func(t *testing.T) {
			test.m.run(100)
			if !test.f(test.m) {
				t.Errorf("Test validation failed with %s", test.m.String())
				fmt.Println(test.m)
			}
		})
	}
}

func TestMachine_String(t *testing.T) {
	min64I := int64(math.MinInt64) // -9223372036854775808
	max64I := int64(math.MaxInt64) // 9223372036854775807

	tests := []struct {
		n string
		m machine
		e string
	}{
		{
			n: "machine with boundary int64 values",
			m: machine{a: min64I, b: 0, c: max64I, seq: []int64{min64I, 0, max64I}, pos: 2},
			e: "machine{a:-9223372036854775808, b:0, c:9223372036854775807, seq:[-9223372036854775808 0 9223372036854775807], out:[], pos:2 (9223372036854775807)}",
		},
		{
			n: "machine with positive out of bounds index",
			m: machine{a: 0, b: 1, c: 2, seq: []int64{1, 2, 3, 4, 5}, out: []int64{1, 2, 3}, pos: 5},
			e: "machine{a:0, b:1, c:2, seq:[1 2 3 4 5], out:[1 2 3], pos:5 (out of bounds)}",
		},
		{
			n: "machine with negative out of bounds index",
			m: machine{a: 0, b: 1, c: 2, seq: []int64{1, 2, 3, 4, 5}, pos: -1},
			e: "machine{a:0, b:1, c:2, seq:[1 2 3 4 5], out:[], pos:-1 (out of bounds)}",
		},
	}

	for _, test := range tests {
		t.Run(test.n, func(t *testing.T) {
			actual := test.m.String()
			if actual != test.e {
				t.Errorf("expected: '%s'\nbut got: '%s'", test.e, actual)
			}
		})
	}
}
