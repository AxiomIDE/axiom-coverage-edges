package nodes_test

import (
	"context"
	"testing"

	gen "axiom-official/axiom-coverage-edges/gen"
	"axiom-official/axiom-coverage-edges/nodes"
)

func TestEmitNested(t *testing.T) {
	got, err := nodes.EmitNested(context.Background(), newTestContext(t), &gen.NestedMsg{Inner: &gen.Inner{Value: "hello"}})
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if got.GetInner().GetValue() != "hello" {
		t.Errorf("EmitNested should pass inner.value through; got %q want %q", got.GetInner().GetValue(), "hello")
	}
}

func TestEmitNestedNilInner(t *testing.T) {
	got, err := nodes.EmitNested(context.Background(), newTestContext(t), &gen.NestedMsg{})
	if err != nil {
		t.Fatalf("unexpected error on nil inner: %v", err)
	}
	if got.GetInner().GetValue() != "" {
		t.Errorf("EmitNested with nil inner should produce zero-value inner; got %q", got.GetInner().GetValue())
	}
}
