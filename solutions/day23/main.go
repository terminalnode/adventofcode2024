package main

import (
	"fmt"
	"github.com/terminalnode/adventofcode2024/common"
	"github.com/terminalnode/adventofcode2024/common/util"
	"slices"
	"strings"
)

func main() {
	common.Setup(23, part1, part2)
}

func part1(
	input util.AocInput,
) (util.AocSolution, util.AocError) {
	cm, err := parse(input.Input)
	if err != nil {
		return util.NewAocError(err.Error(), util.InputParsingError)
	}

	visited := make(map[string]bool)
	for name1, conn1 := range cm {
		for name2, conn2 := range conn1.connMap {
			for name3, conn3 := range conn2.connMap {
				if conn3.connMap[name1] != nil {
					if name1[0] == 't' || name2[0] == 't' || name3[0] == 't' {
						names := []string{name1, name2, name3}
						slices.Sort(names)
						visited[strings.Join(names, "-")] = true
					}
				}
			}
		}
	}

	return util.FormatAocSolution("Clusters of 3 with computers starting with t: %d", len(visited))
}

func part2(
	input util.AocInput,
) (util.AocSolution, util.AocError) {
	cliques, err := getCliques(input.Input)
	if err != nil {
		msg := fmt.Sprintf("Failed to get cliques: %v", err)
		return util.NewAocError(msg, util.InputParsingError)
	}

	var biggest []string
	for _, c := range cliques {
		if len(c) > len(biggest) {
			biggest = c
		}
	}
	slices.Sort(biggest)

	return util.FormatAocSolution("Password: %s", strings.Join(biggest, ","))
}

func getCliques(
	input string,
) ([][]string, error) {
	graph, intToName, err := parseGonumGraph(input)
	if err != nil {
		return [][]string{}, err
	}

	nodes := graph.Nodes()
	p := make([]int64, 0, nodes.Len())
	for nodes.Next() {
		p = append(p, nodes.Node().ID())
	}

	cliques := bronKerbosch([]int64{}, p, graph)
	out := make([][]string, len(cliques))
	for i, clique := range cliques {
		strClique := make([]string, len(clique))
		for j, node := range clique {
			if name, ok := intToName[node]; ok {
				strClique[j] = name
			} else {
				return out, fmt.Errorf("could not find name for node %d", node)
			}
		}
		out[i] = strClique
	}

	return out, nil
}
