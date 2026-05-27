package cleaner

import (
	"context"
	"time"

	"github.com/kyverno/chainsaw/pkg/client"
	"github.com/kyverno/chainsaw/pkg/model"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

type cleanupEntry struct {
	client client.Client
	object client.Object
}

type CleanerCollector interface {
	Empty() bool
	Add(client.Client, client.Object)
}

type Cleaner interface {
	CleanerCollector
	Run(ctx context.Context, stepReport *model.StepReport) []error
}

func New(timeout time.Duration, waitForDeletion bool, delay *time.Duration, propagation metav1.DeletionPropagation) Cleaner {
	_ = "STUB: not implemented"
	return *new(Cleaner)
}

type cleaner struct {
	delay           *time.Duration
	timeout         time.Duration
	propagation     metav1.DeletionPropagation
	waitForDeletion bool
	entries         []cleanupEntry
}

func (c *cleaner) Add(client client.Client, object client.Object) {
	_ = "STUB: not implemented"
	return
}

func (c *cleaner) Empty() bool { _ = "STUB: not implemented"; return false }

func (c *cleaner) Run(ctx context.Context, stepReport *model.StepReport) []error {
	_ = "STUB: not implemented"
	return nil
}

func (c *cleaner) delete(ctx context.Context, entry cleanupEntry) error {
	_ = "STUB: not implemented"
	return nil
}
