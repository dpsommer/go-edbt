package goedbt

import utils "github.com/dpsommer/go-utils"

// Selector defines a composite Behaviour that checks each of its children,
// returning the first non-Failure status or Failure if all children fail
type Selector struct {
	*composite

	utils.Iterator[Behaviour]
}

func NewSelector(bt *BehaviourTree) *Selector {
	return &Selector{
		composite: &composite{
			tree:      bt,
			behaviour: &behaviour{state: Invalid},
			children:  make(utils.Set[Behaviour]),
		},
	}
}

func (n *Selector) initialize() {
	n.Iterator = n.Children()
	n.tree.Start(n.Current(), n.onChildComplete)
}

func (n *Selector) onChildComplete(s Status) {
	c := n.Current()

	switch s := c.State(); s {
	case Success:
		n.tree.Stop(&Event{n, nil}, s)
	default:
		if _, ok := n.Next(); !ok {
			// reached last child, set state to failure
			n.tree.Stop(&Event{n, nil}, Failure)
			return
		}
		n.tree.Start(n.Current(), n.onChildComplete)
	}
}

func (n *Selector) update() Status {
	return Running
}

func (n *Selector) teardown() {}
func (n *Selector) abort()    {}
