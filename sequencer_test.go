package goedbt_test

import (
	"testing"

	goedbt "github.com/dpsommer/go-edbt"
)

func TestSequencer(t *testing.T) {
	sequencer := goedbt.NewSequencer()
	sequencer.AddChild(&goedbt.SuccessBehaviour{})
	tree := goedbt.NewBehaviourTree(sequencer)

	status := goedbt.Tick(tree.Root)
	if status != goedbt.Success {
		t.Errorf("Sequencer got %d, want %d", status, goedbt.Success)
	}

	sequencer.AddChild(&goedbt.FailureBehaviour{})
	status = goedbt.Tick(tree.Root)
	if status != goedbt.Failure {
		t.Errorf("Sequencer got %d, want %d", status, goedbt.Failure)
	}
}
