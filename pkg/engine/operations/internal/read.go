package internal

import (
	"context"

	"github.com/kyverno/chainsaw/pkg/client"
	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
)

func Read(ctx context.Context, expected client.Object, c client.Client) ([]unstructured.Unstructured, error) {
	_ = "STUB: not implemented"
	return nil, nil
}
