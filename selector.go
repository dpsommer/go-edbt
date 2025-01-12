package goedbt

type SelectorNode struct{}

func (n *SelectorNode) Tick() Status {
	return Success
}
