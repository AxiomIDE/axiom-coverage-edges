package nodes_test

import (
	"context"
	"testing"

	gen "axiom-official/axiom-coverage-edges/gen"
	"axiom-official/axiom-coverage-edges/nodes"
)

func TestEmitInt32(t *testing.T) {
	got, err := nodes.EmitInt32(context.Background(), newTestContext(t), &gen.Int32Msg{Value: 42})
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if got.GetValue() != 42 {
		t.Errorf("EmitInt32 should pass Int32Msg.value through; got %d want %d", got.GetValue(), 42)
	}
}
