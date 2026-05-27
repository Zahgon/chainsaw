package config

import (
	"io"

	"github.com/kyverno/chainsaw/pkg/apis/v1alpha2"
	"github.com/spf13/cobra"
	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
)

const schema = "# yaml-language-server: $schema=https://raw.githubusercontent.com/kyverno/chainsaw/main/.schemas/json/configuration-chainsaw-v1alpha2.json"

func Command() *cobra.Command { _ = "STUB: not implemented"; return nil }

func execute(stdout io.Writer, stderr io.Writer, save bool, file string) error {
	_ = "STUB: not implemented"
	return nil
}

func loadConfig(file string) (*v1alpha2.Configuration, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func loadDefaultConfig() (*v1alpha2.Configuration, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func buildConfigPatch(c *v1alpha2.Configuration) (*unstructured.Unstructured, error) {
	_ = "STUB: not implemented"
	return nil, nil
}
