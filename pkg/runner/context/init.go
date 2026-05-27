package context

import (
	"github.com/kyverno/chainsaw/pkg/model"
	"k8s.io/client-go/rest"
)

func InitContext(config model.Configuration, defaultCluster *rest.Config, values any) (TestContext, error) {
	_ = "STUB: not implemented"
	return *new(TestContext), nil
}

// cleanup options

// templating options

// discovery options

// execution options

// deletion options

// error options

// timeouts

// values

// clusters
