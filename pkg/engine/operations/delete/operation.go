package delete

import (
	"context"

	"github.com/kyverno/chainsaw/pkg/apis"
	"github.com/kyverno/chainsaw/pkg/apis/v1alpha1"
	"github.com/kyverno/chainsaw/pkg/client"
	"github.com/kyverno/chainsaw/pkg/engine/namespacer"
	"github.com/kyverno/chainsaw/pkg/engine/operations"
	"github.com/kyverno/chainsaw/pkg/engine/outputs"
	"github.com/kyverno/kyverno-json/pkg/core/compilers"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
)

type operation struct {
	compilers         compilers.Compilers
	client            client.Client
	base              unstructured.Unstructured
	namespacer        namespacer.Namespacer
	template          bool
	expect            []v1alpha1.Expectation
	propagationPolicy metav1.DeletionPropagation
}

func New(
	compilers compilers.Compilers,
	client client.Client,
	obj unstructured.Unstructured,
	namespacer namespacer.Namespacer,
	template bool,
	propagationPolicy metav1.DeletionPropagation,
	expect ...v1alpha1.Expectation,
) operations.Operation {
	_ = "STUB: not implemented"
	return *new(operations.Operation)
}

func (o *operation) Exec(ctx context.Context, bindings apis.Bindings) (_ outputs.Outputs, _err error) {
	_ = "STUB: not implemented"
	return *new(outputs.Outputs), nil
}

func (o *operation) execute(ctx context.Context, bindings apis.Bindings, obj unstructured.Unstructured) error {
	_ = "STUB: not implemented"
	return nil
}

func (o *operation) getResourcesToDelete(ctx context.Context, obj unstructured.Unstructured) ([]unstructured.Unstructured, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (o *operation) deleteResources(ctx context.Context, bindings apis.Bindings, resources ...unstructured.Unstructured) error {
	_ = "STUB: not implemented"
	return nil
}

// if the resource was successfully deleted, record it to track actual deletion

// check if the result was the expected one

func (o *operation) deleteResource(ctx context.Context, resource unstructured.Unstructured) error {
	_ = "STUB: not implemented"
	return nil
}

func (o *operation) waitForDeletion(ctx context.Context, resource unstructured.Unstructured) error {
	_ = "STUB: not implemented"
	return nil
}

func (o *operation) handleCheck(ctx context.Context, bindings apis.Bindings, resource unstructured.Unstructured, err error) error {
	_ = "STUB: not implemented"
	return nil
}
