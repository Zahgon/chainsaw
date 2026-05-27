package namespacer

import (
	"github.com/kyverno/chainsaw/pkg/client"
)

type Namespacer interface {
	Apply(client.Client, client.Object) error
	GetNamespace() string
}

type namespacer struct {
	namespace string
}

func New(namespace string) Namespacer { _ = "STUB: not implemented"; return *new(Namespacer) }

func (n *namespacer) Apply(client client.Client, resource client.Object) error {
	_ = "STUB: not implemented"
	return nil
}

func (n *namespacer) GetNamespace() string { _ = "STUB: not implemented"; return "" }
