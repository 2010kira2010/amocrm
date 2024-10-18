package amocrm

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestEndpoint_Path(t *testing.T) {
	e := endpoint("example")
	path := e.path()
	require.IsType(t, "", path)
	require.Contains(t, path, "/api/v")
	require.Contains(t, path, "/example")
}
