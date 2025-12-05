package day05

import (
	"aoc2025/day05/ranges"
	"aoc2025/util/iters"
	"aoc2025/util/set"
	"aoc2025/util/strpipe"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func Day(filename string, part int) {
	// r := ranges.New()
	// r.AddRange(10, 18)
	// fmt.Println(r)
	// r.AddRange(16, 20)
	// fmt.Println(r)
	d := Input(filename)
	fmt.Println(d.Fresh)
	switch part {
	case 1:
		fmt.Println(Part01(d))
	case 2:
		fmt.Println(Part02(d))
	}
}

type Database struct {
	Fresh     ranges.Ranges
	Inventory set.Set[int]
}

func Part01(d Database) int {
	spoiled := 0
	for id := range d.Inventory {
		if d.Fresh.HasN(id) {
			spoiled++
			continue
		}
	}
	return spoiled
}

func Part02(d Database) int {
	d.Fresh = d.Fresh.Merge()
	fmt.Println(d.Fresh)
	freshIds := 0
	for _, r := range d.Fresh {
		ids := r.End - r.Start + 1
		fmt.Println(ids, r)
		freshIds += ids
	}
	return freshIds
}

func Input(filename string) Database {
	bts, err := os.ReadFile(filename)
	if err != nil {
		panic(err)
	}
	pts := strings.Split(string(bts), "\n\n")
	freshStr, invStr := pts[0], pts[1]

	fresh := ranges.New()
	for line := range strings.Lines(freshStr) {
		l := strpipe.New(line).Trim().SplitDash().CollectStrings()
		start, err := strconv.Atoi(l[0])
		if err != nil {
			panic(err)
		}
		end, err := strconv.Atoi(l[1])
		if err != nil {
			panic(err)
		}
		fresh = append(fresh, ranges.Range{Start: start, End: end})
	}

	inventory := set.New[int]()
	for line := range iters.Map(strpipe.New(invStr).Lines(), strings.TrimSpace) {
		id, err := strconv.Atoi(line)
		if err != nil {
			panic(err)
		}
		inventory.Add(id)
	}
	return Database{
		Fresh:     fresh,
		Inventory: inventory,
	}
}
