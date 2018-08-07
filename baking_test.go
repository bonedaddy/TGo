package tgo_test

import (
	"testing"
	"time"

	tgo "github.com/postables/TGo"
)

func TestGetBakingRightsForDelegateAtCycle(t *testing.T) {
	client := tgo.GenerateClient(tgo.RpcURL, time.Minute)

	err := client.GetBakingRightsForDelegateAtCycle("tz1bhL4zwmLJvHJK5ejDDKdeatpqorvJdc2s", "13", "2")
	if err != nil {
		t.Fatal(err)
	}

}
