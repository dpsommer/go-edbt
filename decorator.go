package goedbt

type Decorator interface {
	Behaviour

	// ...
}

type decorator struct {
	*behaviour

	tree  *BehaviourTree
	child Behaviour
}
