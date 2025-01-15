package goedbt

import "iter"

type iterator[T any] struct {
	seq iter.Seq[T]
	next
}

type Set[K comparable] map[K]struct{}

func keys[K comparable, T any](m map[K]T) []K {
	keys := make([]K, len(m))

	i := 0
	for k := range m {
		keys[i] = k
		i++
	}

	return keys
}

func pop[T any](s []T) (T, []T) {
	return s[len(s)-1], s[:len(s)-1]
}
