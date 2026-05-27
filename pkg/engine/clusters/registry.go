package clusters

import (
	"github.com/kyverno/chainsaw/pkg/client"
	"k8s.io/client-go/rest"
)

const DefaultClient = ""

type Registry interface {
	Register(string, Cluster) Registry
	Lookup(string) Cluster
	Build(Cluster) (*rest.Config, client.Client, error)
}

type clientFactory = func(Cluster) (*rest.Config, client.Client, error)

func defaultClientFactory(cluster Cluster) (*rest.Config, client.Client, error) {
	_ = "STUB: not implemented"
	return nil, *new(client.Client), nil
}

type registry struct {
	clientFactory clientFactory
	clusters      map[string]Cluster
}

func NewRegistry(f clientFactory) Registry { _ = "STUB: not implemented"; return *new(Registry) }

func (c registry) Register(name string, cluster Cluster) Registry {
	_ = "STUB: not implemented"
	return *new(Registry)
}

func (c registry) Lookup(name string) Cluster { _ = "STUB: not implemented"; return *new(Cluster) }

func (c registry) Build(cluster Cluster) (*rest.Config, client.Client, error) {
	_ = "STUB: not implemented"
	return nil, *new(client.Client), nil
}
