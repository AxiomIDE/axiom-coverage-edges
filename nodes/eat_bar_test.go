package nodes_test

import (
	"context"
	"testing"

	gen "axiom-official/axiom-coverage-edges/gen"
	"axiom-official/axiom-coverage-edges/nodes"
)

func TestEatBar(t *testing.T) {
	got, err := nodes.EatBar(context.Background(), newTestContext(t), &gen.BarMsg{Bar: "hi"})
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if got.GetBar() != "eaten:hi" {
		t.Errorf("EatBar should prepend 'eaten:'; got %q want %q", got.GetBar(), "eaten:hi")
	}
}
