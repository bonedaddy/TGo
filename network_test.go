package tgo_test

import (
	"fmt"
	"testing"
	"time"

	tgo "github.com/postables/TGo"
)

func TestGetconnections(t *testing.T) {
	client := tgo.GenerateClient(tgo.RpcURL, time.Minute)

	connections, err := client.GetConnections()
	if err != nil {
		t.Fatal(err)
	}

	fmt.Printf("%+v\n", connections)
}
