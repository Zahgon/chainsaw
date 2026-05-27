package diff

import (
	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
)

func prune(expected map[string]any, actual map[string]any) { _ = "STUB: not implemented"; return }

func pruneMetadata(expected map[string]any, actual map[string]any) {
	_ = "STUB: not implemented"
	return
}

func pruneRoot(expected map[string]any, actual map[string]any) { _ = "STUB: not implemented"; return }

func PrettyDiff(expected unstructured.Unstructured, actual unstructured.Unstructured) (string, error) {
	_ = "STUB: not implemented"
	return "", nil
}
