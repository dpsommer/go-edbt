package goedbt

type BehaviourTree struct {
	blackboard map[string]any
	Root       Node
}

func NewBehaviourTree() *BehaviourTree {
	return &BehaviourTree{
		blackboard: map[string]any{},
		Root:       &RootNode{},
	}
}

func (t *BehaviourTree) Start() {
	for {
		// XXX: just ignore the status for now
		t.Root.Tick()
	}
}
