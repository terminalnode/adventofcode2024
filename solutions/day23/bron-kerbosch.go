package main

import (
	"gonum.org/v1/gonum/graph/simple"
)

// I believe something like this is implemented in gonum already, seems
// like it anyway, but I've already cheated on the graph DS so might at
// least implement the algorithm myself.

func bronKerbosch(
	r []int64,
	p []int64,
	g *simple.UndirectedGraph,
) [][]int64 {
	// If P is empty, this is a maximal click
	if len(p) == 0 {
		return [][]int64{r}
	}
	out := make([][]int64, 0, len(p))
	x := make(map[int64]bool)

	existInP := make(map[int64]bool)
	for _, pEntry := range p {
		existInP[pEntry] = true
	}

	for _, v := range p {
		skipV := false

		fromV := g.From(v)
		newP := make([]int64, 0, fromV.Len())
		for fromV.Next() {
			id := fromV.Node().ID()
			if !existInP[id] {
				continue
			}

			if x[id] {
				// X set would not be empty, and thus can't return anything
				skipV = true
				break
			}
			newP = append(newP, id)
		}

		if !skipV {
			newR := append([]int64{v}, r...)
			out = append(out, bronKerbosch(newR, newP, g)...)
		}
		x[v] = true
	}

	return out
}
