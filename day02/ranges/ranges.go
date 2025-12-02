package ranges

import (
	"aoc2025/util/iters"
	"iter"
	"slices"
	"strconv"
)

type Range struct{ Start, End int }

func (r Range) InvalidIds() iter.Seq[int] {
	return func(yield func(int) bool) {
		for i := r.Start; i <= r.End; i++ {
			s := strconv.Itoa(i)
			if len(s)%2 == 0 {
				// check for doubled
				half := len(s) / 2
				first, last := s[:half], s[half:]
				if first == last && !yield(i) {
					return
				}
			}
		}
	}
}

func (r Range) InvalidIds2() iter.Seq[int] {
	return func(yield func(int) bool) {
		for i := r.Start; i <= r.End; i++ {
			s := strconv.Itoa(i)
			half := len(s) / 2
			for n := half; n > 0; n-- {
				chunks := slices.Collect(iters.Map(slices.Chunk([]rune(s), n), func(r []rune) string { return string(r) }))
				compacted := slices.Compact(chunks)
				if len(compacted) == 1 {
					// fmt.Println(r, i)
					if !yield(i) {
						return
					}
					break
				}
			}
		}
	}
}
