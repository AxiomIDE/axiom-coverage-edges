package nodes

import (
	"context"

	"axiom-official/axiom-coverage-edges/axiom"
	gen "axiom-official/axiom-coverage-edges/gen"
)

// EatLeaf prepends "leaf:" to its LeafMsg input. Sink side for the
// edges-conversion/nested-extract variant — receives the projection of
// NestedMsg.inner.value via the edge adapter, so the prefix lets a future
// assertion distinguish "Eat ran" from "Eat received the input verbatim".
// Coverage FR-011.
func EatLeaf(ctx context.Context, ax axiom.Context, input *gen.LeafMsg) (*gen.LeafMsg, error) {
	_ = ctx
	ax.Log().Info("EatLeaf", "leaf", input.GetLeaf())
	return &gen.LeafMsg{Leaf: "leaf:" + input.GetLeaf()}, nil
}
