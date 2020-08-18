package integration

import (
	"github.com/ju-zp/tasker/integration/services/tasker/client/operations"
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test__Can_ping_successfully(t *testing.T) {
	_, err := taskerApi.Operations.GetPing(operations.NewGetPingParams())
	assert.NoError(t, err)
}
