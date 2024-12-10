package main

import (
	"fmt"
	"github.com/terminalnode/adventofcode2024/common"
)

type block struct {
	id   int
	size int
}

func (b *block) isFile() bool {
	return b.id >= 0
}

func main() {
	common.Setup(9, part1, nil)
}

func parseBlocks(
	input string,
) []block {
	isFile := true
	out := make([]block, len(input))

	for idx, c := range input {
		var id int
		if isFile {
			id = idx / 2
		} else {
			id = -1
		}

		out[idx] = block{
			id:   id,
			size: int(c - '0'),
		}

		isFile = !isFile
	}
	return out
}

func part1(
	input string,
) string {
	blocks := parseBlocks(input)
	sum := 0

	pos := 0
	for len(blocks) > 0 {
		b := blocks[0]
		if b.isFile() {
			i := 0
			for i < b.size {
			}
		}
	}

	return fmt.Sprintf("I don't know!")
}

func popLastBlock(
	blocks []block,
) (block, []block) {
	l := len(blocks)
	lastIdx := l - 1

	if l <= 1 {
		// We will panic with array index out of bounds if blocks is 0,
		// this is on purpose to make errors consistent with go built-ins
		return blocks[0], []block{}
	}

	lastBlock := blocks[lastIdx]
	newBlocks := blocks[:lastIdx]
	if lastBlock.isFile() {
		// Last block is a file, all is great
		return lastBlock, newBlocks
	}

	// Grab blocks until we get one that is a file
	trailingSpace := block{id: -1, size: lastBlock.size}
	l = len(newBlocks)
	for !lastBlock.isFile() && l > 0 {
		lastBlock = newBlocks[len(newBlocks)-1]
		l = len(newBlocks)
	}

	return lastBlock, newBlocks
}
