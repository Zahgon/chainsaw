package command

import (
	"context"
	"os/exec"

	"github.com/kyverno/chainsaw/pkg/apis"
	"github.com/kyverno/chainsaw/pkg/apis/v1alpha1"
	"github.com/kyverno/chainsaw/pkg/engine/operations"
	"github.com/kyverno/chainsaw/pkg/engine/outputs"
	"github.com/kyverno/kyverno-json/pkg/core/compilers"
	"k8s.io/client-go/rest"
)

type operation struct {
	compilers compilers.Compilers
	command   v1alpha1.Command
	basePath  string
	namespace string
	cfg       *rest.Config
}

func New(
	compilers compilers.Compilers,
	command v1alpha1.Command,
	basePath string,
	namespace string,
	cfg *rest.Config,
) operations.Operation {
	_ = "STUB: not implemented"
	return *new(operations.Operation)
}

func (o *operation) Exec(ctx context.Context, bindings apis.Bindings) (_ outputs.Outputs, _err error) {
	_ = "STUB: not implemented"
	return *new(outputs.Outputs), nil
}

func (o *operation) createCommand(ctx context.Context, bindings apis.Bindings) (*exec.Cmd, context.CancelFunc, error) {
	_ = "STUB: not implemented"
	return nil, *new(context.CancelFunc), nil
}

//nolint:gosec

func (o *operation) execute(ctx context.Context, bindings apis.Bindings, cmd *exec.Cmd) (_outputs outputs.Outputs, _err error) {
	_ = "STUB: not implemented"
	return *new(outputs.Outputs), nil
}
