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
	lines := strings.Lines(util.LoadFile(filename))
	equations := Parse(lines)
	fmt.Println(equations)
	switch part {
	case 1:
		fmt.Println(Part1(equations))
	case 2:
		fmt.Println(Part2(equations))
	}
}

func Part1(equations []Equation) int {
	return iters.Sum(iters.Map(slices.Values(equations), Equation.Eval))
}

func Part2(equations []Equation) int {
	return -1
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
