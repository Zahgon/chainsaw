package create

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

func (o *operation) Exec(ctx context.Context, bindings apis.Bindings) (_ outputs.Outputs, _err error) {
	_ = "STUB: not implemented"
	return *new(outputs.Outputs), nil
}

func (o *operation) execute(ctx context.Context, bindings apis.Bindings, obj unstructured.Unstructured) (outputs.Outputs, error) {
	_ = "STUB: not implemented"
	return *new(outputs.Outputs), nil
}

// Check if the error is retryable

// Conflict errors should be retried

// Server timeout errors should be retried

// Too many requests errors should be retried

// Service unavailable errors should be retried

// AlreadyExists error should not be retried as it's a permanent condition

// Non-retryable error

func (o *operation) tryCreateResource(ctx context.Context, bindings apis.Bindings, obj unstructured.Unstructured) (outputs.Outputs, error) {
	_ = "STUB: not implemented"
	// First check if the resource exists
	return *new(outputs.Outputs), nil
}

// If there was an error other than NotFound, propagate it

// Resource doesn't exist, try to create it

// Resource created successfully

// Check if the error matches any expectations

// If the error is not AlreadyExists, propagate it

// Resource already exists (race condition)

// Resource already exists

func (o *operation) handleCheck(ctx context.Context, bindings apis.Bindings, obj unstructured.Unstructured, err error) (_outputs outputs.Outputs, _err error) {
	_ = "STUB: not implemented"
	return *new(outputs.Outputs), nil
}
