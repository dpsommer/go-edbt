package goedbt

// Selector defines a composite Behaviour that checks each of its children,
// returning the first non-Failure status or Failure if all children fail
type Selector struct {
	*composite

	iterator[Behaviour]
}

func NewSelector() *Selector {
	return &Selector{
		composite: &composite{
			node:     &node{state: Invalid},
			children: make(Set[Behaviour]),
		},
	}
}

func (n *Selector) Initialize() {
	n.iterator = n.Children()
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

func (n *Selector) Teardown() {}
func (n *Selector) Abort()    {}
