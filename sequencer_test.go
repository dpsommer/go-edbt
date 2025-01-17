package goedbt_test

import (
	"testing"

	goedbt "github.com/dpsommer/go-edbt"
)

func setupCompositeTree(t *goedbt.BehaviourTree, c goedbt.Composite, bb ...goedbt.Behaviour) {
	for _, b := range bb {
		c.AddChild(b)
	}

	t.Start(c, nil)
}

func TestSequencer(t *testing.T) {
	tt := map[string]struct {
		behaviours []goedbt.Behaviour
		expected   []goedbt.Status
	}{
		"returns success then failure": {
			behaviours: []goedbt.Behaviour{
				&goedbt.XThenY{X: goedbt.Success, Y: goedbt.Failure},
			},
			expected: []goedbt.Status{goedbt.Success, goedbt.Failure},
		},
		"returns running": {
			behaviours: []goedbt.Behaviour{
				&goedbt.SuccessBehaviour{},
				&goedbt.RunningBehaviour{},
			},
			expected: []goedbt.Status{goedbt.Running},
		},
		"returns running then success": {
			behaviours: []goedbt.Behaviour{
				&goedbt.SuccessBehaviour{},
				&goedbt.XThenY{X: goedbt.Running, Y: goedbt.Success},
			},
			expected: []goedbt.Status{goedbt.Running, goedbt.Success},
		},
	}

	for name, tc := range tt {
		t.Run(name, func(t *testing.T) {
			tree := goedbt.NewBehaviourTree()
			sequencer := goedbt.NewSequencer(tree)
			setupCompositeTree(tree, sequencer, tc.behaviours...)

			for _, s := range tc.expected {
				tree.Tick()
				if sequencer.State() != s {
					t.Errorf("Selector got %d, want %d", sequencer.State(), s)
				}
			}
		})
	}
}
