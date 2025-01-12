package goedbt

type BehaviourTree struct {
	blackboard map[string]any
	root       Node
}

func NewBehaviourTree() *BehaviourTree {
	return &BehaviourTree{
		blackboard: map[string]any{},
		root:       &RootNode{},
	}
}

func (t *BehaviourTree) Start() {
	for {
		// XXX: just ignore the status for now
		t.root.Tick()
	}
}
