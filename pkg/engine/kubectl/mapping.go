package kubectl

import (
	"context"

	"github.com/kyverno/chainsaw/pkg/apis"
	"github.com/kyverno/chainsaw/pkg/apis/v1alpha1"
	"github.com/kyverno/chainsaw/pkg/client"
	"github.com/kyverno/kyverno-json/pkg/core/compilers"
	"k8s.io/apimachinery/pkg/api/meta"
	"k8s.io/apimachinery/pkg/runtime/schema"
)

func mapResource(ctx context.Context, compilers compilers.Compilers, client client.Client, tc apis.Bindings, resource v1alpha1.ObjectType) (string, bool, error) {
	_ = "STUB: not implemented"
	return "", false, nil
}

func mapResourceFromApiVersionAndKind(client client.Client, apiVersion string, kind string) (string, bool, error) {
	_ = "STUB: not implemented"
	return "", false, nil
}

func mapResourceFromGVK(mapper meta.RESTMapper, gvk schema.GroupVersionKind) (string, bool, error) {
	_ = "STUB: not implemented"
	return "", false, nil
}
