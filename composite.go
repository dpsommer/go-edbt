package goedbt

import utils "github.com/dpsommer/go-utils"

type Composite interface {
	Behaviour

	Children() utils.Iterator[Behaviour]
	AddChild(child Behaviour)
	RemoveChild(child Behaviour)
	ClearChildren()
}

type composite struct {
	*behaviour

	tree     *BehaviourTree
	children utils.Set[Behaviour]
}

func (n *composite) Children() utils.Iterator[Behaviour] {
	// copy the children map keys to a list so that modifications to it
	// while we hold an active iterator don't affect iteration. use a list
	// so that we can replay the same keys if needed
	cc := utils.Keys(n.children)
	return utils.NewIterator(cc)
}

func (n *composite) AddChild(child Behaviour) {
	n.children[child] = struct{}{}
}

func (n *composite) RemoveChild(child Behaviour) {
	// ensure that the node is re-initialized on next tick
	n.state = Aborted
	delete(n.children, child)
}

func (n *composite) ClearChildren() {
	n.state = Aborted
	n.children = make(utils.Set[Behaviour])
}
