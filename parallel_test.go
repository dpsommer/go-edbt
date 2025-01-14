package goedbt_test

import (
	"testing"

	goedbt "github.com/dpsommer/go-edbt"
)

func TestParallel(t *testing.T) {
	parallel := goedbt.NewParallel()
	parallel.AddChild(&goedbt.FailureBehaviour{})
	parallel.AddChild(&goedbt.SuccessBehaviour{})

	tree := goedbt.NewBehaviourTree(parallel)

	status := goedbt.Tick(tree.Root)

	if status != goedbt.Success {
		t.Errorf("ParallelNode got %d, want %d", status, goedbt.Success)
	}
}
