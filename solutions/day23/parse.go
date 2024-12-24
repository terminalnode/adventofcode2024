package main

import (
	"fmt"
	"gonum.org/v1/gonum/graph"
	"gonum.org/v1/gonum/graph/simple"
	"strings"
)

type connection struct {
	name    nodeName
	connMap connMap
}
type nodeName = string
type connMap map[nodeName]*connection

type textNode struct {
	id   int64
	name string
}

func (n textNode) ID() int64 {
	return n.id
}

func parseGonumGraph(
	input string,
) (*simple.UndirectedGraph, map[int64]string, error) {
	out := simple.NewUndirectedGraph()
	visited := make(map[string]graph.Node)
	intToName := make(map[int64]string)

	idCounter := int64(0)

	lines := strings.Split(input, "\n")
	for _, line := range lines {
		split := strings.Split(line, "-")
		if len(split) != 2 {
			return out, intToName, fmt.Errorf("length after splitting '%s' should be 2, was %d", line, len(split))
		}
		name1 := split[0]
		name2 := split[1]

		if visited[name1] == nil {
			idCounter++
			visited[name1] = simple.Node(idCounter)
			intToName[idCounter] = name1
			out.AddNode(visited[name1])
		}
		if visited[name2] == nil {
			idCounter++
			visited[name2] = simple.Node(idCounter)
			intToName[idCounter] = name2
			out.AddNode(visited[name2])
		}
		out.SetEdge(simple.Edge{F: visited[name1], T: visited[name2]})
	}

	return out, intToName, nil
}

func parse(
	input string,
) (connMap, error) {
	lines := strings.Split(input, "\n")
	out := make(connMap)

	for _, line := range lines {
		split := strings.Split(line, "-")
		if len(split) != 2 {
			return out, fmt.Errorf("length after splitting '%s' should be 2, was %d", line, len(split))
		}
		out.addConnection(split[0], split[1])
	}

	return out, nil
}

func (m connMap) addConnection(
	n1 nodeName,
	n2 nodeName,
) {
	node1 := m.getOrCreateNode(n1)
	node2 := m.getOrCreateNode(n2)
	node1.connMap[n2] = node2
	node2.connMap[n1] = node1
}

func (m connMap) getOrCreateNode(
	name nodeName,
) *connection {
	if n, ok := m[name]; ok {
		return n
	}
	m[name] = &connection{name: name, connMap: make(connMap)}
	return m[name]
}
