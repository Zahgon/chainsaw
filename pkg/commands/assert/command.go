package assert

import (
	"github.com/kyverno/chainsaw/pkg/client"
	nspacer "github.com/kyverno/chainsaw/pkg/engine/namespacer"
	"github.com/spf13/cobra"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
	"k8s.io/client-go/tools/clientcmd"
)

type options struct {
	timeout             metav1.Duration
	namespace           string
	noColor             bool
	assertPath          string
	resourcePath        string
	kubeConfigOverrides clientcmd.ConfigOverrides
}

func Command() *cobra.Command { _ = "STUB: not implemented"; return nil }

func preRunE(opts *options, _ *cobra.Command, args []string) error {
	_ = "STUB: not implemented"
	return nil
}

func runE(opts options, cmd *cobra.Command, client client.Client, namespacer nspacer.Namespacer) error {
	_ = "STUB: not implemented"
	return nil
}

// TODO: we should improve the lookup logic here

func assert(opts options, client client.Client, resource unstructured.Unstructured, namespacer nspacer.Namespacer) error {
	_ = "STUB: not implemented"
	return nil
}
