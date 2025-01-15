package goedbt_test

import (
	"testing"

	goedbt "github.com/dpsommer/go-edbt"
)

func TestSelector(t *testing.T) {
	tt := map[string]struct {
		behaviours []goedbt.Behaviour
		expected   []goedbt.Status
	}{
		"returns success": {
			behaviours: []goedbt.Behaviour{
				&goedbt.FailureBehaviour{},
				&goedbt.SuccessBehaviour{},
			},
			expected: []goedbt.Status{goedbt.Success},
		},
		"returns running": {
			behaviours: []goedbt.Behaviour{
				&goedbt.FailureBehaviour{},
				&goedbt.RunningBehaviour{},
			},
			expected: []goedbt.Status{goedbt.Running},
		},
		"returns running then success": {
			behaviours: []goedbt.Behaviour{
				&goedbt.FailureBehaviour{},
				&goedbt.XThenY{X: goedbt.Running, Y: goedbt.Success},
			},
			expected: []goedbt.Status{goedbt.Running, goedbt.Success},
		},
	}

	for name, tc := range tt {
		t.Run(name, func(t *testing.T) {
			tree := setupCompositeTree(goedbt.NewSelector(), tc.behaviours...)

			for _, s := range tc.expected {
				status := goedbt.Tick(tree.Root)
				if status != s {
					t.Errorf("Selector got %d, want %d", status, s)
				}
			}
		})
	}
}
