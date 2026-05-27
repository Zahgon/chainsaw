package expressions

import (
	"context"
	"regexp"
)

var (
	escapeRegex = regexp.MustCompile(`^\\(.+)\\$`)
	engineRegex = regexp.MustCompile(`^\((?:(\w+);)?(.+)\)$`)
)

type Expression struct {
	Statement string
	Engine    string
}

func Parse(ctx context.Context, value string) *Expression { _ = "STUB: not implemented"; return nil }

func parseExpressionRegex(_ context.Context, in string) *Expression {
	_ = "STUB: not implemented"
	return nil

	// 1. match escape, if there's no escaping then match engine
}

// account for default engine

// parse statement
