package internal

import (
	"bytes"
	"fmt"
)

type CommandOutput struct {
	Stdout bytes.Buffer
	Stderr bytes.Buffer
}

func (c *CommandOutput) Out() string { _ = "STUB: not implemented"; return "" }

func (c *CommandOutput) Err() string { _ = "STUB: not implemented"; return "" }

func (c *CommandOutput) Sections() []fmt.Stringer { _ = "STUB: not implemented"; return nil }
