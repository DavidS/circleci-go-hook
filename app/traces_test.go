package app

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestTraceparent(t *testing.T) {
	workflow_id := "cedf175c-c537-49f8-9b5c-95892f0b2407"
	job_id := "3f6a7c4a-0bb6-463a-bd0c-3d5e0b1b1d34"
	assert.Equal(t, "export TRACEPARENT=00-cedf175cc53749f89b5c95892f0b2407-3f6a7c4a0bb6463a-01", translate_traceparent(workflow_id, job_id), `unexpected translate_traceparent result`)
}
