package goedbt

// Sequencer defines a composite Behaviour that checks each of its children,
// returning the first non-Success status or Success if all children succeed
type Sequencer struct {
	*composite

	iterator[Behaviour]
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
	n.iterator = n.Children()
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
