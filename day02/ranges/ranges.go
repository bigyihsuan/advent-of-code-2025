package ranges

import (
	"iter"
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
