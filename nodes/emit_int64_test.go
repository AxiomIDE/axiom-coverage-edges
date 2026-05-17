package nodes_test

import (
	"context"
	"testing"

	gen "axiom-official/axiom-coverage-edges/gen"
	"axiom-official/axiom-coverage-edges/nodes"
)

func TestEmitInt64(t *testing.T) {
	got, err := nodes.EmitInt64(context.Background(), newTestContext(t), &gen.Int64Msg{Value: 7})
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if got.GetValue() != 7 {
		t.Errorf("EmitInt64 should pass Int64Msg.value through; got %d want %d", got.GetValue(), 7)
	}
}
