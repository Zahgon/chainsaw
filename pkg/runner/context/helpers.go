package context

import (
	"github.com/kyverno/chainsaw/pkg/apis/v1alpha1"
)

func WithBindings(tc TestContext, variables ...v1alpha1.Binding) (TestContext, error) {
	_ = "STUB: not implemented"
	return *new(TestContext), nil
}

func WithClusters(tc TestContext, c map[string]v1alpha1.Cluster) TestContext {
	_ = "STUB: not implemented"
	return *new(TestContext)
}

func WithCurrentCluster(tc TestContext, name string) (TestContext, error) {
	_ = "STUB: not implemented"
	return *new(TestContext), nil
}

func withValues(tc TestContext, values any) TestContext {
	_ = "STUB: not implemented"
	return *new(TestContext)
}
