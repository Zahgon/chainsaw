package dryrun

import (
	"context"

	"github.com/kyverno/chainsaw/pkg/client"
	"k8s.io/apimachinery/pkg/api/meta"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apimachinery/pkg/types"
	ctrlclient "sigs.k8s.io/controller-runtime/pkg/client"
)

type Client = client.Client

type dryRunClient struct {
	inner Client
}

func (c *dryRunClient) Create(ctx context.Context, obj client.Object, opts ...client.CreateOption) error {
	_ = "STUB: not implemented"
	return nil
}

func (c *dryRunClient) Update(ctx context.Context, obj client.Object, opts ...client.UpdateOption) error {
	_ = "STUB: not implemented"
	return nil
}

func (c *dryRunClient) Delete(ctx context.Context, obj client.Object, opts ...client.DeleteOption) error {
	_ = "STUB: not implemented"
	return nil
}

func (c *dryRunClient) Get(ctx context.Context, key types.NamespacedName, obj client.Object, opts ...client.GetOption) error {
	_ = "STUB: not implemented"
	return nil
}

func (c *dryRunClient) IsObjectNamespaced(obj runtime.Object) (bool, error) {
	_ = "STUB: not implemented"
	return false, nil
}

func (c *dryRunClient) List(ctx context.Context, list client.ObjectList, opts ...client.ListOption) error {
	_ = "STUB: not implemented"
	return nil
}

func (c *dryRunClient) Patch(ctx context.Context, obj client.Object, patch client.Patch, opts ...client.PatchOption) error {
	_ = "STUB: not implemented"
	return nil
}

func (c *dryRunClient) RESTMapper() meta.RESTMapper {
	_ = "STUB: not implemented"
	return *new(meta.RESTMapper)
}

func (c *dryRunClient) SubResource(subResource string) ctrlclient.SubResourceClient {
	_ = "STUB: not implemented"
	return *new(ctrlclient.SubResourceClient)
}

func New(inner Client) Client { _ = "STUB: not implemented"; return *new(Client) }
