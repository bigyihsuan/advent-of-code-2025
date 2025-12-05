package ranges

import (
	"cmp"
	"fmt"
	"slices"
	"strings"
)

type Range struct{ Start, End int }

func (r Range) String() string {
	return fmt.Sprintf("%d-%d", r.Start, r.End)
}

func (r Range) Len() int {
	return r.End - r.Start + 1
}

// checks if two ranges overlap.
func (r Range) Overlaps(o Range) bool {
	a := r.Start <= o.Start && o.Start < r.End && r.End <= o.End // o starts in r, ends after r
	b := o.Start <= r.Start && r.Start < o.End && o.End <= r.End // o starts before r, ends in r
	c := r.Start <= o.Start && o.End <= r.End                    // o is inside or equal to r
	return a || b || c
}

func (r Range) SharesEndPoint(o Range) bool {
	return r.Start == o.End || r.End == o.Start
}

func (r Range) CanMerge(o Range) bool {
	return r.Overlaps(o) || r.SharesEndPoint(o)
}

// joins together two ranges if and only if they overlap.
// returns the original range if they do not overlap.
func (r Range) Join(o Range) Range {
	if !r.CanMerge(o) {
		return r
	}
	return Range{
		Start: min(r.Start, o.Start),
		End:   max(r.End, o.End),
	}
}

func RangeJoinMany(rs ...Range) Range {
	r := rs[0]
	for _, rb := range rs[1:] {
		r = r.Join(rb)
	}
	return r
}

func (r Range) Contains(n int) bool {
	return r.Start <= n && n <= r.End
}

type Ranges []Range

func New() Ranges {
	return Ranges{}
}

func (r Ranges) String() string {
	var sb strings.Builder
	sb.WriteString("{")
	es := []string{}
	for _, e := range r {
		es = append(es, fmt.Sprint(e))
	}
	sb.WriteString(strings.Join(es, ","))
	sb.WriteString("}")
	return sb.String()
}

func (r Ranges) HasN(n int) bool {
	for _, ra := range r {
		if ra.Contains(n) {
			return true
		}
	}
	return false
}

func (r Ranges) Merge() Ranges {
	slices.SortFunc(r, func(a, b Range) int {
		return cmp.Or(
			cmp.Compare(a.Start, b.Start),
			-cmp.Compare(a.Len(), b.Len()),
		)
	})
	fmt.Println(r)
	out := []Range{r[0]}
	for i := 1; i < len(r); i++ {
		last := out[len(out)-1]
		curr := r[i]
		if last.SharesEndPoint(curr) || last.Overlaps(curr) {
			out[len(out)-1] = last.Join(curr)
		} else {
			out = append(out, curr)
		}
	}
	return out
}
