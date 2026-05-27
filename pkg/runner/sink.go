package runner

import (
	"github.com/kyverno/chainsaw/pkg/logging"
	"k8s.io/utils/clock"
)

const eraser = "\b\b\b\b\b\b\b\b\b\b\b\b"

// newSink creates a new sink for logging.
// If quietMode is true, it will buffer non-error messages and flush them when an error occurs.
// This ensures diagnostic details (catch output, failed command output) are shown on failures.
// Otherwise, it will log all messages immediately.
func newSink(clock clock.PassiveClock, quiet bool, log func(args ...any)) logging.SinkFunc {
	_ = "STUB: not implemented"
	// Use a map to maintain separate buffers per test+step combination
	// This ensures parallel tests don't mix their output
	return *new(logging.SinkFunc)
}

// Track which test+step combinations have encountered errors
// Once an error occurs, all subsequent logs for that test+step are shown immediately

// Not in quiet mode - log everything immediately

// In quiet mode with error/internal: flush buffer and mark error occurred

// Clear buffer after flushing

// After an error, show all subsequent logs immediately (including catch blocks)

// In quiet mode before any error: buffer non-error messages per test+step
