package goedbt

import (
	"context"
)

// ParallelNode defines a composite node that ticks its children
// concurrently. Returns Running until at least one child succeeds.
// If all children fail, returns Failure.
type ParallelNode struct {
	children []Node
	// TODO: define a duration to re-tick children when Running
}

func (n *ParallelNode) Tick() Status {
	status := Failure
	results := make(chan Status, 1)

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	for _, c := range n.children {
		go func(ctx context.Context, c Node) {
			select {
			case results <- c.Tick():
			case <-ctx.Done():
			}
		}(ctx, c)
	}

	for range len(n.children) {
		res := <-results
		if res == Success {
			return Success
		} else if res == Running {
			status = Running
		}
	}

	return status
}

func (n *ParallelNode) Children() []Node {
	return n.children
}

func (n *ParallelNode) AddChild(child Node) error {
	n.children = append(n.children, child)
	return nil
}
