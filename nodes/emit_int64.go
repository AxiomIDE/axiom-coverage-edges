package nodes

import (
	"context"

	"axiom-official/axiom-coverage-edges/axiom"
	gen "axiom-official/axiom-coverage-edges/gen"
)

// EmitInt64 passes its Int64Msg input through unchanged. Source side for
// int64→int32 narrowing edge tests — the reverse direction of the wave-001
// EmitInt32→EatInt64 widen pair. Coverage FR-010.
func EmitInt64(ctx context.Context, ax axiom.Context, input *gen.Int64Msg) (*gen.Int64Msg, error) {
	_ = ctx
	ax.Log().Info("EmitInt64", "value", input.GetValue())
	return &gen.Int64Msg{Value: input.GetValue()}, nil
}
