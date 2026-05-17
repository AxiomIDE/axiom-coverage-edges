package nodes_test

import (
	"context"
	"testing"

	gen "axiom-official/axiom-coverage-edges/gen"
	"axiom-official/axiom-coverage-edges/nodes"
)

func TestEatLeaf(t *testing.T) {
	got, err := nodes.EatLeaf(context.Background(), newTestContext(t), &gen.LeafMsg{Leaf: "x"})
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if got.GetLeaf() != "leaf:x" {
		t.Errorf("EatLeaf should prepend 'leaf:'; got %q want %q", got.GetLeaf(), "leaf:x")
	}
}
