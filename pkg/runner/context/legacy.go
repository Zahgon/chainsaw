package context

import (
	"github.com/kyverno/chainsaw/pkg/apis/v1alpha1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

type ContextData struct {
	BasePath            *string
	Catch               []v1alpha1.CatchFinally
	Cluster             *string
	Clusters            v1alpha1.Clusters
	DelayBeforeCleanup  *metav1.Duration
	DeletionPropagation *metav1.DeletionPropagation
	DryRun              *bool
	SkipDelete          *bool
	Templating          *bool
	TerminationGrace    *metav1.Duration
	Timeouts            *v1alpha1.Timeouts
}

func SetupContext(tc TestContext, data ContextData) (TestContext, error) {
	_ = "STUB: not implemented"
	return *new(TestContext), nil
}

func SetupBindings(tc TestContext, bindings ...v1alpha1.Binding) (TestContext, error) {
	_ = "STUB: not implemented"
	return *new(TestContext), nil
}

func SetupContextAndBindings(tc TestContext, data ContextData, bindings ...v1alpha1.Binding) (TestContext, error) {
	_ = "STUB: not implemented"
	return *new(TestContext), nil
}
