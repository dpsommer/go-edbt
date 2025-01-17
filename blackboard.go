package goedbt

type blackboard struct {
	values    map[string]any
	observers map[string]*BlackboardObserver
}

func newBlackboard() *blackboard {
	return &blackboard{
		values:    map[string]any{},
		observers: map[string]*BlackboardObserver{},
	}
}
