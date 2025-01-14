package goedbt

type Status int

const (
	Invalid Status = -1
	Success Status = iota
	Failure
	Running
)
