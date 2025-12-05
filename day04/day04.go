package day04

import (
	"aoc2025/util"
	"aoc2025/util/grid"
	"fmt"
	"slices"
)

const (
	ROLL  = '@'
	EMPTY = '.'
)

func Day(filename string, part int) {
	sss := util.LoadStringGrid(filename)
	switch part {
	case 1:
		fmt.Println(Part01(sss))
	case 2:
		fmt.Println(Part02(sss))
	}
}

var offsets = []struct{ x, y int }{{-1, -1}, {-1, 0}, {-1, 1}, {0, 1}, {1, 1}, {1, 0}, {1, -1}, {0, -1}}

func Part01(g grid.Grid) int {
	rolls := 0
	for i := range g.Height() {
		for j := range g.Width() {
			if g[i][j] == ROLL {
				adjacentRolls := 0
				for _, o := range offsets {
					x, y := i+o.x, j+o.y
					if x < 0 || x >= g.Height() || y < 0 || y >= g.Width() {
						continue
					} else if g[x][y] == ROLL {
						adjacentRolls++
					}
				}
				if adjacentRolls < 4 {
					rolls++
				}
			}
		}
	}
	return rolls
}

func Part02(g grid.Grid) int {
	rolls := 0
	curr := g.Clone()
	next := g.Clone()
	for {
		fmt.Println(curr)
		for i := range curr.Height() {
			for j := range curr.Width() {
				if curr[i][j] == ROLL {
					adjacentRolls := 0
					for _, o := range offsets {
						x, y := i+o.x, j+o.y
						if x < 0 || x >= curr.Height() || y < 0 || y >= curr.Width() {
							continue
						} else if curr[x][y] == ROLL {
							adjacentRolls++
						}
					}
					if adjacentRolls < 4 {
						rolls++
						next[i][j] = EMPTY
					}
				}
			}
		}
		if slices.EqualFunc(curr, next, slices.Equal) {
			break
		} else {
			curr = next.Clone()
		}
	}
	return rolls
}
