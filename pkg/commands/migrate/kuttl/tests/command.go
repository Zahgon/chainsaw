package tests

import (
	"io"

	"github.com/kyverno/chainsaw/pkg/apis/v1alpha1"
	"github.com/kyverno/chainsaw/pkg/discovery"
	"github.com/spf13/cobra"
	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
)

const schema = "# yaml-language-server: $schema=https://raw.githubusercontent.com/kyverno/chainsaw/main/.schemas/json/test-chainsaw-v1alpha1.json"

func Command() *cobra.Command { _ = "STUB: not implemented"; return nil }

func execute(stdout io.Writer, stderr io.Writer, save, cleanup bool, paths ...string) error {
	_ = "STUB: not implemented"
	return nil
}

func processFolder(stdout io.Writer, stderr io.Writer, folder string, save, cleanup bool) error {
	_ = "STUB: not implemented"
	return nil
}

func isKuttl(resource unstructured.Unstructured) bool { _ = "STUB: not implemented"; return false }

func processStep(stderr io.Writer, step *v1alpha1.TestStep, s discovery.Step, folder string, save bool) error {
	_ = "STUB: not implemented"
	return nil
}

// no cleanup

// no cleanup

// no cleanup

// no cleanup

// no cleanup

// no cleanup

func saveResources(stderr io.Writer, folder, file string, resources ...unstructured.Unstructured) error {
	_ = "STUB: not implemented"
	return nil
}

func prepend[T any](slice []T, elems ...T) []T { _ = "STUB: not implemented"; return nil }

func testStep(to *v1alpha1.TestStepSpec, in unstructured.Unstructured) error {
	_ = "STUB: not implemented"
	return nil
}

func testAssert(to *v1alpha1.TestStepSpec, in unstructured.Unstructured) error {
	_ = "STUB: not implemented"
	return nil
}

// TODO: timeout
