package main

import (
	"fmt"
	"github.com/terminalnode/adventofcode2024/common"
	"log"
	"slices"
	"strconv"
	"strings"
)

func main() {
	common.Setup(24, part1, nil)
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

	for len(depSet) > 0 {
		newDepSet := make(map[name]bool)
		for dep := range depSet {
			depDeps, w := resolveDeps(r, wm, dep)
			for _, depDep := range depDeps {
				newDepSet[depDep] = true
			}

			if len(depDeps) == 1 {
				w.execute(r)
			}
		}
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
