package day06

import (
	"aoc2025/util"
	"aoc2025/util/iters"
	"aoc2025/util/strpipe"
	"fmt"
	"iter"
	"regexp"
	"slices"
	"strconv"
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
	lines := strings.Lines(util.LoadFile(filename))
	equations := Parse(lines)
	fmt.Println(equations)
	return iters.Sum(iters.Map(slices.Values(equations), Equation.Eval))
}

func Part2(filename string) int {
	lines := slices.Collect(strings.Lines(util.LoadFile(filename)))

	n, o := lines[:len(lines)-1], lines[len(lines)-1]
	grid := slices.Collect(
		iters.Map(slices.Values(n),
			func(s string) []string {
				return slices.Collect(
					iters.Map(slices.Values([]rune(strings.Trim(s, "\n"))), func(r rune) string { return string(r) }),
				)
			},
		),
	)
	grid = util.RotateGridCCW(grid)
	nums := [][]int{}
	numRow := []int{}
	for _, row := range grid {
		if c := slices.Clone(row); slices.Equal(slices.Compact(c), []string{" "}) {
			nums = append(nums, numRow)
			numRow = []int{}
			continue
		}
		n, _ := strconv.Atoi(strings.TrimSpace(strings.Join(row, "")))
		numRow = append(numRow, n)
	}
	nums = append(nums, numRow)

	rs := []rune(o)
	slices.Reverse(rs)
	rev := string(rs)
	ops := strings.Fields(strings.TrimSpace(rev))

	eqs := []Equation{}
	for i, op := range ops {
		var e Equation
		switch op {
		case string(PLUS):
			e.Operation = PLUS
		case string(STAR):
			e.Operation = STAR
		}
		e.Nums = nums[i]
		eqs = append(eqs, e)
	}
	return iters.Sum(iters.Map(slices.Values(eqs), Equation.Eval))
}

type Equation struct {
	Nums []int
	Operation
}

func (e Equation) Eval() int {
	switch e.Operation {
	case PLUS:
		return iters.Sum(slices.Values(e.Nums))
	case STAR:
		return iters.Product(slices.Values(e.Nums))
	default:
		panic(fmt.Errorf("equation eval: should have an PLUS or STAR, but got %q", e.Operation))
	}
}

type Operation string

const (
	PLUS Operation = "+"
	STAR Operation = "*"
)

func Parse(l iter.Seq[string]) []Equation {
	eqs := []Equation{}
	lines := strpipe.StrPipeIter(iters.Map(l, func(s string) strpipe.StrPipe { return strpipe.New(s).Trim().Replace(*regexp.MustCompile(`\s+`), " ") })).CollectStrings()
	n, o := lines[:len(lines)-1], lines[len(lines)-1]
	ops := strings.Fields(o)
	nums := slices.Collect(iters.Map(slices.Values(n), func(s string) []int {
		return slices.Collect(iters.Map(strings.FieldsSeq(s), func(nn string) int { n, _ := strconv.Atoi(nn); return n }))
	}))

	for i, op := range ops {
		var e Equation
		switch op {
		case string(PLUS):
			e.Operation = PLUS
		case string(STAR):
			e.Operation = STAR
		}
		for _, row := range nums {
			e.Nums = append(e.Nums, row[i])
		}
		eqs = append(eqs, e)
	}
	return eqs
}
