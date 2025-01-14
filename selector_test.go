package goedbt_test

import (
	"testing"

	goedbt "github.com/dpsommer/go-edbt"
)

func setupSimpleSelectorTree(tasks ...goedbt.Behaviour) *goedbt.BehaviourTree {
	selector := goedbt.NewSelector()
	for _, t := range tasks {
		selector.AddChild(t)
	}

	tree := goedbt.NewBehaviourTree(selector)

	return tree
}

func TestSelector(t *testing.T) {
	tree := setupSimpleSelectorTree(
		&goedbt.FailureBehaviour{},
		&goedbt.SuccessBehaviour{},
	)
	status := goedbt.Tick(tree.Root)

	if status != goedbt.Success {
		t.Errorf("Selector got %d, want %d", status, goedbt.Success)
	}
}

func TestSelectorRunning(t *testing.T) {
	tree := setupSimpleSelectorTree(
		&goedbt.FailureBehaviour{},
		&goedbt.RunningBehaviour{},
	)
	status := goedbt.Tick(tree.Root)

	if status != goedbt.Running {
		t.Errorf("Selector got %d, want %d", status, goedbt.Running)
	}
}
