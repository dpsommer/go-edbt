package goedbt

// SuccessBehaviour
type SuccessBehaviour struct {
	behaviour
}

func (n *SuccessBehaviour) initialize()    {}
func (n *SuccessBehaviour) update() Status { return Success }
func (n *SuccessBehaviour) teardown()      {}
func (n *SuccessBehaviour) abort()         {}

// FailureBehaviour
type FailureBehaviour struct {
	behaviour
}

func (n *FailureBehaviour) initialize()    {}
func (n *FailureBehaviour) update() Status { return Failure }
func (n *FailureBehaviour) teardown()      {}
func (n *FailureBehaviour) abort()         {}

// RunningBehaviour
type RunningBehaviour struct {
	behaviour
}

func (n *RunningBehaviour) initialize()    {}
func (n *RunningBehaviour) update() Status { return Running }
func (n *RunningBehaviour) teardown()      {}
func (n *RunningBehaviour) abort()         {}

// XThenY
type XThenY struct {
	behaviour

	accessed bool
	X, Y     Status
}

func (n *XThenY) initialize() { n.state = n.X }
func (n *XThenY) update() Status {
	if n.accessed {
		n.state = n.Y
		return n.Y
	}
	n.accessed = true
	return n.X
}
func (n *XThenY) teardown() {}
func (n *XThenY) abort()    {}
