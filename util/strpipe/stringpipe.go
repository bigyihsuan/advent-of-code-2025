package strpipe

import (
	"aoc2025/util/iters"
	"iter"
	"slices"
	"strconv"
	"strings"
)

type StrPipe string
type StrPipeIter iter.Seq[StrPipe]

func New(s string) StrPipe {
	return StrPipe(s)
}

func NewFromRunes(rs []rune) StrPipe {
	return New(string(rs))
}

func (sp StrPipe) String() string {
	return string(sp)
}

func (sp StrPipe) Runes() []rune {
	return []rune(sp)
}

// chaining

func (sp StrPipe) Trim() StrPipe {
	return StrPipe(strings.TrimSpace(string(sp)))
}

// collections

func (sp StrPipe) SplitCommas() iter.Seq[StrPipe] {
	return iters.Map(slices.Values(strings.Split(string(sp), ",")), New)
}

func (sp StrPipe) Chunk(n int) iter.Seq[StrPipe] {
	return func(yield func(StrPipe) bool) {
		chunks := iters.Map(slices.Chunk(sp.Runes(), n), NewFromRunes)
		for chunk := range chunks {
			if !yield(chunk) {
				return
			}
		}
	}
}

func (sp StrPipe) ToInts(digits int) iter.Seq[int] {
	return func(yield func(int) bool) {
		chunks := iters.Map(sp.Chunk(digits), func(s StrPipe) int {
			i, err := strconv.Atoi(s.String())
			if err != nil {
				panic(err)
			}
			return i
		})
		for chunk := range chunks {
			if !yield(chunk) {
				return
			}
		}
	}
}

// StrPipes

func (it StrPipeIter) CollectPipes() []StrPipe {
	return slices.Collect(iter.Seq[StrPipe](it))
}

func (it StrPipeIter) CollectStrings() []string {
	return slices.Collect(iters.Map(iter.Seq[StrPipe](it), StrPipe.String))
}
