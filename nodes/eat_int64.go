package nodes

import (
	"context"

	"axiom-official/axiom-coverage-edges/axiom"
	gen "axiom-official/axiom-coverage-edges/gen"
)

// EatInt64 adds 1 to its Int64Msg input. Sink side for int32→int64 widening
// edge tests — the +1 lets a future assertion distinguish "Eat ran" from
// "Eat received the input verbatim".
func EatInt64(ctx context.Context, ax axiom.Context, input *gen.Int64Msg) (*gen.Int64Msg, error) {
	_ = ctx
	ax.Log().Info("EatInt64", "value", input.GetValue())
	return &gen.Int64Msg{Value: input.GetValue() + 1}, nil
}
