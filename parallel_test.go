package goedbt_test

import (
	"testing"

	goedbt "github.com/dpsommer/go-edbt"
)

func TestParallel(t *testing.T) {
	tree := goedbt.NewBehaviourTree()
	parallel := &goedbt.ParallelNode{}
	parallel.AddChild(&goedbt.FailureTask{})
	parallel.AddChild(&goedbt.SuccessTask{})
	tree.Root.AddChild(parallel)
	status := tree.Root.Tick()

	if status != goedbt.Success {
		t.Errorf("ParallelNode got %d, want %d", status, goedbt.Success)
	}
}
