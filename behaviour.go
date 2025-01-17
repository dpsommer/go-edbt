package goedbt

type Behaviour interface {
	State() Status
	SetState(s Status)

	initialize()
	update() Status
	teardown()
	abort()
}

type behaviour struct {
	state Status
}

func (n *behaviour) State() Status { return n.state }
func (n *behaviour) SetState(s Status) {
	n.state = s
}

func tick(b Behaviour) Status {
	if b.State() != Running {
		b.initialize()
	}

	state := b.update()
	b.SetState(state)

	if state != Running {
		b.teardown()
	}

	return state
}
