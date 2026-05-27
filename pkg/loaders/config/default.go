package config

import (
	"sync"

	"github.com/kyverno/chainsaw/pkg/apis/v1alpha2"
)

func defaultConfiguration(_fs func() ([]byte, error)) (*v1alpha2.Configuration, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

var DefaultConfiguration = sync.OnceValues(func() (*v1alpha2.Configuration, error) { return defaultConfiguration(nil) })
