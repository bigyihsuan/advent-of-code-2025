package iters

import "iter"

func Map[T, U any](ts iter.Seq[T], f func(T) U) iter.Seq[U] {
	return func(yield func(U) bool) {
		for t := range ts {
			if !yield(f(t)) {
				return
			}
		}
	}
}

func Sum(ints iter.Seq[int]) int {
	s := 0
	for i := range ints {
		s += i
	}
	return s
}

func Product(ints iter.Seq[int]) int {
	s := 1
	for i := range ints {
		s *= i
	}
	return s
}
