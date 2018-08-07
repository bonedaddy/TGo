package tgo

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
)

func (rpc *RPC) GetChainID(chainAlias string) error {
	url := fmt.Sprintf("%s/chains/%s/chain_id", rpc.URL, chainAlias)
	resp, err := rpc.Client.Get(url)
	if err != nil {
		return err
	}
	respBytes, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return err
	}
	fmt.Println(string(respBytes))
	return nil
}

func (rpc *RPC) GetHeadBlock(chainAlias string) (map[string]interface{}, error) {
	url := fmt.Sprintf("%s/chains/%s/blocks/head", rpc.URL, chainAlias)
	resp, err := rpc.Client.Get(url)
	if err != nil {
		return nil, err
	}
	respBytes, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	m := make(map[string]interface{})

	err = json.Unmarshal(respBytes, &m)
	if err != nil {
		return nil, err
	}
	fmt.Printf("%+v\n", m["hash"])
	return m, nil
}
