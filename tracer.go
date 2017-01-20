package trace

import (
	"fmt"
	"io"
)

// Tracer is the interface that describes an object capable of tracing events
// throughout code.
type Tracer interface {
	Trace(...interface{})
}

type tracer struct {
	out io.Writer
}

func (t *tracer) Trace(a ...interface{}) {
	fmt.Fprint(t.out, a...)
	fmt.Fprintln(t.out)
}

// New returns a new tracer that writes to the given Writer
func New(w io.Writer) Tracer {
	return &tracer{out: w}
}

type nilTracer struct{}

func (t *nilTracer) Trace(a ...interface{}) {}

// Off returns a new tracer that doesn't write anywhere
func Off() Tracer {
	return &nilTracer{}
}
