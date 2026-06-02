package nodes_test

import (
	"context"
	"testing"

	"axiom-official/axiom-coverage-edges/axiom"
)

type testContext struct {
	t          *testing.T
	secretsMap map[string]string
}

func newTestContext(t *testing.T) *testContext {
	return &testContext{t: t, secretsMap: map[string]string{}}
}

type testLogger struct{ t *testing.T }

func (l *testLogger) Debug(msg string, args ...any) { l.t.Logf("DEBUG  %s %v", msg, args) }
func (l *testLogger) Info(msg string, args ...any)  { l.t.Logf("INFO   %s %v", msg, args) }
func (l *testLogger) Warn(msg string, args ...any)  { l.t.Logf("WARN   %s %v", msg, args) }
func (l *testLogger) Error(msg string, args ...any) { l.t.Logf("ERROR  %s %v", msg, args) }

type testSecrets struct{ m map[string]string }

func (s testSecrets) Get(name string) (string, bool) { v, ok := s.m[name]; return v, ok }

type testAgent struct{}

func (testAgent) Memory() axiom.AgentMemory { return testAgentMemory{} }

type testAgentMemory struct{}

func (testAgentMemory) Session(_ string) axiom.SessionMemory { return testSessionMemory{} }
func (testAgentMemory) Search(_ context.Context, _ string, _ int) ([]axiom.MemoryEntry, error) {
	return nil, nil
}
func (testAgentMemory) Write(_ context.Context, _ string, _ float32) (string, error) {
	return "", nil
}

type testSessionMemory struct{}

func (testSessionMemory) Search(_ context.Context, _ string, _ int) ([]axiom.MemoryEntry, error) {
	return nil, nil
}
func (testSessionMemory) Write(_ context.Context, _ string, _ float32) (string, error) {
	return "", nil
}
func (testSessionMemory) History() axiom.SessionHistory { return testSessionHistory{} }
func (testSessionMemory) End(_ context.Context) error   { return nil }

type testSessionHistory struct{}

func (testSessionHistory) Last(_ context.Context, _ int) ([]axiom.ConversationTurn, error) {
	return nil, nil
}
func (testSessionHistory) Append(_ context.Context, _, _ string) error { return nil }

func (c *testContext) Log() axiom.Logger            { return &testLogger{c.t} }
func (c *testContext) Secrets() axiom.Secrets       { return testSecrets{c.secretsMap} }
func (c *testContext) Agent() axiom.Agent           { return testAgent{} }
func (c *testContext) ExecutionID() string          { return "test-execution-id" }
func (c *testContext) FlowID() string               { return "test-flow-id" }
func (c *testContext) TenantID() string             { return "test-tenant-id" }
func (c *testContext) Reflection() axiom.Reflection { return emptyReflection{} }

// ADR-051 (2026-05-26): no-op Mutation stub. axiom push regenerates the
// fixture's axiom/ package; once SAF-001 lands axiom.Context requires
// Mutation() so the testContext interface assertion fails without it.
// Coverage-edges fixtures don't exercise the mutation path.
func (c *testContext) Mutation() axiom.Mutation { return testMutation{} }

type testMutation struct{}

func (testMutation) Flow() axiom.FlowMutation { return testFlowMutation{} }

type testFlowMutation struct{}

func (testFlowMutation) AddNode(_, _ string, _ *axiom.CanvasPosition) uint32 { return 0 }
func (testFlowMutation) AddEdge(_, _ uint32, _ *axiom.EdgeCondition)         {}

var _ axiom.Context = (*testContext)(nil)

// ADR-050 (2026-05-26): the axiom.Context interface added Reflection() so
// every node SDK has access to ax.Reflection().Flow().*. None of the
// fixture nodes here use reflection (they're pure type-converters for edge
// tests), so the test injects an empty reflection rather than building a
// populated FlowReflection. axiom build regenerates axiom/context.go with
// the Reflection types referenced below.

type emptyReflection struct{}

func (emptyReflection) Flow() axiom.FlowReflection { return emptyFlowReflection{} }

type emptyFlowReflection struct{}

func (emptyFlowReflection) Nodes() []axiom.ReflectionNode     { return nil }
func (emptyFlowReflection) Edges() []axiom.ReflectionEdge     { return nil }
func (emptyFlowReflection) LoopEdges() []axiom.ReflectionEdge { return nil }
func (emptyFlowReflection) Position() axiom.FlowPosition      { return axiom.FlowPosition{} }
func (emptyFlowReflection) GraphID() string                   { return "" }
