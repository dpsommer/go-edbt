package goedbt

type Composite interface {
	Children() []Behaviour
	AddChild(child Behaviour)
	RemoveChild(child Behaviour)
	ClearChildren()
}

type composite struct {
	// Golang has no Set structure, so use a map to mimic one
	children map[Behaviour]struct{}
}

func (n *composite) Children() []Behaviour {
	return Keys(n.children)
}

func (n *composite) AddChild(child Behaviour) {
	n.children[child] = struct{}{}
}

func (n *composite) RemoveChild(child Behaviour) {
	delete(n.children, child)
}

func (n *composite) ClearChildren() {
	n.children = make(map[Behaviour]struct{})
}
