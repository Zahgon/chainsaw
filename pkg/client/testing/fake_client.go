package testing

import (
	"context"

	"k8s.io/apimachinery/pkg/api/meta"
	"k8s.io/apimachinery/pkg/runtime"
	client "sigs.k8s.io/controller-runtime/pkg/client"
)

// TODO: not thread safe
type FakeClient struct {
	GetFn                func(ctx context.Context, call int, key client.ObjectKey, obj client.Object, opts ...client.GetOption) error
	CreateFn             func(ctx context.Context, call int, obj client.Object, opts ...client.CreateOption) error
	UpdateFn             func(ctx context.Context, call int, obj client.Object, opts ...client.UpdateOption) error
	DeleteFn             func(ctx context.Context, call int, obj client.Object, opts ...client.DeleteOption) error
	ListFn               func(ctx context.Context, call int, list client.ObjectList, opts ...client.ListOption) error
	PatchFn              func(ctx context.Context, call int, obj client.Object, patch client.Patch, opts ...client.PatchOption) error
	SubResourceFn        func(subResource string) client.SubResourceClient
	IsObjectNamespacedFn func(call int, obj runtime.Object) (bool, error)
	RESTMapperFn         func(call int) meta.RESTMapper
	numCalls             int
}

func (c *FakeClient) Get(ctx context.Context, key client.ObjectKey, obj client.Object, opts ...client.GetOption) error {
	_ = "STUB: not implemented"
	return nil
}

func (c *FakeClient) List(ctx context.Context, list client.ObjectList, opts ...client.ListOption) error {
	_ = "STUB: not implemented"
	return nil
}

func (c *FakeClient) Create(ctx context.Context, obj client.Object, opts ...client.CreateOption) error {
	_ = "STUB: not implemented"
	return nil
}

func (c *FakeClient) Update(ctx context.Context, obj client.Object, opts ...client.UpdateOption) error {
	_ = "STUB: not implemented"
	return nil
}

func (c *FakeClient) Delete(ctx context.Context, obj client.Object, opts ...client.DeleteOption) error {
	_ = "STUB: not implemented"
	return nil
}

func (c *FakeClient) Patch(ctx context.Context, obj client.Object, patch client.Patch, opts ...client.PatchOption) error {
	_ = "STUB: not implemented"
	return nil
}

func (c *FakeClient) IsObjectNamespaced(obj runtime.Object) (bool, error) {
	_ = "STUB: not implemented"
	return false, nil
}

func (c *FakeClient) RESTMapper() meta.RESTMapper {
	_ = "STUB: not implemented"
	return *new(meta.RESTMapper)
}

func (c *FakeClient) SubResource(subResource string) client.SubResourceClient {
	_ = "STUB: not implemented"
	return *new(client.SubResourceClient)
}

func (c *FakeClient) NumCalls() int { _ = "STUB: not implemented"; return 0 }

type FakeSubResourceWriter struct {
	GetFn    func(ctx context.Context, obj client.Object, subResource client.Object, opts ...client.SubResourceGetOption) error
	CreateFn func(ctx context.Context, obj client.Object, subResource client.Object, opts ...client.SubResourceCreateOption) error
	UpdateFn func(ctx context.Context, obj client.Object, opts ...client.SubResourceUpdateOption) error
	PatchFn  func(ctx context.Context, obj client.Object, patch client.Patch, opts ...client.SubResourcePatchOption) error
	ApplyFn  func(ctx context.Context, obj runtime.ApplyConfiguration, opts ...client.SubResourceApplyOption) error
}

func (f *FakeSubResourceWriter) Get(ctx context.Context, obj client.Object, subResource client.Object, opts ...client.SubResourceGetOption) error {
	_ = "STUB: not implemented"
	return nil
}

func (f *FakeSubResourceWriter) Create(ctx context.Context, obj client.Object, subResource client.Object, opts ...client.SubResourceCreateOption) error {
	_ = "STUB: not implemented"
	return nil
}

func (f *FakeSubResourceWriter) Update(ctx context.Context, obj client.Object, opts ...client.SubResourceUpdateOption) error {
	_ = "STUB: not implemented"
	return nil
}

func (f *FakeSubResourceWriter) Patch(ctx context.Context, obj client.Object, patch client.Patch, opts ...client.SubResourcePatchOption) error {
	_ = "STUB: not implemented"
	return nil
}

func (f *FakeSubResourceWriter) Apply(ctx context.Context, obj runtime.ApplyConfiguration, opts ...client.SubResourceApplyOption) error {
	_ = "STUB: not implemented"
	return nil
}

func NewFakeSubResourceWriter() *FakeSubResourceWriter { _ = "STUB: not implemented"; return nil }
