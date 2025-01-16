package goedbt

type Observer func(s Status)

type Behaviour interface {
	State() Status
	SetState(s Status)

	Initialize()
	Update() Status
	Teardown()
	Abort()

	FireObserver(s Status)
}

type behaviour struct {
	state Status

	Observer
}

func (n *behaviour) State() Status         { return n.state }
func (n *behaviour) SetState(s Status)     { n.state = s }
func (n *behaviour) FireObserver(s Status) { n.Observer(s) }

func tick(b Behaviour) Status {
	if b.State() != Running {
		b.Initialize()
	}

	state := b.Update()
	if state != Running {
		b.Teardown()
	}

	return state
}
