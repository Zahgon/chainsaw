package discovery

import (
	"k8s.io/apimachinery/pkg/labels"
)

func DiscoverTests(fileName string, selector labels.Selector, remarshal bool, paths ...string) ([]Test, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func discoverTests(fileName string, selector labels.Selector, remarshal bool, folders ...string) ([]Test, error) {
	_ = "STUB: not implemented"
	return nil, nil
}
