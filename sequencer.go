package goedbt

// Sequencer defines a Behaviour BehaviourNode that checks each of its children,
// returning the first non-Success status or Success if all children succeed
type Sequencer struct {
	*node
	*composite
}

func NewSequencer() *Sequencer {
	return &Sequencer{
		node: &node{state: Invalid},
		composite: &composite{
			children: make(map[Behaviour]struct{}),
		},
	}
}

func (n *Sequencer) Initialize() {}
func (n *Sequencer) Terminate()  {}

func (n *Sequencer) Update() Status {
	for c := range n.children {
		if status := tick(c); status != Success {
			n.state = status
			return n.state
		}
	}

	n.state = Success
	return n.state
}
