package main

import (
	"fmt"
	"strings"
)

type name = string
type registry = map[name]bool
type wireMap = map[name]wire

type op int

const (
	AND = op(iota)
	OR
	XOR
)

type wire struct {
	p1  name
	op  op
	p2  name
	out name
}

func parse(
	input string,
) (registry, wireMap, error) {
	rawRegistry, rawWires, err := splitInput(input)
	outRegistry := make(registry)
	outWires := make(wireMap)
	if err != nil {
		return outRegistry, outWires, err
	}

	for _, raw := range rawRegistry {
		n, v, err := parseReg(raw)
		if err != nil {
			return outRegistry, outWires, err
		}
		outRegistry[n] = v
	}

	for _, raw := range rawWires {
		w, err := parseWire(raw)
		if err != nil {
			return outRegistry, outWires, err
		}
		outWires[w.out] = w
	}

	return outRegistry, outWires, nil
}

func splitInput(
	input string,
) ([]string, []string, error) {
	split := strings.Split(input, "\n\n")
	if len(split) != 2 {
		return []string{}, []string{}, fmt.Errorf("failed to split input, got %d parts", len(split))
	}

	return strings.Split(split[0], "\n"), strings.Split(split[1], "\n"), nil
}

func parseReg(
	raw string,
) (name, bool, error) {
	split := strings.Split(raw, ": ")
	if len(split) != 2 {
		return "", false, fmt.Errorf("failed to split raw registry %s, got %d parts (%v)", raw, len(split), split)
	} else if split[1] != "1" && split[1] != "0" {
		return "", false, fmt.Errorf("failed to parse registry %s with initial value %s", raw, split[1])
	}

	return split[0], split[1] == "1", nil
}

func parseWire(
	raw string,
) (wire, error) {
	split := strings.Split(raw, " ")
	if len(split) != 5 {
		return wire{}, fmt.Errorf("failed to split raw wire %s, got %d parts (%v)", raw, len(split), split)
	}

	p1 := split[0]
	rawOp := split[1]
	p2 := split[2]
	out := split[4]

	var realOp op
	switch rawOp {
	case "AND":
		realOp = AND
	case "OR":
		realOp = OR
	case "XOR":
		realOp = XOR
	default:
		return wire{}, fmt.Errorf("failed to parse operation %s", rawOp)
	}

	return wire{p1: p1, op: realOp, p2: p2, out: out}, nil
}
