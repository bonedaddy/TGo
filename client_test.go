package tgo_test

import (
	"fmt"
	"testing"
	"time"

	tgo "github.com/postables/TGo"
)

var rpcURL = "http://127.0.0.1:3090"

func TestClient(t *testing.T) {
	client := tgo.GenerateClient(rpcURL, time.Minute)
	connections, err := client.GetConnections()
	if err != nil {
		t.Fatal(err)
	}
	fmt.Printf("%+v\n", connections)

	peer, err := client.GetPeerID(connections[0].PeerID)
	if err != nil {
		t.Fatal(err)
	}
	fmt.Printf("%+v\n", peer)
}
