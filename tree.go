package goedbt

type Search func(*Event) bool

type BehaviourTree struct {
	blackboard map[string]any
	scheduler  *deque[*Event]
}

func NewBehaviourTree() *BehaviourTree {
	return &BehaviourTree{
		blackboard: map[string]any{},
		scheduler:  &deque[*Event]{},
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
	e := t.scheduler.PopFront()
	if e == nil {
		return false
	}

	state := tick(e)
	if state != Running && e.Observer != nil {
		e.Observer(state)
	} else {
		t.scheduler.PushBack(e)
	}

	return true
}

func (t *BehaviourTree) Start(b Behaviour, o Observer) {
	t.scheduler.PushFront(&Event{Behaviour: b, Observer: o})
}

func (t *BehaviourTree) Stop(e *Event, s Status) {
	if s == Running {
		panic("can't set Running state for stopped behaviour")
	}

	e.SetState(s)
	if e.Observer != nil {
		e.Observer(s)
	}
}

func (t *BehaviourTree) Abort(b Behaviour, s Search) {
	b.abort()
	t.scheduler.Remove(t.scheduler.RIndex(s))
}
