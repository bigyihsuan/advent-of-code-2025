package grid

import (
	"aoc2025/util/iters"
	"fmt"
	"slices"
	"strings"
)

type Offset struct{ x, y int }

func (o Offset) Apply(x, y int) (int, int) {
	return x + o.x, y + o.y
}

type Kernel []Offset

var (
	KernelAll           = Kernel{{-1, -1}, {-1, 0}, {-1, +1}, {0, +1}, {+1, +1}, {+1, 0}, {+1, -1}, {0, -1}}
	KernelAboveOnly     = Kernel{{-1, 0}}
	KernelBelowOnly     = Kernel{{+1, 0}}
	KernelDiagonalBelow = Kernel{{+1, -1}, {+1, +1}}
	KernelLeft          = Kernel{{0, -1}}
	KernelRight         = Kernel{{0, +1}}
	KernelSides         = Kernel{{0, -1}, {0, +1}}
)

type Grid[T comparable] [][]T

func NewFromStringGrid(g string) Grid[rune] {
	return Grid[rune](slices.Collect(
		iters.Map(
			strings.Lines(g),
			func(s string) []rune {
				return []rune(strings.TrimSpace(s))
			},
		),
	))
}

func (g Grid[T]) Height() int {
	return len(g)
}

func (g Grid[T]) Width() int {
	return len(g[0])
}

func (g Grid[T]) String() string {
	var sb strings.Builder
	for i := range g {
		for j := range g[i] {
			sb.WriteString(fmt.Sprint(g[i][j]))
		}
		sb.WriteString("\n")
	}
	return sb.String()
}

func (g Grid[T]) Clone() Grid[T] {
	var n Grid[T]
	for _, row := range g {
		r := []T{}
		r = append(r, row...)
		n = append(n, r)
	}
	return n
}

func (g Grid[T]) InBounds(x, y int) bool {
	return x >= 0 && x < g.Height() && y >= 0 && y < g.Width()
}

func (g Grid[T]) KernelOnCellHasElement(x, y int, kernel Kernel, e T) bool {
	for _, o := range kernel {
		i := x + o.x
		j := y + o.y
		if !g.InBounds(i, j) {
			continue
		}
		if g[i][j] == e {
			return true
		}
	}
	return false
}
func (g Grid[T]) KernelOnCellHasAnyElement(x, y int, kernel Kernel, es []T) bool {
	for _, o := range kernel {
		i := x + o.x
		j := y + o.y
		if !g.InBounds(i, j) {
			continue
		}
		if slices.Contains(es, g[i][j]) {
			return true
		}
	}
	return false
}

func (g Grid[T]) AllCellsInKernelHaveElement(x, y int, kernel Kernel, e T) bool {
	for _, o := range kernel {
		i := x + o.x
		j := y + o.y
		if !g.InBounds(i, j) {
			continue
		}
		if g[i][j] != e {
			return false
		}
	}
	return true
}

func (g Grid[T]) UpdateCellsWithKernel(x, y int, kernel Kernel, newElement T) (changed int) {
	for _, offset := range kernel {
		ox, oy := offset.Apply(x, y)
		if !g.InBounds(ox, oy) {
			continue
		}
		g[ox][oy] = newElement
		changed++
	}
	return
}
