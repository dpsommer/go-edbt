package goedbt

type Status int

const (
	Success Status = iota
	Failure
	Running
)

type Node interface {
	Tick() Status
}
