package goedbt_test

import (
	"testing"

	goedbt "github.com/dpsommer/go-edbt"
)

func TestSelector(t *testing.T) {
	tree := goedbt.NewBehaviourTree()
	selector := &goedbt.SelectorNode{}
	selector.AddChild(&goedbt.FailureTask{})
	selector.AddChild(&goedbt.SuccessTask{})
	tree.Root.AddChild(selector)
	status := tree.Root.Tick()

	if status != goedbt.Success {
		t.Errorf("SelectorNode got %d, want %d", status, goedbt.Success)
	}
}
