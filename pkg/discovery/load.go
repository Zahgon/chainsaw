package discovery

import (
	"github.com/kyverno/chainsaw/pkg/apis/v1alpha1"
)

func tryLoadTestFile(file string, remarshal bool) ([]*v1alpha1.Test, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func tryLoadTestFiles(fileName string, path string, remarshal bool) ([]*v1alpha1.Test, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func LoadTest(fileName string, path string, remarshal bool) ([]Test, error) {
	_ = "STUB: not implemented"
	// first, try to load a test manifest
	return nil, nil
}

// an absolute path points at a single specific file - don't fall back to
// per-folder step discovery which would ignore the path entirely

// next, look at files
