package day04

import (
	"aoc2025/util"
	"aoc2025/util/grid"
	"fmt"
)

func Day04(filename string, part int) {
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
			if g[i][j] == '@' {
				adjacentRolls := 0
				for _, o := range offsets {
					x, y := i+o.x, j+o.y
					if x < 0 || x >= g.Height() || y < 0 || y >= g.Width() {
						continue
					} else if g[x][y] == '@' {
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
	return -1
}
