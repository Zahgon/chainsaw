package config

import (
	"io"

	"github.com/kyverno/chainsaw/pkg/apis/v1alpha1"
	"github.com/spf13/cobra"
	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
)

const schema = "# yaml-language-server: $schema=https://raw.githubusercontent.com/kyverno/chainsaw/main/.schemas/json/configuration-chainsaw-v1alpha1.json"

func Command() *cobra.Command { _ = "STUB: not implemented"; return nil }

func execute(stdout io.Writer, stderr io.Writer, save, cleanup bool, path string) error {
	_ = "STUB: not implemented"
	return nil
}

func isKuttl(resource unstructured.Unstructured) bool { _ = "STUB: not implemented"; return false }

func migrate(stderr io.Writer, path string, resource unstructured.Unstructured) (*v1alpha1.Configuration, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func testSuite(in unstructured.Unstructured) (*v1alpha1.Configuration, error) {
	_ = "STUB: not implemented"
	return nil, nil
}
