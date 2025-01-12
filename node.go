package goedbt

type Status int

const (
	Success Status = iota
	Failure
	Running
)

type Node interface {
	Tick() Status
	Children() []Node
	AddChild(child Node) error
}
