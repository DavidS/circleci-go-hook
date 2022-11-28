package app

import (
	"fmt"
	"strings"

	"go.opentelemetry.io/otel/trace"
)

func translate_traceparent(workflow_id string, job_id string) string {
	trace_id, err := trace.TraceIDFromHex(strings.ReplaceAll(workflow_id,
		"-", ""))
	if err != nil {
		panic(err)
	}
	span_id, err := trace.SpanIDFromHex(strings.ReplaceAll(job_id,
		"-", "")[:16])
	if err != nil {
		panic(err)
	}
	return fmt.Sprintf("export TRACEPARENT=%02x-%v-%v-%v", 0, trace_id,
		span_id, trace.FlagsSampled,
	)
}
