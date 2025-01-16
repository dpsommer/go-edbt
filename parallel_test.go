package goedbt_test

import (
	"testing"

	goedbt "github.com/dpsommer/go-edbt"
)

func TestParallel(t *testing.T) {
	tt := map[string]struct {
		successPolicy goedbt.Policy
		behaviours    []goedbt.Behaviour
		expected      []goedbt.Status
	}{
		"RequireOne success": {
			successPolicy: goedbt.RequireOne,
			behaviours: []goedbt.Behaviour{
				&goedbt.FailureBehaviour{},
				&goedbt.SuccessBehaviour{},
			},
			expected: []goedbt.Status{goedbt.Success},
		},
		"RequireOne failure": {
			successPolicy: goedbt.RequireOne,
			behaviours: []goedbt.Behaviour{
				&goedbt.FailureBehaviour{},
				&goedbt.FailureBehaviour{},
			},
			expected: []goedbt.Status{goedbt.Failure},
		},
		"RequireAll success": {
			successPolicy: goedbt.RequireAll,
			behaviours: []goedbt.Behaviour{
				&goedbt.SuccessBehaviour{},
				&goedbt.SuccessBehaviour{},
			},
			expected: []goedbt.Status{goedbt.Success},
		},
		"RequireAll failure": {
			successPolicy: goedbt.RequireAll,
			behaviours: []goedbt.Behaviour{
				&goedbt.SuccessBehaviour{},
				&goedbt.SuccessBehaviour{},
				&goedbt.FailureBehaviour{},
			},
			expected: []goedbt.Status{goedbt.Failure},
		},
		"returns running": {
			successPolicy: goedbt.RequireOne,
			behaviours: []goedbt.Behaviour{
				&goedbt.RunningBehaviour{},
				&goedbt.FailureBehaviour{},
				&goedbt.FailureBehaviour{},
			},
			expected: []goedbt.Status{goedbt.Running},
		},
		"returns success despite running child": {
			successPolicy: goedbt.RequireOne,
			behaviours: []goedbt.Behaviour{
				&goedbt.RunningBehaviour{},
				&goedbt.SuccessBehaviour{},
				&goedbt.FailureBehaviour{},
			},
			expected: []goedbt.Status{goedbt.Success},
		},
		"returns running then success": {
			successPolicy: goedbt.RequireOne,
			behaviours: []goedbt.Behaviour{
				&goedbt.FailureBehaviour{},
				&goedbt.XThenY{X: goedbt.Running, Y: goedbt.Success},
				&goedbt.FailureBehaviour{},
			},
			expected: []goedbt.Status{goedbt.Running, goedbt.Success},
		},
		"returns failure despite running child": {
			successPolicy: goedbt.RequireAll,
			behaviours: []goedbt.Behaviour{
				&goedbt.RunningBehaviour{},
				&goedbt.SuccessBehaviour{},
				&goedbt.FailureBehaviour{},
			},
			expected: []goedbt.Status{goedbt.Failure},
		},
		"returns running then failure": {
			successPolicy: goedbt.RequireAll,
			behaviours: []goedbt.Behaviour{
				&goedbt.SuccessBehaviour{},
				&goedbt.XThenY{X: goedbt.Running, Y: goedbt.Failure},
				&goedbt.SuccessBehaviour{},
			},
			expected: []goedbt.Status{goedbt.Running, goedbt.Failure},
		},
	}

	for name, tc := range tt {
		t.Run(name, func(t *testing.T) {
			tree := goedbt.NewBehaviourTree()
			parallel := goedbt.NewParallel(tree, tc.successPolicy)
			setupCompositeTree(tree, parallel, tc.behaviours...)

			for _, s := range tc.expected {
				tree.Tick()
				if parallel.State() != s {
					t.Errorf("Selector got %d, want %d", parallel.State(), s)
				}
			}
		})
	}
}
