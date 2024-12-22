package main

import (
	"fmt"
	"strconv"
	"strings"
)

func main() {
	// 245810 too high
	// 234958 too low
	fmt.Println(part1(exIn))
	fmt.Println(part1(realIn))
	// common.Setup(21, nil, nil)
}

func part1(
	input string,
) string {
	var out string

	sequences := strings.Split(input, "\n")
	sum := 0
	for _, rawSequence := range sequences {
		fmt.Println(rawSequence)
		numPart, err := numericPart(rawSequence)
		if err != nil {
			return "Fail: " + err.Error()
		}
		sequence := stringToCharArr(rawSequence)

		sequence = getMovements(kpNumeric, sequence)
		printIntArr("Numpad movements: ", sequence)
		sequence = getMovements(kpDirectional, sequence)
		printIntArr("First bot movements: ", sequence)
		sequence = getMovements(kpDirectional, sequence)
		printIntArr("Second bot movements: ", sequence)
		fmt.Println()
		l := len(sequence)

		out += fmt.Sprintf("%s: %d * %03d = %d\n", rawSequence, l, numPart, l*numPart)
		sum += l * numPart
	}
	out += fmt.Sprintf("Sum: %d", sum)
	return out
}

func getMovements(
	pad costMap,
	sequence []char,
) []char {
	state := 'A'
	kpPresses := make([][]char, len(sequence))
	for i, target := range sequence {
		newSeq := pad[state][target]
		kpPresses[i] = newSeq
		state = target
	}
	return flattenCharArr(kpPresses)
}

func numericPart(
	sequence string,
) (int, error) {
	nSeq := sequence[:len(sequence)-1]
	n, err := strconv.ParseInt(nSeq, 10, 0)
	return int(n), err
}

func stringToCharArr(
	s string,
) []char {
	out := make([]char, len(s))
	for i, ch := range s {
		out[i] = ch
	}
	return out
}

func flattenCharArr(
	arr [][]char,
) []char {
	size := 0
	for _, subarr := range arr {
		size += len(subarr)
	}

	out := make([]char, 0, size)
	for _, subarr := range arr {
		out = append(out, subarr...)
	}

	return out
}

func printIntArr(
	pre string,
	arr []char,
) {
	fmt.Printf(pre)
	for _, n := range arr {
		fmt.Printf("%c", n)
	}
	fmt.Println()
}
