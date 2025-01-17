package goedbt

import "iter"

type next[T any] func() (T, bool)

type current[T any] func() T

type iterator[T any] struct {
	seq iter.Seq[T]
	current[T]
	next[T]
}

func newIterator[T any](tt []T) iterator[T] {
	var i int

	// return an iterator and a closure that increments the iterator so that we
	// can resume iteration from the same key if a child is running
	return iterator[T]{
		seq: func(yield func(T) bool) {
			for j := 0; j < len(tt); j += 1 {
				if !yield(tt[j]) {
					return
				}
			}
		},
		current: func() T { return tt[i] },
		next: func() (T, bool) {
			if i+1 >= len(tt) {
				return *new(T), false
			}
			i += 1
			return tt[i], true
		},
	}
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
