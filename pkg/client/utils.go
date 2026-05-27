package client

import (
	"context"
	"time"

	"github.com/kyverno/pkg/ext/output/color"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime"
)

const PollInterval = 50 * time.Millisecond

func Key(obj metav1.Object) ObjectKey { _ = "STUB: not implemented"; return *new(ObjectKey) }

func Name(key ObjectKey) string { _ = "STUB: not implemented"; return "" }

func ColouredName(key ObjectKey, color *color.Color) string { _ = "STUB: not implemented"; return "" }

func PatchObject(actual, expected runtime.Object) (runtime.Object, error) {
	_ = "STUB: not implemented"
	return *new(runtime.Object), nil
}

func WaitForDeletion(ctx context.Context, client Client, object Object) error {
	_ = "STUB: not implemented"
	return nil
}
