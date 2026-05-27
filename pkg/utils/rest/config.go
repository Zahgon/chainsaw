package rest

import (
	"io"

	"k8s.io/client-go/rest"
	"k8s.io/client-go/tools/clientcmd"
)

func DefaultConfig(overrides clientcmd.ConfigOverrides) (*rest.Config, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func Config(kubeconfigPath string, overrides clientcmd.ConfigOverrides) (*rest.Config, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func load(loader clientcmd.ClientConfigLoader, overrides clientcmd.ConfigOverrides) (*rest.Config, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func Save(cfg *rest.Config, w io.Writer) error { _ = "STUB: not implemented"; return nil }
