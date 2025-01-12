package goedbt_test

import (
	"testing"

	goedbt "github.com/dpsommer/go-edbt"
)

func TestSequencer(t *testing.T) {
	tree := goedbt.NewBehaviourTree()
	sequencer := &goedbt.SequencerNode{}
	tree.Root.AddChild(sequencer)
	sequencer.AddChild(&goedbt.SuccessTask{})
	status := tree.Root.Tick()
	if status != goedbt.Success {
		t.Errorf("SequencerNode got %d, want %d", status, goedbt.Success)
	}
	sequencer.AddChild(&goedbt.FailureTask{})
	status = tree.Root.Tick()
	if status != goedbt.Failure {
		t.Errorf("SequencerNode got %d, want %d", status, goedbt.Failure)
	}
}
