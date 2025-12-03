package day03

import (
	"aoc2025/util"
	"aoc2025/util/iters"
	"aoc2025/util/strpipe"
	"fmt"
	"iter"
	"slices"
)

func Day03(filename string, part int) {
	banks := util.LoadLines(filename, Parse)
	switch part {
	case 1:
		fmt.Println(Part01(banks))
	case 2:
		fmt.Println(Part02(banks))
	}
}

func Part01(banks iter.Seq[Bank]) int {
	return iters.Sum(iters.Map(banks, Bank.LargestPair))
}
func Part02(banks iter.Seq[Bank]) int {
	return iters.Sum(iters.Map(banks, Bank.LargestPair))
}

type Bank []int

func Parse(s string) Bank {
	return Bank(slices.Collect(strpipe.New(s).Trim().ToInts(1)))
}

func (b Bank) LargestPair() int {
	m := 0
	for i := 0; i < len(b)-1; i++ {
		for j := i + 1; j < len(b); j++ {
			a, b := b[i], b[j]
			num := a*10 + b
			m = max(m, num)
		}
	}
	return m
}
