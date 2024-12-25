package main

import (
	"fmt"
	"github.com/terminalnode/adventofcode2024/common"
	"log"
	"math/bits"
	"slices"
	"strconv"
	"strings"
)

func main() {
	common.Setup(24, part1, part2)
}

func part1(
	input string,
) string {
	r, wm, err := parse(input)
	if err != nil {
		return fmt.Sprintf("Failed to parse input: %v", err)
	}
	zNames := findIndexes('z', r, wm)
	resolveNames(r, wm, zNames)
	out, strOut, err := toInt(r, zNames)
	if err != nil {
		return fmt.Sprintf("Failed to read %s as binary: %v", strOut, err)
	}

	return fmt.Sprintf("Decimal output is %d (binary %s)", out, strOut)
}

func part2(
	input string,
) string {
	r, wm, err := parse(input)
	if err != nil {
		return fmt.Sprintf("Failed to parse input: %v", err)
	}

	zWrong := make([]name, 0)
	carryWrong := make([]name, 0)

	// Rules deduced by people who know how a ripple adder works and who have figured out the structure of the
	// input. Most people just plopped it into a graphing tool or solved it with pen and paper it seems.
	// I hated every second of this problem with a passion.
	for out, w := range wm {
		if out[0] == 'z' && out != "z45" && w.op != XOR {
			zWrong = append(zWrong, out)
		}

		inputsNotXY := out[0] != 'z' &&
			w.p1[0] != 'x' && w.p1[0] != 'y' &&
			w.p2[0] != 'x' && w.p2[0] != 'y'
		if inputsNotXY && out[0] != 'z' && w.op == XOR {
			carryWrong = append(carryWrong, out)
		}
	}

	pairs := make([][2]name, 0, 3)
	for _, carry := range carryWrong {
		zOutput := findFirstZ(carry, wm)
		zNum, _ := strconv.Atoi(string(zOutput[1:]))
		pairs = append(pairs, [2]name{carry, name(fmt.Sprintf("z%02d", zNum-1))})
	}

	xNames := findIndexes('x', r, wm)
	yNames := findIndexes('y', r, wm)
	zNames := findIndexes('z', r, wm)
	resolveNames(r, wm, zNames, yNames, xNames)

	xOut, _, _ := toInt(r, xNames)
	yOut, _, _ := toInt(r, yNames)
	zOut, _, _ := toInt(r, zNames)

	expected := xOut + yOut
	wrongBits := zOut ^ expected
	falseCarry := bits.TrailingZeros64(uint64(wrongBits))

	lastPair := make([]name, 0, 2)
	for out, w := range wm {
		suffix := fmt.Sprintf("%02d", falseCarry)
		if strings.HasSuffix(w.p1, suffix) &&
			strings.HasSuffix(w.p2, suffix) {
			lastPair = append(lastPair, out)
		}
	}

	final := make([]name, 0, 6)
	final = append(final, zWrong...)
	final = append(final, carryWrong...)
	final = append(final, lastPair...)
	slices.Sort(final)

	return strings.Join(final, ",")
}

func findFirstZ(
	start name,
	wm wireMap,
) name {
	visited := make(map[name]bool)
	current := start

	for {
		if current[0] == 'z' {
			return current
		}
		visited[current] = true

		for _, w := range wm {
			if (w.p1 == current || w.p2 == current) && !visited[w.out] {
				current = w.out
				break
			}
		}
	}
}

func toInt(
	r registry,
	names []name,
) (int, string, error) {
	values := make([]string, len(names))
	for i, n := range names {
		if r[n] {
			values[i] = "1"
		} else {
			values[i] = "0"
		}
	}

	slices.Reverse(values)
	outStr := strings.Join(values, "")
	out64, err := strconv.ParseInt(outStr, 2, 0)
	outInt := int(out64)
	if err != nil {
		return outInt, outStr, fmt.Errorf("failed to read %s as binary: %v", outStr, err)
	}

	return outInt, outStr, nil
}

func resolveNames(
	r registry,
	wm wireMap,
	nameSets ...[]name,
) {
	depSet := make(map[name]bool)
	for _, nameSet := range nameSets {
		for _, n := range nameSet {
			depSet[n] = true
		}
	}

	anyNew := true
	for len(depSet) > 0 && anyNew {
		anyNew = false
		newDepSet := make(map[name]bool)
		for dep := range depSet {
			depDeps, w := resolveDeps(r, wm, dep)
			for _, depDep := range depDeps {
				anyNew = anyNew || !depSet[depDep]
				newDepSet[depDep] = true
			}

			if len(depDeps) == 1 {
				w.execute(r)
			}
		}

		anyNew = anyNew || len(depSet) != len(newDepSet)
		depSet = newDepSet
	}
}

func resolveDeps(
	r registry,
	wm wireMap,
	target name,
) ([]name, wire) {
	out := make([]name, 0, 3)
	if _, ok := r[target]; ok {
		return out, wire{}
	}
	out = append(out, target)

	w, ok := wm[target]
	if !ok {
		log.Println("WARNING: target unresolvable", target)
		return out, w
	}

	if _, ok := r[w.p1]; !ok {
		out = append(out, w.p1)
	}
	if _, ok := r[w.p2]; !ok {
		out = append(out, w.p2)
	}

	return out, w
}

func (w wire) execute(
	r registry,
) {
	p1 := r[w.p1]
	p2 := r[w.p2]
	switch w.op {
	case AND:
		r[w.out] = p1 && p2
	case OR:
		r[w.out] = p1 || p2
	case XOR:
		r[w.out] = (p1 || p2) && (p1 != p2)
	}
}

func findIndexes(
	prefix rune,
	r registry,
	wm wireMap,
) []name {
	out := make([]name, 0, 30)
	for curr := 0; ; curr++ {
		key := name(fmt.Sprintf("%c%02d", prefix, curr))
		_, inR := r[key]
		_, inWM := wm[key]
		if !inR && !inWM {
			break
		}
		out = append(out, key)
	}
	slices.Sort(out)
	return out
}
