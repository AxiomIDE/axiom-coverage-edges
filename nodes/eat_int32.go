package nodes

import (
	"context"

	"axiom-official/axiom-coverage-edges/axiom"
	gen "axiom-official/axiom-coverage-edges/gen"
)

// EatInt32 adds 1 to its Int32Msg input. Sink side for int64→int32 narrowing
// edge tests — the reverse direction of the wave-001 EmitInt32→EatInt64
// widen pair. The +1 lets a future assertion distinguish "Eat ran" from
// "Eat received the input verbatim". Coverage FR-010.
func EatInt32(ctx context.Context, ax axiom.Context, input *gen.Int32Msg) (*gen.Int32Msg, error) {
	_ = ctx
	ax.Log().Info("EatInt32", "value", input.GetValue())
	return &gen.Int32Msg{Value: input.GetValue() + 1}, nil
}
