package goedbt

type Policy int

const (
	RequireOne Policy = iota
	RequireAll
)

// Parallel defines a composite Behaviour that ticks its children
// concurrently. Returns Running until at least one child succeeds.
// If all children fail, returns Failure.
type Parallel struct {
	*composite

	// only specify a successPolicy - this implicitly defines the failure
	// policy and avoids ambiguity (e.g. if successPolicy and failurePolicy
	// are both RequireAll)
	successPolicy Policy

	runningNodes []Behaviour
}

func NewParallel(successPolicy Policy) *Parallel {
	return &Parallel{
		successPolicy: successPolicy,
		composite: &composite{
			node:     &node{state: Invalid},
			children: make(Set[Behaviour]),
		},
	}
}

func (n *Parallel) Initialize() {
	n.runningNodes = keys(n.children)
}

func (n *Parallel) Update() Status {
	var b Behaviour
	// track success count and required number of successes for RequireAll
	// we only need len(runningNodes) successes each tick as we can imply that:
	// * if RequireOne, we haven't succeeded yet but any Success will be enough
	// * if RequireAll, completed tasks must have been Success as we fail fast
	var successes int
	needSuccesses := len(n.runningNodes)

	// default to Failure so we don't need to count Failures for RequireOne
	status := Failure
	// list all still-running tasks; check only that subset on subsequent ticks
	stillRunning := make([]Behaviour, needSuccesses)

	for len(n.runningNodes) > 0 {
		b, n.runningNodes = pop(n.runningNodes)
		res := tick(b)

		if res == Success {
			successes += 1
			if n.successPolicy == RequireOne || successes == needSuccesses {
				n.state = Success
				return n.state
			}
		} else if res == Running {
			status = Running
			stillRunning = append(stillRunning, b)
		} else if n.successPolicy == RequireAll {
			// fail fast if a child task failed or was aborted
			n.state = res
			return n.state
		}
	}

	if status == Running {
		n.runningNodes = stillRunning
	}

	n.state = status
	return n.state
}

func (n *Parallel) Teardown() {
	// if this node is configured to Teardown early,
	// abort any running children
	for c := range n.children {
		c.Abort()
	}
}

func (n *Parallel) Abort() {}
