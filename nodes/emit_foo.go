package nodes

import (
	"context"

	"axiom-official/axiom-coverage-edges/axiom"
	gen "axiom-official/axiom-coverage-edges/gen"
)

// EmitFoo passes its FooMsg input through unchanged. Used as the source side
// of an edge under test so the receiver sees a known FooMsg.foo value.
func EmitFoo(ctx context.Context, ax axiom.Context, input *gen.FooMsg) (*gen.FooMsg, error) {
	_ = ctx
	ax.Log().Info("EmitFoo", "foo", input.GetFoo())
	return &gen.FooMsg{Foo: input.GetFoo()}, nil
}
