package day08

import (
	"aoc2025/util"
	"aoc2025/util/iters"
	"aoc2025/util/strpipe"
	"cmp"
	"fmt"
	"iter"
	"maps"
	"slices"
	"sort"
	"strconv"

	"github.com/ihebu/dsu"
	"github.com/kr/pretty"
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
	boxes := slices.Collect(util.LoadLines(filename, ParseBox))

	edges := MakeEdges(boxes)
	fmt.Println(edges.Len())
	forest := Kruskal(edges, len(boxes))
	counts := make(map[Box]int)
	for _, b := range boxes {
		boxSet := forest.Find(b)
		if boxSet == nil {
			continue // this box is not connected to anything
		}
		box := boxSet.(Box)
		if _, ok := counts[box]; !ok {
			counts[box] = 1
		} else {
			counts[box]++
		}
	}
	pretty.Println(counts)

	cts := slices.SortedFunc(maps.Values(counts), func(a, b int) int { return cmp.Compare(b, a) })
	fmt.Println(cts)
	firstThree := cts[:3]
	fmt.Println(firstThree)
	return iters.Product(slices.Values(firstThree))
}

func Part2(filename string) int {
	boxes := slices.Collect(util.LoadLines(filename, ParseBox))

	edges := MakeEdges(boxes)
	// fmt.Println(edges.Len())
	forest := Kruskal(edges, len(boxes))
	counts := make(map[Box]int)
	for _, b := range boxes {
		boxSet := forest.Find(b)
		if boxSet == nil {
			continue // this box is not connected to anything
		}
		box := boxSet.(Box)
		if _, ok := counts[box]; !ok {
			counts[box] = 1
		} else {
			counts[box]++
		}
	}
	// pretty.Println(counts)

	cts := slices.SortedFunc(maps.Values(counts), func(a, b int) int { return cmp.Compare(b, a) })
	fmt.Println(cts)
	firstThree := cts[:3]
	fmt.Println(firstThree)
	return iters.Product(slices.Values(firstThree))
}

// https://en.wikipedia.org/wiki/Kruskal%27s_algorithm
func Kruskal(edges Edges, verticesCount int) *dsu.DSU {
	forest := dsu.New()
	count := 0
	for _, e := range edges {
		forest.Add(e.A)
		forest.Add(e.B)
		if forest.Find(e.A) != forest.Find(e.B) {
			forest.Union(e.A, e.B)
			count++
			if count == verticesCount-1 {
				fmt.Println(e.A, e.B)
				fmt.Println(e.A.X * e.B.X)
				break
			}
		}
	}

	return forest
}

type Box struct{ X, Y, Z int }

func ParseBox(line string) Box {
	nums := slices.Collect(iters.Map(iter.Seq[strpipe.StrPipe](strpipe.New(line).Trim().SplitCommas()), func(s strpipe.StrPipe) int { n, _ := strconv.Atoi(s.String()); return n }))
	x, y, z := nums[0], nums[1], nums[2]
	return Box{X: x, Y: y, Z: z}
}

type Edge struct {
	Distance int
	A, B     Box
}

type Edges []Edge

func (a Edges) Len() int           { return len(a) }
func (a Edges) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
func (a Edges) Less(i, j int) bool { return a[i].Distance < a[j].Distance }

func MakeEdges(boxes []Box) Edges {
	edges := Edges{}
	for i, a := range boxes {
		for _, b := range boxes[i+1:] {
			if a == b {
				continue
			}
			edges = append(edges, Edge{Distance(a, b), a, b})
		}
	}
	sort.Sort(edges)
	return edges
}

func Distance(a, b Box) int {
	x := a.X - b.X
	x *= x
	y := a.Y - b.Y
	y *= y
	z := a.Z - b.Z
	z *= z
	return x + y + z
}
