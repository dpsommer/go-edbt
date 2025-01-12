package goedbt

type SequencerNode struct {
	children []Node
}

func (n *SequencerNode) Tick() Status {
	for _, c := range n.children {
		status := c.Tick()
		if status == Failure {
			return Failure
		}
	}
	return Success
}

func (n *SequencerNode) Children() []Node {
	return n.children
}

func (n *SequencerNode) AddChild(child Node) error {
	n.children = append(n.children, child)
	return nil
}
