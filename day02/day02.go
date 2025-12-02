package day02

import (
	"aoc2025/day02/ranges"
	"aoc2025/util"
	"aoc2025/util/iters"
	"fmt"
	"iter"
	"strconv"
	"strings"
)

func Day02(filename string, part int) {
	rs := util.LoadLineOfElements(filename, Parse)
	switch part {
	case 1:
		fmt.Println(Part01(rs))
	case 2:
		fmt.Println(Part02(rs))
	}
}

func Part01(rs iter.Seq[ranges.Range]) int {
	sum := 0
	for r := range rs {
		sum += iters.Sum(r.InvalidIds())
		// fmt.Println(r, slices.Collect(r.InvalidIds()))
	}
	return sum
}
func Part02(rs iter.Seq[ranges.Range]) int {
	panic("TODO")
	// return -1
}

func Parse(s string) ranges.Range {
	ss := strings.Split(s, "-")
	startStr, endStr := ss[0], ss[1]
	start, err := strconv.Atoi(startStr)
	if err != nil {
		panic(fmt.Errorf("parsing range: %w", err))
	}
	end, err := strconv.Atoi(endStr)
	if err != nil {
		panic(fmt.Errorf("parsing range: %w", err))
	}
	return ranges.Range{
		Start: start,
		End:   end,
	}
}
