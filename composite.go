package goedbt

import (
	"iter"
)

type next func()

type Composite interface {
	Behaviour

	Children() (iter.Seq[Behaviour], next)
	AddChild(child Behaviour)
	RemoveChild(child Behaviour)
	ClearChildren()
}

type composite struct {
	*node

	// Golang has no Set structure, so use a map to mimic one
	children map[Behaviour]struct{}
}

func (n *composite) Children() (iter.Seq[Behaviour], next) {
	// copy the children map keys to a list so that modifications to it
	// while we hold an active iterator don't affect iteration. use a list
	// so that we can replay the same keys if needed
	cc := make([]Behaviour, len(n.children))

	var i int
	for k := range n.children {
		cc[i] = k
		i += 1
	}

	i = 0

	// return an iterator and a closure that increments the iterator so that we
	// can resume iteration from the same key if a child is running
	return func(yield func(Behaviour) bool) {
		for {
			if i >= len(cc) {
				return
			}
			if !yield(cc[i]) {
				return
			}
		}
	}, func() { i += 1 }
}

func (n *composite) AddChild(child Behaviour) {
	n.children[child] = struct{}{}
}

func (n *composite) RemoveChild(child Behaviour) {
	// ensure that the node is re-initialized on next tick
	n.state = Aborted
	delete(n.children, child)
}

func (n *composite) ClearChildren() {
	n.state = Aborted
	n.children = make(map[Behaviour]struct{})
}
