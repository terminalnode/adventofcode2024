package util

import (
	"fmt"
	"strings"
)

type Matrix[T any] struct {
	matrix [][]T
	MaxX   int
	MaxY   int
}

type IntMatrix = Matrix[int]
type Int8Matrix = Matrix[int8]
type Int16Matrix = Matrix[int16]
type Int32Matrix = Matrix[int32]
type Int64Matrix = Matrix[int64]

type UIntMatrix = Matrix[uint]
type UInt8Matrix = Matrix[uint8]
type UInt16Matrix = Matrix[uint16]
type UInt32Matrix = Matrix[uint32]
type UInt64Matrix = Matrix[uint64]

type CharMatrix = UInt8Matrix
type BoolMatrix = Matrix[bool]

func NewMatrixFromRows[T any](
	matrix [][]T,
) (Matrix[T], error) {
	maxX := len(matrix[0]) - 1
	maxY := len(matrix) - 1

	// Verify that the Matrix is consistent
	for i, row := range matrix[1:] {
		if len(row)-1 != maxX {
			err := fmt.Errorf("matrix[%d] length (%d) differs from matrix[0] (%d)", i+1, len(row), maxX+1)
			return Matrix[T]{}, err
		}
	}

	return Matrix[T]{
		matrix: matrix,
		MaxX:   maxX,
		MaxY:   maxY,
	}, nil
}

func NewCharMatrix(
	input string,
) (CharMatrix, error) {
	lines := strings.Split(input, "\n")
	rows := make([][]uint8, len(lines))

	for i, line := range lines {
		rows[i] = []uint8(line)
	}

	return NewMatrixFromRows[uint8](rows)
}

func (m Matrix[T]) IsInMatrix(
	x int,
	y int,
) bool {
	invalid := x < 0 || y < 0 || x > m.MaxX || y > m.MaxY
	return !invalid
}

func (m Matrix[T]) Get(
	x int,
	y int,
) (T, error) {
	if !m.IsInMatrix(x, y) {
		var zeroValue T
		return zeroValue, fmt.Errorf("invalid point (%d,%d)", x, y)
	}
	return m.matrix[y][x], nil
}

func (m Matrix[T]) GetOrDefault(
	x int,
	y int,
	defaultReturn T,
) (T, error) {
	v, err := m.Get(x, y)
	if err != nil {
		return defaultReturn, err
	}
	return v, nil
}

func (m Matrix[T]) Set(
	x int,
	y int,
	value T,
) error {
	if !m.IsInMatrix(x, y) {
		return fmt.Errorf("invalid point (%d,%d)", x, y)
	}

	m.matrix[y][x] = value
	return nil
}
