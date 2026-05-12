package nodes

import (
	"context"

	"axiom-official/axiom-coverage-edges/axiom"
	gen "axiom-official/axiom-coverage-edges/gen"
)

// EatBar prepends "eaten:" to its BarMsg input. Used as the sink side of an
// edge under test so the conversion from FooMsg→BarMsg is observable.
func EatBar(ctx context.Context, ax axiom.Context, input *gen.BarMsg) (*gen.BarMsg, error) {
	_ = ctx
	ax.Log().Info("EatBar", "bar", input.GetBar())
	return &gen.BarMsg{Bar: "eaten:" + input.GetBar()}, nil
}
