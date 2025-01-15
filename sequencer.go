package goedbt

import "iter"

// Sequencer defines a Behaviour BehaviourNode that checks each of its children,
// returning the first non-Success status or Success if all children succeed
type Sequencer struct {
	*composite

	seq iter.Seq[Behaviour]
	next
}

func NewSequencer() *Sequencer {
	return &Sequencer{
		composite: &composite{
			node:     &node{state: Invalid},
			children: make(Set[Behaviour]),
		},
	}
}

func (n *Sequencer) Initialize() {
	n.seq, n.next = n.Children()
}

func (n *Sequencer) Update() Status {
	for c := range n.seq {
		if status := tick(c); status != Success {
			n.state = status
			return n.state
		}
		n.next()
	}

	n.state = Success
	return n.state
}

func (n *Sequencer) Teardown() {}
func (n *Sequencer) Abort()    {}
