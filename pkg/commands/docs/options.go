package docs

import (
	"github.com/spf13/cobra"
)

type options struct {
	path       string
	website    bool
	autogenTag bool
}

func (o options) validate(root *cobra.Command) error { _ = "STUB: not implemented"; return nil }

func (o options) execute(root *cobra.Command) error { _ = "STUB: not implemented"; return nil }
