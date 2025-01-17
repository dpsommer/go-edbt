package goedbt

type Policy int

const (
	RequireOne Policy = iota
	RequireAll
)

// Parallel defines a composite Behaviour that ticks all of its children
// in parallel each tick. Accepts a successPolicy of either RequireOne or
// RequireAll.
//
// With RequireOne, returns Success on the first Success from a child. If
// no child succeeds, continues to tick any Running children and returns
// Failure once all children complete unsuccessfully.
//
// With RequireAll, returns Failure if a non-Success status is received from
// a completed child. Continues to tick any Running children if all completed
// children have been successful, returning Success if all children succeed.
type Parallel struct {
	*composite

	// only specify a successPolicy - this implicitly defines the failure
	// policy and avoids ambiguity
	successPolicy Policy
	successes     int
	failures      int
}

func NewParallel(bt *BehaviourTree, successPolicy Policy) *Parallel {
	return &Parallel{
		successPolicy: successPolicy,
		composite: &composite{
			tree:      bt,
			behaviour: &behaviour{state: Invalid},
			children:  make(Set[Behaviour]),
		},
	}
}

func (n *Parallel) initialize() {
	n.successes = 0
	n.failures = 0
	for c := range n.children {
		n.tree.Start(c, n.onChildComplete)
	}
}

func (n *Parallel) onChildComplete(s Status) {
	switch s {
	case Success:
		n.successes += 1
		if n.successPolicy == RequireOne || n.successes == len(n.children) {
			n.tree.Stop(&Event{n, nil}, s)
		}
	default:
		n.failures += 1
		if n.successPolicy == RequireAll || n.failures == len(n.children) {
			n.tree.Stop(&Event{n, nil}, Failure)
		}
	}
	// teardown and abort any running children if we
	// stop before all children have finished
	if n.state != Running && n.successes+n.failures < len(n.children) {
		n.teardown()
	}
}

func (n *Parallel) update() Status {
	return Running
}

func (n *Parallel) teardown() {
	for c := range n.children {
		if c.State() == Running {
			n.tree.Abort(c, func(e *Event) bool {
				return e.Behaviour == c
			})
		}
	}
}

func (n *Parallel) abort() {}
