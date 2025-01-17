package goedbt

type condition func(bod *BlackboardObserver) bool

func HasValue(bod *BlackboardObserver) bool {
	v, ok := bod.tree.blackboard.values[bod.key]
	return ok && v != nil
}

// TODO: how should abort rules be implemented?
// they need to be evaluated and abort neighbouring
// nodes (in the example case of lower-priority) if
// the BOD condition is true. does that mean they
// need to know about their parent?
type abortRule func()

// TODO: the BOD should implement an observer decorator intf
// ... or embed an observerDecorator struct?
type ObserverDecorator interface{}

type BlackboardObserver struct {
	*decorator

	key string
	condition
	abortRule
}

type BODOptions struct {
	key   string
	child Behaviour

	condition
	abortRule
}

func NewBlackboardObserver(bt *BehaviourTree, opts BODOptions) *BlackboardObserver {
	return &BlackboardObserver{
		key:       opts.key,
		condition: opts.condition,
		abortRule: opts.abortRule,
		decorator: &decorator{
			tree:      bt,
			behaviour: &behaviour{state: Invalid},
			child:     opts.child,
		},
	}
}

func (d *BlackboardObserver) initialize() {

}

func (d *BlackboardObserver) update() Status {
	if d.condition(d) {
		d.abortRule()

		d.tree.Start(d.child, d.onChildComplete)
		return Running
	}

	return Failure
}

func (d *BlackboardObserver) onChildComplete(s Status) {
	if s == Success {
		d.state = Success
	} else {
		d.state = Failure
	}
}

func (d *BlackboardObserver) teardown() {}

func (d *BlackboardObserver) abort() {}
