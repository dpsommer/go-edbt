package goedbt

// Selector defines a composite Behaviour that checks each of its children,
// returning the first non-Failure status or Failure if all children fail
type Selector struct {
	*BehaviourTree
	*composite

	iterator[Behaviour]
}

func NewSelector(bt *BehaviourTree) *Selector {
	return &Selector{
		BehaviourTree: bt,
		composite: &composite{
			behaviour: &behaviour{state: Invalid},
			children:  make(Set[Behaviour]),
		},
	}
}

func (n *Selector) initialize() {
	n.iterator = n.Children()
}

func (n *Selector) update() Status {
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

func (n *Selector) teardown() {}
func (n *Selector) abort()    {}
