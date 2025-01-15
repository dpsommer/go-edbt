package goedbt

type Status int

const (
	Invalid Status = iota - 1
	Success
	Failure
	Running
	Aborted
)
