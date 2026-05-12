package nodes

import (
	"context"

	"axiom-official/axiom-coverage-edges/axiom"
	gen "axiom-official/axiom-coverage-edges/gen"
)

// EmitInt32 passes its Int32Msg input through unchanged. Source side for
// int32→int64 widening edge tests.
func EmitInt32(ctx context.Context, ax axiom.Context, input *gen.Int32Msg) (*gen.Int32Msg, error) {
	_ = ctx
	ax.Log().Info("EmitInt32", "value", input.GetValue())
	return &gen.Int32Msg{Value: input.GetValue()}, nil
}
