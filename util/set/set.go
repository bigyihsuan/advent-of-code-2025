package set

import (
	"fmt"
	"strings"
)

type Set[T comparable] map[T]struct{}

func New[T comparable]() Set[T] {
	return make(Set[T])
}

func (s Set[T]) Add(t T) Set[T] {
	s[t] = struct{}{}
	return s
}

func (s Set[T]) Remove(t T) Set[T] {
	delete(s, t)
	return s
}

func (s Set[T]) Has(t T) bool {
	_, ok := s[t]
	return ok
}

func (s Set[T]) String() string {
	var sb strings.Builder
	sb.WriteString("{")
	es := []string{}
	for e := range s {
		es = append(es, fmt.Sprint(e))
	}
	sb.WriteString(strings.Join(es, ","))
	sb.WriteString("}")
	return sb.String()
}
