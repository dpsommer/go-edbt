package goedbt

import "iter"

// Sequencer defines a Behaviour BehaviourNode that checks each of its children,
// returning the first non-Success status or Success if all children succeed
type Sequencer struct {
	*node
	*composite

	seq  iter.Seq[Behaviour]
	next func()
}

func NewSequencer() *Sequencer {
	return &Sequencer{
		node: &node{state: Invalid},
		composite: &composite{
			children: make(map[Behaviour]struct{}),
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

func (n *Sequencer) Terminate() {}
