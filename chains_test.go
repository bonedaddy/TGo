package tgo_test

import (
	"fmt"
	"testing"
	"time"

	"github.com/postables/TGo"
)

func TestChains(t *testing.T) {
	client := tgo.GenerateClient(tgo.RpcURL, time.Minute)
	err := client.GetChainID("main")
	if err != nil {
		t.Fatal(err)
	}

	blockHeader, err := client.GetHeadBlock("main")
	if err != nil {
		t.Fatal(err)
	}
	fmt.Printf("%+v\n", blockHeader)
}
