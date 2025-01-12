package goedbt

import "fmt"

type RootNode struct {
	child Node
}

func (n *RootNode) Tick() Status {
	return n.child.Tick()
}

func (n *RootNode) Children() []Node {
	return []Node{n.child}
}

func (n *RootNode) AddChild(child Node) error {
	if n.child == nil {
		n.child = child
		return nil
	}

	return fmt.Errorf("root nodes can only have a single child")
}
