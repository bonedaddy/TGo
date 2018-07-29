package tgo

import (
	"net/http"
	"time"
)

type RPC struct {
	URL    string
	Client *http.Client
}

func GenerateClient(rpcURL string, timeout time.Duration) *RPC {
	rpc := RPC{}
	client := http.DefaultClient
	client.Timeout = timeout
	rpc.Client = client
	rpc.URL = rpcURL
	return &rpc
}
