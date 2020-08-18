package integration

import (
	httptransport "github.com/go-openapi/runtime/client"
	"github.com/go-openapi/strfmt"
	"github.com/ju-zp/tasker/integration/services/tasker/client"
	"os"
	"testing"
)

var taskerApi *client.JustaTodo

func TestMain(m *testing.M) {
	transport := httptransport.New("localhost:3000", client.DefaultBasePath, client.DefaultSchemes)
	taskerApi = client.New(transport, strfmt.Default)

	code := m.Run()
	os.Exit(code)
}