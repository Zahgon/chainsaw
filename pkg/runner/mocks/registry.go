package mocks

import (
	"github.com/kyverno/chainsaw/pkg/client"
	"github.com/kyverno/chainsaw/pkg/engine/clusters"
	"k8s.io/client-go/rest"
)

type Registry struct {
	Client client.Client
}

func (r Registry) Register(string, clusters.Cluster) clusters.Registry {
	_ = "STUB: not implemented"
	return *new(clusters.Registry)
}

func (r Registry) Lookup(string) clusters.Cluster {
	_ = "STUB: not implemented"
	return *new(clusters.Cluster)
}

func (r Registry) Build(clusters.Cluster) (*rest.Config, client.Client, error) {
	_ = "STUB: not implemented"
	return nil, *new(client.Client), nil
}
