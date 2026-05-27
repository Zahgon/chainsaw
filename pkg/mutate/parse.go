package mutate

import (
	"context"

	"github.com/kyverno/chainsaw/pkg/apis"
	"github.com/kyverno/kyverno-json/pkg/core/compilers"
	"k8s.io/apimachinery/pkg/util/validation/field"
)

func Parse(ctx context.Context, mutation any) Mutation {
	_ = "STUB: not implemented"
	return *new(Mutation)
}

// mapNode is the mutation type represented by a map.
// it is responsible for projecting the analysed resource and passing the result to the descendant
type mapNode map[any]Mutation

func (n mapNode) mutate(ctx context.Context, c compilers.Compilers, path *field.Path, value any, bindings apis.Bindings) (any, error) {
	_ = "STUB: not implemented"
	return *new(any), nil
}

// sliceNode is the mutation type represented by a slice.
// it first compares the length of the analysed resource with the length of the descendants.
// if lengths match all descendants are evaluated with their corresponding items.
type sliceNode []Mutation

func (n sliceNode) mutate(ctx context.Context, c compilers.Compilers, path *field.Path, value any, bindings apis.Bindings) (any, error) {
	_ = "STUB: not implemented"
	return *new(any), nil
}

// scalarNode is a terminal type of mutation.
// it receives a value and compares it with an expected value.
// the expected value can be the result of an expression.
type scalarNode struct {
	rhs any
}

func (n *scalarNode) mutate(ctx context.Context, c compilers.Compilers, path *field.Path, value any, bindings apis.Bindings) (any, error) {
	_ = "STUB: not implemented"
	return *new(any), nil
}

// we only project if the expression uses the engine syntax
// this is to avoid the case where the value is a map and the RHS is a string
