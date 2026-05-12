package nodes_test

import (
	"context"
	"testing"

	gen "axiom-official/axiom-coverage-edges/gen"
	"axiom-official/axiom-coverage-edges/nodes"
)

func TestEmitFoo(t *testing.T) {
	got, err := nodes.EmitFoo(context.Background(), newTestContext(t), &gen.FooMsg{Foo: "hello-foo"})
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if got.GetFoo() != "hello-foo" {
		t.Errorf("EmitFoo should pass FooMsg.foo through; got %q want %q", got.GetFoo(), "hello-foo")
	}
}
