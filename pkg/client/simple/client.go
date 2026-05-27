package simple

import (
	"github.com/kyverno/chainsaw/pkg/client"
	"k8s.io/client-go/rest"
)

func New(cfg *rest.Config) (client.Client, error) {
	_ = "STUB: not implemented"
	return *new(client.Client), nil
}
