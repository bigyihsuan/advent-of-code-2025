//go:generate stringer -type=Space -linecomment
package day07

import (
	"aoc2025/util"
	"aoc2025/util/grid"
	"aoc2025/util/iters"
	"aoc2025/util/set"
	"fmt"
	"slices"
	"strings"
)

func Day(filename string, part int) {
	switch part {
	case 1:
		fmt.Println(Part1(filename))
	case 2:
		fmt.Println(Part2(filename))
	}
}

func Part1(filename string) int {
	manifold := Parse(filename)
	splits := 0
	beams := set.New(slices.Index(manifold.Grid[0], SOURCE)) // indexes of beams
	for _, row := range manifold.Grid[1:] {
		newBeams := set.New[int]()
		for beam := range beams {
			if row[beam] == SPLITTER {
				splits++
				newBeams.Add(beam-1, beam+1)
			} else {
				newBeams.Add(beam)
			}
		}
		beams = newBeams
	}
	return splits
}

func Part2(filename string) int {
	m := Parse(filename).Grid
	start := slices.Index(m[0], SOURCE)
	beams := map[int]int{start: 1}
	for _, row := range m[1:] {
		for idx, count := range beams {
			if row[idx] == SPLITTER {
				delete(beams, idx)
				beams[idx-1] = count + beams[idx-1]
				beams[idx+1] = count + beams[idx+1]
			}
		}
		fmt.Println(beams)
	}
	s := 0
	for _, c := range beams {
		s += c
	}
	return s
}

type Space int
type Space2 struct {
	Kind      Space
	BeamCount int
}

const (
	EMPTY    Space = iota // .
	SOURCE                // S
	SPLITTER              // ^
	BEAM                  // |
)

type Manifold struct{ grid.Grid[Space] }

func (m Manifold) String() string {
	o := strings.Join(slices.Collect(iters.Map(
		slices.Values(m.Grid),
		func(s []Space) string {
			return strings.Join(slices.Collect(iters.Map(
				slices.Values(s),
				Space.String,
			)), "")
		})), "\n")
	return o
}

func Parse(filename string) Manifold {
	parse := func(line string) (out []Space) {
		l := strings.TrimSpace(line)
		for _, r := range l {
			ele := EMPTY
			switch r {
			case '.':
				ele = EMPTY
			case 'S':
				ele = SOURCE
			case '^':
				ele = SPLITTER
			case '|':
				ele = BEAM
			}
			out = append(out, ele)
		}
		return
	}
	return Manifold{slices.Collect(util.LoadLines(filename, parse))}
}
func Parse2(filename string) grid.Grid[Space2] {
	parse := func(line string) (out []Space2) {
		l := strings.TrimSpace(line)
		for _, r := range l {
			ele := EMPTY
			switch r {
			case '.':
				ele = EMPTY
			case 'S':
				ele = SOURCE
			case '^':
				ele = SPLITTER
			case '|':
				ele = BEAM
			}
			out = append(out, Space2{ele, 0})
		}
		return
	}
	return slices.Collect(util.LoadLines(filename, parse))
}
