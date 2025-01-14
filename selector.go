package goedbt

// Selector defines a Behaviour BehaviourNode that checks each of its children,
// returning the first non-Failure status or Failure if all children fail
type Selector struct {
	*node
	*composite
}

func NewSelector() *Selector {
	return &Selector{
		node: &node{state: Invalid},
		composite: &composite{
			children: make(map[Behaviour]struct{}),
		},
	}
}

func (n *Selector) Initialize() {}
func (n *Selector) Terminate()  {}

func (n *Selector) Update() Status {
	for c := range n.children {
		if status := tick(c); status != Failure {
			n.state = status
			return n.state
		}
	}

	n.state = Failure
	return n.state
}
