package util

import (
	"math"
	"slices"
)

func SliceWithoutIndexes[T any, S ~[]T](s S, idxs ...int) S {
	o := S{}
	for i, e := range s {
		if slices.Contains(idxs, i) {
			continue
		}
		o = append(o, e)
	}
	return o
}

func IntFromDigits(digits []int) int {
	n := 0
	for i, d := range digits {
		n += d * int(math.Pow10(len(digits)-i-1))
	}
	return n
}

func RotateGridCCW[T any](grid [][]T) [][]T {
	// https://stackoverflow.com/a/8664879/8143168
	for i, row := range grid {
		slices.Reverse(row)
		grid[i] = row
	}
	return Transpose(grid)
}

func Transpose[T any](grid [][]T) [][]T {
	// https://stackoverflow.com/a/26199060/8143168
	m, n := len(grid), len(grid[0])
	newGrid := make([][]T, n)
	for i := range newGrid {
		newGrid[i] = make([]T, m)
	}

	for i := range n {
		for j := range m {
			newGrid[i][j] = grid[j][i]
		}
	}

	return newGrid
}
