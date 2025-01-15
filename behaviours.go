package goedbt

// SuccessBehaviour
type SuccessBehaviour struct {
	node
}

func (n *SuccessBehaviour) Initialize() {}
func (n *SuccessBehaviour) Update() Status {
	return Success
}
func (n *SuccessBehaviour) Teardown() {}
func (n *SuccessBehaviour) Abort()    {}

// FailureBehaviour
type FailureBehaviour struct {
	node
}

func (n *FailureBehaviour) Initialize() {}
func (n *FailureBehaviour) Update() Status {
	return Failure
}
func (n *FailureBehaviour) Teardown() {}
func (n *FailureBehaviour) Abort()    {}

// RunningBehaviour
type RunningBehaviour struct {
	node
}

func (n *RunningBehaviour) Initialize() {}
func (n *RunningBehaviour) Update() Status {
	return Running
}
func (n *RunningBehaviour) Teardown() {}
func (n *RunningBehaviour) Abort()    {}

// XThenY
type XThenY struct {
	node

	accessed bool
	X, Y     Status
}

func (n *XThenY) Initialize() {}
func (n *XThenY) Update() Status {
	if n.accessed {
		return n.Y
	}
	n.accessed = true
	return n.X
}
func (n *XThenY) Teardown() {}
func (n *XThenY) Abort()    {}
