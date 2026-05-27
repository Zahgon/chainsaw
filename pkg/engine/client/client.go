package client

import (
	"context"

	"github.com/kyverno/chainsaw/pkg/client"
	"github.com/kyverno/chainsaw/pkg/logging"
	"k8s.io/apimachinery/pkg/api/meta"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apimachinery/pkg/types"
	ctrlclient "sigs.k8s.io/controller-runtime/pkg/client"
)

func New(inner client.Client) client.Client { _ = "STUB: not implemented"; return *new(client.Client) }

type runnerClient struct {
	inner client.Client
}

func (c *runnerClient) Create(ctx context.Context, obj client.Object, opts ...client.CreateOption) (_err error) {
	_ = "STUB: not implemented"
	return nil
}

func (c *runnerClient) Update(ctx context.Context, obj client.Object, opts ...client.UpdateOption) (_err error) {
	_ = "STUB: not implemented"
	return nil
}

func (c *runnerClient) Delete(ctx context.Context, obj client.Object, opts ...client.DeleteOption) (_err error) {
	_ = "STUB: not implemented"
	return nil
}

func (c *runnerClient) Get(ctx context.Context, key types.NamespacedName, obj client.Object, opts ...client.GetOption) error {
	_ = "STUB: not implemented"
	return nil
}

func (c *runnerClient) List(ctx context.Context, list client.ObjectList, opts ...client.ListOption) (_err error) {
	_ = "STUB: not implemented"
	return nil
}

func (c *runnerClient) Patch(ctx context.Context, obj client.Object, patch client.Patch, opts ...client.PatchOption) (_err error) {
	_ = "STUB: not implemented"
	return nil
}

func (c *runnerClient) IsObjectNamespaced(obj runtime.Object) (bool, error) {
	_ = "STUB: not implemented"
	return false, nil
}

func (c *runnerClient) RESTMapper() meta.RESTMapper {
	_ = "STUB: not implemented"
	return *new(meta.RESTMapper)
}

func (c *runnerClient) ok(ctx context.Context, op logging.Operation, obj client.Object) {
	_ = "STUB: not implemented"
	return
}

func (c *runnerClient) error(ctx context.Context, op logging.Operation, obj client.Object, err error) {
	_ = "STUB: not implemented"
	return
}

func (c *runnerClient) SubResource(subResource string) ctrlclient.SubResourceClient {
	_ = "STUB: not implemented"
	return *new(ctrlclient.SubResourceClient)
}
