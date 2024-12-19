package assert

import (
	"io"
	"os"
)

var Writter io.Writer = os.Stderr

// ContextWindow represents the number of source code lines to display above and
// below the line that caused the assertion failure.
var ContextWindow int = 5

// MaxTraceDepth represents the maximum number of stack frames to display in the
// stack trace.
var MaxTraceDepth int = 20

// ReturnValue represents the exit code to use when an assertion fails.
var ReturnValue int = 1 // Not using 42069 requires a lot of self control
