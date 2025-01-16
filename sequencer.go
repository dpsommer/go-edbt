package goedbt

// Sequencer defines a composite Behaviour that checks each of its children,
// returning the first non-Success status or Success if all children succeed
type Sequencer struct {
	*BehaviourTree
	*composite

	iterator[Behaviour]
}

// XXX: don't love having to pass the BT to each node - can this even be called
// a tree at this point?
func NewSequencer(bt *BehaviourTree) *Sequencer {
	return &Sequencer{
		BehaviourTree: bt,
		composite: &composite{
			behaviour: &behaviour{state: Invalid},
			children:  make(Set[Behaviour]),
		},
	}
}

func (n *Sequencer) initialize() {
	n.iterator = n.Children()
	n.pushCurrentChild()
}

func (n *Sequencer) pushCurrentChild() {
	n.Start(&Event{
		Behaviour: n.current(),
		Observer:  n.onChildComplete,
	})
}

func (n *Sequencer) onChildComplete(s Status) {
	c := n.current()
	switch s := c.State(); s {
	case Failure:
		n.Stop(&Event{n, nil}, s)
	case Running:
		n.pushCurrentChild()
	case Success:
		if _, ok := n.next(); !ok {
			// reached last child, set state to success
			n.Stop(&Event{n, nil}, s)
			return
		}
		n.pushCurrentChild()
	}
}

func (n *Sequencer) update() Status {
	n.state = Running
	return Running
}

func (n *Sequencer) teardown() {}
func (n *Sequencer) abort()    {}
