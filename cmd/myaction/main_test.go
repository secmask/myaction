package main

import (
	"github.com/stretchr/testify/require"
	"io"
	"net/http"
	"testing"
)

func TestAA(t *testing.T) {

}

func TestAB(t *testing.T) {
	resp, err := http.Get("https://ifconfig.me")
	require.NoError(t, err)
	data, err := io.ReadAll(resp.Body)
	require.NoError(t, err)
	t.Log(string(data))
}
