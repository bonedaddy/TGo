package tgo

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
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

// GetConnections calls GET /network/connections
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

// GetPeerID calls GET /network/connections/<peer_id>
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

// RemovePeers can be used to remove multiple peers at once
// Calls DELETE /network/connections/<peer_id>
func (rpc *RPC) RemovePeers(peers map[string]bool) ([]string, error) {
	processedPeers := []string{}
	for k, v := range peers {
		err := rpc.RemovePeer(k, v)
		if err != nil {
			return processedPeers, err
		}
		processedPeers = append(processedPeers, k)
	}
	return processedPeers, nil
}

// RemovePeer calls DELETE /network/connections/<peer_id>
func (rpc *RPC) RemovePeer(peerID string, wait bool) error {
	url := fmt.Sprintf("%s/network/connections/%s", rpc.URL, peerID)
	if wait {
		url = fmt.Sprintf("%s?wait", url)
	}
	req, err := http.NewRequest("DELETE", url, nil)
	if err != nil {
		return err
	}
	resp, err := rpc.Client.Do(req)
	defer resp.Body.Close()
	if resp.Status != "200 OK" {
		return fmt.Errorf("expected status '200 OK' got %s", resp.Status)
	}
	return nil
}

// ClearGreylist calls GET /network/greylist/clear
func (rpc *RPC) ClearGreylist() error {
	resp, err := rpc.Client.Get(fmt.Sprintf("%s/network/greylist/clear", rpc.URL))
	if err != nil {
		return err
	}
	defer resp.Body.Close()
	if resp.Status != "200 OK" {
		return fmt.Errorf("expected status '200 OK' got %s", resp.Status)
	}
	return nil
}
