package goedbt

import (
	"context"
)

// Parallel defines a composite Behaviour that ticks its children
// concurrently. Returns Running until at least one child succeeds.
// If all children fail, returns Failure.
type Parallel struct {
	*node
	*composite
}

func NewParallel() *Parallel {
	return &Parallel{
		node: &node{state: Invalid},
		composite: &composite{
			children: make(map[Behaviour]struct{}),
		},
	}
}

func (n *Parallel) Initialize() {}
func (n *Parallel) Terminate()  {}

func (n *Parallel) Update() Status {
	status := Failure
	results := make(chan Status, 1)

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	for c := range n.children {
		go func(ctx context.Context, c Behaviour) {
			select {
			case results <- tick(c):
			case <-ctx.Done():
			}
		}(ctx, c)
	}

	for range len(n.children) {
		res := <-results
		if res == Success {
			n.state = Success
			return n.state
		} else if res == Running {
			status = Running
		}
	}

	n.state = status
	return n.state
}
