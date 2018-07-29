package tgo

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
)

// ConnectionsResponse holds the response from `GET /network/connections`
type ConnectionsResponse struct {
	Incoming bool   `json:"incoming"`
	PeerID   string `json:"peer_id"`
	IDPoint  struct {
		Address string `json:"addr"`
		Port    int64  `json:"port"`
	} `json:"id_point"`
	RemoteSocketPort int64 `json:"remote_socket_port"`
	Versions         []struct {
		Name  string `json:"name"`
		Major int64  `json:"magor"`
		Minor int64  `json:"miner"`
	} `json:"versions"`
	Private       bool `json:"private"`
	LocalMetadata struct {
		DisableMempool bool `json:"disable_mempool"`
		PrivateNode    bool `json:"private_node"`
	} `json:"local_metadata"`
	RemoteMetadata struct {
		DisableMempool bool `json:"disable_mempool"`
		PrivateNode    bool `json:"private_node"`
	} `json:"remote_metadata"`
}

// GetConnections calls /network/connections
func (rpc *RPC) GetConnections() ([]ConnectionsResponse, error) {
	resp, err := rpc.Client.Get(fmt.Sprintf("%s/network/connections", rpc.URL))
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	respBytes, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	cp := []ConnectionsResponse{}
	err = json.Unmarshal(respBytes, &cp)
	if err != nil {
		return nil, err
	}
	return cp, nil
}

func (rpc *RPC) GetPeerID(peerID string) (ConnectionsResponse, error) {
	resp, err := rpc.Client.Get(fmt.Sprintf("%s/network/connections/%s", rpc.URL, peerID))
	if err != nil {
		return ConnectionsResponse{}, err
	}
	defer resp.Body.Close()
	respBytes, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return ConnectionsResponse{}, err
	}
	cp := ConnectionsResponse{}
	err = json.Unmarshal(respBytes, &cp)
	if err != nil {
		return ConnectionsResponse{}, err
	}
	return cp, nil
}
