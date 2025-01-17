package goedbt

// Selector defines a composite Behaviour that checks each of its children,
// returning the first non-Failure status or Failure if all children fail
type Selector struct {
	*composite

	iterator[Behaviour]
}

func NewSelector(bt *BehaviourTree) *Selector {
	return &Selector{
		composite: &composite{
			tree:      bt,
			behaviour: &behaviour{state: Invalid},
			children:  make(Set[Behaviour]),
		},
	}
}

func (n *Selector) initialize() {
	n.iterator = n.Children()
	n.tree.Start(n.current(), n.onChildComplete)
}

func (n *Selector) onChildComplete(s Status) {
	c := n.current()

	switch s := c.State(); s {
	case Success:
		n.tree.Stop(&Event{n, nil}, s)
	default:
		if _, ok := n.next(); !ok {
			// reached last child, set state to failure
			n.tree.Stop(&Event{n, nil}, Failure)
			return
		}
		n.tree.Start(n.current(), n.onChildComplete)
	}
}

func (n *Selector) update() Status {
	return Running
}

func (n *Selector) teardown() {}
func (n *Selector) abort()    {}
