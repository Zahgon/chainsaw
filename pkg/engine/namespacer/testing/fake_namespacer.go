package testing

import (
	"github.com/kyverno/chainsaw/pkg/client"
)

type FakeNamespacer struct {
	ApplyFn        func(call int, client client.Client, obj client.Object) error
	GetNamespaceFn func(call int) string
	numCalls       int
}

func (n *FakeNamespacer) Apply(client client.Client, obj client.Object) error {
	_ = "STUB: not implemented"
	return nil
}

func (n *FakeNamespacer) GetNamespace() string { _ = "STUB: not implemented"; return "" }
