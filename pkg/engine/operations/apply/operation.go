package apply

import (
	"context"

	"github.com/kyverno/chainsaw/pkg/apis"
	"github.com/kyverno/chainsaw/pkg/apis/v1alpha1"
	"github.com/kyverno/chainsaw/pkg/cleanup/cleaner"
	"github.com/kyverno/chainsaw/pkg/client"
	"github.com/kyverno/chainsaw/pkg/engine/namespacer"
	"github.com/kyverno/chainsaw/pkg/engine/operations"
	"github.com/kyverno/chainsaw/pkg/engine/outputs"
	"github.com/kyverno/kyverno-json/pkg/core/compilers"
	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
)

type operation struct {
	compilers  compilers.Compilers
	client     client.Client
	base       unstructured.Unstructured
	namespacer namespacer.Namespacer
	cleaner    cleaner.CleanerCollector
	template   bool
	expect     []v1alpha1.Expectation
	outputs    []v1alpha1.Output
}

func New(
	compilers compilers.Compilers,
	client client.Client,
	obj unstructured.Unstructured,
	namespacer namespacer.Namespacer,
	cleaner cleaner.CleanerCollector,
	template bool,
	expect []v1alpha1.Expectation,
	outputs []v1alpha1.Output,
) operations.Operation {
	_ = "STUB: not implemented"
	return *new(operations.Operation)
}

func (o *operation) Exec(ctx context.Context, tc apis.Bindings) (_ outputs.Outputs, _err error) {
	_ = "STUB: not implemented"
	return *new(outputs.Outputs), nil
}

func (o *operation) execute(ctx context.Context, tc apis.Bindings, obj unstructured.Unstructured) (outputs.Outputs, error) {
	_ = "STUB: not implemented"
	return *new(outputs.Outputs), nil
}

// TODO: determine if the error can be retried

func (o *operation) tryApplyResource(ctx context.Context, tc apis.Bindings, obj unstructured.Unstructured) (outputs.Outputs, error) {
	_ = "STUB: not implemented"
	return *new(outputs.Outputs), nil
}

func (o *operation) updateResource(ctx context.Context, tc apis.Bindings, actual *unstructured.Unstructured, obj unstructured.Unstructured) (outputs.Outputs, error) {
	_ = "STUB: not implemented"
	return *new(outputs.Outputs), nil
}

func (o *operation) createResource(ctx context.Context, tc apis.Bindings, obj unstructured.Unstructured) (outputs.Outputs, error) {
	_ = "STUB: not implemented"
	return *new(outputs.Outputs), nil
}

func (o *operation) handleCheck(ctx context.Context, tc apis.Bindings, obj unstructured.Unstructured, err error) (_outputs outputs.Outputs, _err error) {
	_ = "STUB: not implemented"
	return *new(outputs.Outputs), nil
}
