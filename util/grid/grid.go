package grid

import (
	"aoc2025/util/iters"
	"slices"
	"strings"
)

type Grid [][]rune

func NewFromStringGrid(g string) Grid {
	return Grid(slices.Collect(
		iters.Map(
			strings.Lines(g),
			func(s string) []rune {
				return []rune(strings.TrimSpace(s))
			},
		),
	))
}

func (g Grid) Height() int {
	return len(g)
}

func (g Grid) Width() int {
	return len(g[0])
}

func (g Grid) String() string {
	var sb strings.Builder
	for i := range g {
		for j := range g[i] {
			sb.WriteString(string(g[i][j]))
		}
		sb.WriteString("\n")
	}
	return sb.String()
}
