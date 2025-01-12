package goedbt

type RootNode struct{}

func (n *RootNode) Tick() Status {
	return Success
}
