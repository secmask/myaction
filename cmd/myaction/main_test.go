package main

import (
	"io"
	"net/http"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestAB(t *testing.T) {
	resp, err := http.Get("https://ifconfig.me")
	require.NoError(t, err)

	data, err := io.ReadAll(resp.Body)
	require.NoError(t, err)
	t.Log(string(data))
}
