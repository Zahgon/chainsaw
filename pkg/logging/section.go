package logging

import (
	"fmt"
)

type section struct {
	name string
	args []any
}

func (s section) String() string { _ = "STUB: not implemented"; return "" }

func Section(name string, args ...any) fmt.Stringer {
	_ = "STUB: not implemented"
	return *new(fmt.Stringer)
}

func ErrSection(err error) fmt.Stringer { _ = "STUB: not implemented"; return *new(fmt.Stringer) }
