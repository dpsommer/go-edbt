package goedbt

type Observer func(Status)

type Event struct {
	Behaviour
	Observer
}
