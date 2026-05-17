package nodes

import (
	"context"

	"axiom-official/axiom-coverage-edges/axiom"
	gen "axiom-official/axiom-coverage-edges/gen"
)

// EmitNested passes its NestedMsg input through unchanged. Source side for
// the edges-conversion/nested-extract variant — adapter expressions project
// "inner.value" through the edge into LeafMsg.leaf. Coverage FR-011.
func EmitNested(ctx context.Context, ax axiom.Context, input *gen.NestedMsg) (*gen.NestedMsg, error) {
	_ = ctx
	inner := input.GetInner()
	if inner == nil {
		inner = &gen.Inner{}
	}
	ax.Log().Info("EmitNested", "inner.value", inner.GetValue())
	return &gen.NestedMsg{Inner: &gen.Inner{Value: inner.GetValue()}}, nil
}
