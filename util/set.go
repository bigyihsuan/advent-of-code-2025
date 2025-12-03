package util

type Set[T comparable] map[T]struct{}

func (s Set[T]) Add(t T) {
	s[t] = struct{}{}
}
