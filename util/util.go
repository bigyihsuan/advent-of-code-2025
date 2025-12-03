package util

import (
	"math"
	"slices"
)

func SliceWithoutIndexes[T any, S ~[]T](s S, idxs ...int) S {
	o := S{}
	for i, e := range s {
		if slices.Contains(idxs, i) {
			continue
		}
		o = append(o, e)
	}
	return o
}

func IntFromDigits(digits []int) int {
	n := 0
	for i, d := range digits {
		n += d * int(math.Pow10(len(digits)-i-1))
	}
	return n
}
