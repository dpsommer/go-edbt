package goedbt

import "fmt"

type SuccessTask struct {
}

func (n *SuccessTask) Tick() Status {
	return Success
}

func (n *SuccessTask) Children() []Node {
	return nil
}

func (n *SuccessTask) AddChild(child Node) error {
	return fmt.Errorf("task nodes cannot have children")
}

type FailureTask struct {
}

func (n *FailureTask) Tick() Status {
	return Failure
}

func (n *FailureTask) Children() []Node {
	return nil
}

func (n *FailureTask) AddChild(child Node) error {
	return fmt.Errorf("task nodes cannot have children")
}
