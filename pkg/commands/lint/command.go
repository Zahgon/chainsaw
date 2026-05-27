package lint

import (
	"io"

	"github.com/spf13/cobra"
)

func Command() *cobra.Command { _ = "STUB: not implemented"; return nil }

func lintInput(input []byte, kind string, format string, writer io.Writer) error {
	_ = "STUB: not implemented"
	return nil
}

func lintSchema(input []byte, kind string, format string, writer io.Writer) error {
	_ = "STUB: not implemented"
	return nil
}
