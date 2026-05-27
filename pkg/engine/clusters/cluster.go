package clusters

import (
	"k8s.io/client-go/rest"
)

type Cluster interface {
	Config() (*rest.Config, error)
}

type fromConfig struct {
	config *rest.Config
}

func NewClusterFromConfig(config *rest.Config) Cluster {
	_ = "STUB: not implemented"
	return *new(Cluster)
}

func (c *fromConfig) Config() (*rest.Config, error) { _ = "STUB: not implemented"; return nil, nil }

type fromKubeconfig struct {
	resolver func() (*rest.Config, error)
}

func NewClusterFromKubeconfig(kubeconfig string, context string) Cluster {
	_ = "STUB: not implemented"
	return *new(Cluster)
}

func (c *fromKubeconfig) Config() (*rest.Config, error) { _ = "STUB: not implemented"; return nil, nil }
