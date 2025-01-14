package goedbt

type BehaviourTree struct {
	blackboard map[string]any
	Root       Behaviour
}

func NewBehaviourTree(b Behaviour) *BehaviourTree {
	return &BehaviourTree{
		blackboard: map[string]any{},
		Root:       b,
	}
}

func (t *BehaviourTree) Start() {
	for {
		// XXX: just ignore the status for now
		tick(t.Root)
	}
}
