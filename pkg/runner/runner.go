package runner

import (
	"context"

	"github.com/kyverno/chainsaw/pkg/apis/v1alpha1"
	"github.com/kyverno/chainsaw/pkg/apis/v1alpha2"
	"github.com/kyverno/chainsaw/pkg/cleanup/cleaner"
	"github.com/kyverno/chainsaw/pkg/discovery"
	"github.com/kyverno/chainsaw/pkg/model"
	enginecontext "github.com/kyverno/chainsaw/pkg/runner/context"
	"github.com/kyverno/chainsaw/pkg/runner/internal"
	"github.com/kyverno/chainsaw/pkg/runner/operations"
	"k8s.io/utils/clock"
)

type Runner interface {
	Run(context.Context, v1alpha2.NamespaceOptions, enginecontext.TestContext, ...discovery.Test) error
}

func New(clock clock.PassiveClock, onFailure func()) Runner {
	_ = "STUB: not implemented"
	return *new(Runner)
}

type runner struct {
	clock     clock.PassiveClock
	onFailure func()
	deps      *internal.TestDeps
}

func (r *runner) Run(ctx context.Context, nsOptions v1alpha2.NamespaceOptions, tc enginecontext.TestContext, tests ...discovery.Test) error {
	_ = "STUB: not implemented"
	return nil
}

func (r *runner) run(ctx context.Context, m mainstart, nsOptions v1alpha2.NamespaceOptions, tc enginecontext.TestContext, tests ...discovery.Test) error {
	_ = "STUB: not implemented"
	return nil
}

// sanity check

// setup logger sink

// setup logger

// setup cleanup

// setup namespace

// loop through tests

// setup logger

// helper to run test

// setup logger sink

// setup concurrency

// setup reporting

// setup summary - must run after all other cleanup (registered first to run last in LIFO order)

// skip check

// setup context

// fail fast check

// prepare namespace
// TODO: should be part of setupContext ?

// setup cleaner

// setup namespace

// setup bindings

// loop through steps

// run test scenarios

// m.Run() returns:
// - 0 if everything went well
// - 1 if some of the tests failed
// - 2 if running the tests was not possible
// In our case, we consider an error only when running the tests was not possible.
// For now, the case where some of the tests failed will be covered by the summary.

func (r *runner) runStep(
	ctx context.Context,
	cleanup func(func()),
	fail func(),
	failed func() bool,
	tc enginecontext.TestContext,
	step v1alpha1.TestStep,
	testReport *model.TestReport,
) bool {
	_ = "STUB: not implemented"
	return false
}

func (r *runner) runOperation(
	ctx context.Context,
	tc enginecontext.TestContext,
	operation v1alpha1.Operation,
	operationId int,
	cleaner cleaner.Cleaner,
	report *model.StepReport,
) (bool, enginecontext.TestContext, error) {
	_ = "STUB: not implemented"
	return false, *new(enginecontext.TestContext), nil
}

func (r *runner) runCatch(
	ctx context.Context,
	tc enginecontext.TestContext,
	operation v1alpha1.CatchFinally,
	operationId int,
) (enginecontext.TestContext, error) {
	_ = "STUB: not implemented"
	return *new(enginecontext.TestContext), nil
}

func (*runner) runAction(
	ctx context.Context,
	action operations.Operation,
	opType model.OperationType,
	operationId int,
	actionId int,
	tc enginecontext.TestContext,
	stepReport *model.StepReport,
) (outputs map[string]any, err error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (r *runner) onFail() { _ = "STUB: not implemented"; return }

func (r *runner) cleanup(ctx context.Context, tc enginecontext.TestContext, cleaner cleaner.Cleaner) error {
	_ = "STUB: not implemented"
	return nil
}

func (r *runner) setupNamespace(ctx context.Context, nsOptions v1alpha2.NamespaceOptions, tc enginecontext.TestContext, cleanup cleaner.Cleaner) (enginecontext.TestContext, error) {
	_ = "STUB: not implemented"
	return *new(enginecontext.TestContext), nil
}

func (r *runner) setupTestContext(ctx context.Context, testId int, scenarioId int, tc enginecontext.TestContext, test discovery.Test, bindings ...v1alpha1.Binding) (enginecontext.TestContext, error) {
	_ = "STUB: not implemented"
	return *new(enginecontext.TestContext), nil
}

func (r *runner) testCleanup(ctx context.Context, tc enginecontext.TestContext, cleaner cleaner.Cleaner, report *model.TestReport) error {
	_ = "STUB: not implemented"
	return nil
}
