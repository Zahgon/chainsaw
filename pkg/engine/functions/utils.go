package functions

// stable is a utility function that returns the function name as-is.
// Used for stable, non-experimental JMESPath functions.
func stable(in string) string {
	_ = "STUB: not implemented"

	// experimental is a utility function that prefixes the function name with "x_".
	// Used for experimental JMESPath functions that may change in future versions.
	return ""
}

func experimental(in string) string {
	_ = "STUB: not implemented"

	// getArgAt retrieves an argument at the specified index from the arguments slice.
	// Returns an error if the index is out of range.
	return ""
}

func getArgAt(arguments []any, index int) (any, error) {
	_ = "STUB: not implemented"
	return *new(any), nil
}

// getArg is a generic function that retrieves and type-asserts an argument at the specified index.
// Arguments:
// - arguments: The slice of arguments to retrieve from
// - index: The index of the argument to retrieve
// - out: A pointer to store the retrieved value
// Returns an error if the index is out of range or if the type assertion fails.
func getArg[T any](arguments []any, index int, out *T) error { _ = "STUB: not implemented"; return nil }
