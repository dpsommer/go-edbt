package goedbt

type Behaviour interface {
	State() Status
	Initialize()
	Update() Status
	Terminate()
}

type node struct {
	state Status
}

func (n *node) State() Status { return n.state }

func tick(b Behaviour) Status {
	if b.State() != Running {
		b.Initialize()
	}

	state := b.Update()
	if state != Running {
		b.Terminate()
	}

	return state
}
