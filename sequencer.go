package goedbt

import utils "github.com/dpsommer/go-utils"

// Sequencer defines a composite Behaviour that checks each of its children,
// returning the first non-Success status or Success if all children succeed
type Sequencer struct {
	*composite

	utils.Iterator[Behaviour]
}

// XXX: don't love having to pass the BT to each node - can this even be called
// a tree at this point?
func NewSequencer(bt *BehaviourTree) *Sequencer {
	return &Sequencer{
		composite: &composite{
			tree:      bt,
			behaviour: &behaviour{state: Invalid},
			children:  make(utils.Set[Behaviour]),
		},
	}
}

func (n *Sequencer) initialize() {
	n.Iterator = n.Children()
	n.tree.Start(n.Current(), n.onChildComplete)
}

func (n *Sequencer) onChildComplete(s Status) {
	c := n.Current()

	switch s := c.State(); s {
	case Success:
		if _, ok := n.Next(); !ok {
			// reached last child, set state to success
			n.tree.Stop(&Event{n, nil}, s)
			return
		}
		n.tree.Start(n.Current(), n.onChildComplete)
	default: // handle failure and aborted cases
		n.tree.Stop(&Event{n, nil}, Failure)
	}
}

func (n *Sequencer) update() Status {
	return Running
}

func (n *Sequencer) teardown() {}
func (n *Sequencer) abort()    {}
