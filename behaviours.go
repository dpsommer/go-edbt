package goedbt

// SuccessBehaviour
type SuccessBehaviour struct {
	node
}

func (n *SuccessBehaviour) Initialize() {}
func (n *SuccessBehaviour) Update() Status {
	return Success
}
func (n *SuccessBehaviour) Terminate() {}

// FailureBehaviour
type FailureBehaviour struct {
	node
}

func (n *FailureBehaviour) Initialize() {}
func (n *FailureBehaviour) Update() Status {
	return Failure
}
func (n *FailureBehaviour) Terminate() {}

// RunningBehaviour
type RunningBehaviour struct {
	node
}

func (n *RunningBehaviour) Initialize() {}
func (n *RunningBehaviour) Update() Status {
	return Running
}
func (n *RunningBehaviour) Terminate() {}
