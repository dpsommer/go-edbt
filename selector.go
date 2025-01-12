package goedbt

type SelectorNode struct {
	children []Node
}

func (n *SelectorNode) Tick() Status {
	// Selector nodes implement OR behaviour - checking each child in sequence
	// and returning on the first Success or Failure if all children fail
	for _, child := range n.children {
		status := child.Tick()
		if status == Success {
			return Success
		}
	}

	return Failure
}

func (n *SelectorNode) Children() []Node {
	return n.children
}

func (n *SelectorNode) AddChild(child Node) error {
	n.children = append(n.children, child)
	return nil
}
