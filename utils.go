package goedbt

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

func pop[K any](s []K) (K, []K) {
	return s[len(s)-1], s[:len(s)-1]
}
