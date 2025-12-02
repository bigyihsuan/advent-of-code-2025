package strpipe

import (
	"aoc2025/util/iters"
	"iter"
	"slices"
	"strings"
)

type StrPipe string

func New(s string) StrPipe {
	return StrPipe(s)
}

func (sp StrPipe) String() string {
	return string(sp)
}

// chaining

func (sp StrPipe) Trim() StrPipe {
	return StrPipe(strings.TrimSpace(string(sp)))
}

// collections

func (sp StrPipe) SplitCommas() iter.Seq[StrPipe] {
	return iters.Map(slices.Values(strings.Split(string(sp), ",")), New)
}
