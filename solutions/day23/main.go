package main

import (
	"fmt"
	"github.com/terminalnode/adventofcode2024/common"
	"slices"
	"strings"
)

func main() {
	common.Setup(23, part1, part2)
}

func part1(
	input string,
) string {
	cm, err := parse(input)
	if err != nil {
		return fmt.Sprintf("Failed to parse input: %v", err)
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

	return fmt.Sprintf("Clusters of 3 with computers starting with t: %d", len(visited))
}

func part2(
	input string,
) string {
	cliques, err := getCliques(input)
	if err != nil {
		return fmt.Sprintf("Failed to get cliques: %v", err)
	}

	var biggest []string
	for _, c := range cliques {
		if len(c) > len(biggest) {
			biggest = c
		}
	}
	slices.Sort(biggest)

	return fmt.Sprintf("Password: %s", strings.Join(biggest, ","))
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
