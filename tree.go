package goedbt

type BehaviourTree struct {
	blackboard map[string]any
	scheduler  *Deque[Behaviour]
}

func NewBehaviourTree() *BehaviourTree {
	return &BehaviourTree{
		blackboard: map[string]any{},
		scheduler:  NewDeque[Behaviour](minCapacity),
	}
}

func (t *BehaviourTree) Tick() {
	// demarcate the end of the current tick schedule with a nil value
	t.scheduler.PushBack(nil)
	for t.step() {
		// continue until we reach the nil value
	}
}

func (t *BehaviourTree) step() bool {
	b := t.scheduler.PopFront()
	if b == nil {
		return false
	}

	state := tick(b)
	if state != Running {
		b.FireObserver(state)
	} else {
		t.scheduler.PushBack(b)
	}

	return true
}

func (t *BehaviourTree) Start(b Behaviour) {
	t.scheduler.PushFront(b)
}

func (t *BehaviourTree) Stop(b Behaviour, s Status) {
	if s == Running {
		panic("can't set Running state for stopped behaviour")
	}

	b.SetState(s)
	b.FireObserver(s)
}
