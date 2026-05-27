package context

import (
	"time"

	"github.com/kyverno/chainsaw/pkg/apis"
	"github.com/kyverno/chainsaw/pkg/apis/v1alpha1"
	"github.com/kyverno/chainsaw/pkg/client"
	"github.com/kyverno/chainsaw/pkg/engine/clusters"
	"github.com/kyverno/chainsaw/pkg/engine/namespacer"
	"github.com/kyverno/chainsaw/pkg/model"
	"github.com/kyverno/kyverno-json/pkg/core/compilers"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/rest"
	"k8s.io/utils/clock"
)

type Timeouts struct {
	Apply   time.Duration
	Assert  time.Duration
	Cleanup time.Duration
	Delete  time.Duration
	Error   time.Duration
	Exec    time.Duration
}

type TestContext struct {
	*model.Summary
	*model.Report
	basePath            string
	bindings            apis.Bindings
	catch               []v1alpha1.CatchFinally
	cluster             clusters.Cluster
	clusters            clusters.Registry
	compilers           compilers.Compilers
	delayBeforeCleanup  *time.Duration
	deletionPropagation metav1.DeletionPropagation
	dryRun              bool
	failFast            bool
	fullName            bool
	namespacer          namespacer.Namespacer
	quiet               bool
	skipDelete          bool
	templating          bool
	terminationGrace    *time.Duration
	timeouts            Timeouts
}

func MakeContext(clock clock.PassiveClock, bindings apis.Bindings, registry clusters.Registry) TestContext {
	_ = "STUB: not implemented"
	return *new(TestContext)
}

func EmptyContext(clock clock.PassiveClock) TestContext {
	_ = "STUB: not implemented"
	return *new(TestContext)
}

func (tc *TestContext) Bindings() apis.Bindings {
	_ = "STUB: not implemented"
	return *new(apis.Bindings)
}

func (tc *TestContext) BasePath() string { _ = "STUB: not implemented"; return "" }

func (tc *TestContext) Catch() []v1alpha1.CatchFinally { _ = "STUB: not implemented"; return nil }

func (tc *TestContext) Cluster(name string) clusters.Cluster {
	_ = "STUB: not implemented"
	return *new(clusters.Cluster)
}

func (tc *TestContext) Clusters() clusters.Registry {
	_ = "STUB: not implemented"
	return *new(clusters.Registry)
}

func (tc *TestContext) Compilers() compilers.Compilers {
	_ = "STUB: not implemented"
	return *new(compilers.Compilers)
}

func (tc *TestContext) CurrentCluster() clusters.Cluster {
	_ = "STUB: not implemented"
	return *new(clusters.Cluster)
}

func (tc *TestContext) CurrentClusterClient() (*rest.Config, client.Client, error) {
	_ = "STUB: not implemented"
	return nil, *new(client.Client), nil
}

func (tc *TestContext) DelayBeforeCleanup() *time.Duration { _ = "STUB: not implemented"; return nil }

func (tc *TestContext) DeletionPropagation() metav1.DeletionPropagation {
	_ = "STUB: not implemented"
	return *new(metav1.DeletionPropagation)
}

func (tc *TestContext) DryRun() bool { _ = "STUB: not implemented"; return false }

func (tc *TestContext) FailFast() bool { _ = "STUB: not implemented"; return false }

func (tc *TestContext) FullName() bool { _ = "STUB: not implemented"; return false }

func (tc *TestContext) Namespacer() namespacer.Namespacer {
	_ = "STUB: not implemented"
	return *new(namespacer.Namespacer)
}

func (tc *TestContext) Quiet() bool { _ = "STUB: not implemented"; return false }

func (tc *TestContext) SkipDelete() bool { _ = "STUB: not implemented"; return false }

func (tc *TestContext) Templating() bool { _ = "STUB: not implemented"; return false }

func (tc *TestContext) TerminationGrace() *time.Duration { _ = "STUB: not implemented"; return nil }

func (tc *TestContext) Timeouts() Timeouts { _ = "STUB: not implemented"; return *new(Timeouts) }

func (tc TestContext) WithBasePath(basePath string) TestContext {
	_ = "STUB: not implemented"
	return *new(TestContext)
}

func (tc TestContext) WithBinding(name string, value any) TestContext {
	_ = "STUB: not implemented"
	return *new(TestContext)
}

func (tc TestContext) WithCatch(catch ...v1alpha1.CatchFinally) TestContext {
	_ = "STUB: not implemented"
	return *new(TestContext)
}

func (tc TestContext) WithCluster(name string, cluster clusters.Cluster) TestContext {
	_ = "STUB: not implemented"
	return *new(TestContext)
}

func (tc TestContext) WithCurrentCluster(name string) TestContext {
	_ = "STUB: not implemented"
	return *new(TestContext)
}

func (tc TestContext) WithDefaultCompiler(name string) TestContext {
	_ = "STUB: not implemented"
	return *new(TestContext)
}

func (tc TestContext) WithDelayBeforeCleanup(delayBeforeCleanup *time.Duration) TestContext {
	_ = "STUB: not implemented"
	return *new(TestContext)
}

func (tc TestContext) WithDeletionPropagation(deletionPropagation metav1.DeletionPropagation) TestContext {
	_ = "STUB: not implemented"
	return *new(TestContext)
}

func (tc TestContext) WithDryRun(dryRun bool) TestContext {
	_ = "STUB: not implemented"
	return *new(TestContext)
}

func (tc TestContext) WithFailFast(failFast bool) TestContext {
	_ = "STUB: not implemented"
	return *new(TestContext)
}

func (tc TestContext) WithFullName(fullName bool) TestContext {
	_ = "STUB: not implemented"
	return *new(TestContext)
}

func (tc TestContext) WithNamespacer(namespacer namespacer.Namespacer) TestContext {
	_ = "STUB: not implemented"
	return *new(TestContext)
}

func (tc TestContext) WithQuiet(quiet bool) TestContext {
	_ = "STUB: not implemented"
	return *new(TestContext)
}

func (tc TestContext) WithSkipDelete(skipDelete bool) TestContext {
	_ = "STUB: not implemented"
	return *new(TestContext)
}

func (tc TestContext) WithTemplating(templating bool) TestContext {
	_ = "STUB: not implemented"
	return *new(TestContext)
}

func (tc TestContext) WithTerminationGrace(terminationGrace *time.Duration) TestContext {
	_ = "STUB: not implemented"
	return *new(TestContext)
}

func (tc TestContext) WithTimeouts(timeouts v1alpha1.Timeouts) TestContext {
	_ = "STUB: not implemented"
	return *new(TestContext)
}
