package nodes_test

import (
	"context"
	"testing"

	gen "axiom-official/axiom-coverage-edges/gen"
	"axiom-official/axiom-coverage-edges/nodes"
)

func TestEatInt32(t *testing.T) {
	got, err := nodes.EatInt32(context.Background(), newTestContext(t), &gen.Int32Msg{Value: 100})
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if got.GetValue() != 101 {
		t.Errorf("EatInt32 should add 1; got %d want %d", got.GetValue(), 101)
	}
}
