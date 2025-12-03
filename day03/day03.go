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
	return iters.Sum(iters.Map(banks, Bank.LargestTwelve))
}

type Bank []int

func Parse(s string) Bank {
	return Bank(slices.Collect(strpipe.New(s).Trim().ToInts(1)))
}

func (bank Bank) LargestPair() int {
	m := 0
	for i := 0; i < len(bank)-1; i++ {
		for j := i + 1; j < len(bank); j++ {
			a, b := bank[i], bank[j]
			num := a*10 + b
			m = max(m, num)
		}
	}
	return m
}

func (bank Bank) LargestTwelve() int {
	digits := []int{}
	skips := len(bank) - 12
	bankClone := slices.Clone(bank)
	for skips > 0 && len(digits) < 12 {
		index := slices.Index(bankClone, slices.Max(bankClone[:skips+1]))
		digit := slices.Max(bankClone[:skips+1])
		digits = append(digits, digit)
		skips -= index
		bankClone = bankClone[index+1:]
	}
	// no more skips, add rest of the digits
	digits = append(digits, bankClone[:12-len(digits)]...)
	num := util.IntFromDigits(digits)
	return num
}
