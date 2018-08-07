package tgo

import (
	"fmt"
	"io/ioutil"
)

// GetBakingRightsForDelegateAtCycle is used to get the baking rights for a particular delegate with a customizable priority
func (rpc *RPC) GetBakingRightsForDelegateAtCycle(delegate, cycle, priority string) error {
	if priority == "" {
		priority = "2"
	}
	url := fmt.Sprintf("%s/chains/main/blocks/head/helpers/baking_rights?cycle=%s&delegate=%s&max_priority=%s", rpc.URL, cycle, delegate, priority)
	resp, err := rpc.Client.Get(url)
	if err != nil {
		return err
	}

	respBytes, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return err
	}

	fmt.Printf("Baking rights for delegate %s at cycle %s\n%v", delegate, cycle, string(respBytes))
	return nil
}
