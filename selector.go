package goedbt

import (
	"iter"
)

// Selector defines a Behaviour BehaviourNode that checks each of its children,
// returning the first non-Failure status or Failure if all children fail
type Selector struct {
	*node
	*composite

	seq  iter.Seq[Behaviour]
	next func()
}

func NewSelector() *Selector {
	return &Selector{
		node: &node{state: Invalid},
		composite: &composite{
			children: make(map[Behaviour]struct{}),
		},
	}
}

func (n *Selector) Initialize() {
	n.seq, n.next = n.Children()
}

func (n *Selector) Update() Status {
	for c := range n.seq {
		status := tick(c)
		if status != Failure {
			n.state = status
			return n.state
		}
		n.next()
	}

	n.state = Failure
	return n.state
}

func (n *Selector) Terminate() {}
