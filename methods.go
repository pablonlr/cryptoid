package cryptoid

import (
	"encoding/json"
	"strconv"
)

type ListUnspentResponse struct {
	UnspentOutputs []UTXOResponse `json:"unspent_outputs"`
}

type UTXOResponse struct {
	TxHash        string  `json:"tx_hash"`
	TxOutputN     int     `json:"tx_ouput_n"`
	Value         float64 `json:"value"`
	Confirmations int     `json:"confirmations"`
	Script        string  `json:"script"`
	Addr          string  `json:"addr"`
}

func (client *CryptoIDClient) MNCount(coin string) (int, error) {
	m := make(map[string]string)
	m["q"] = "masternodecount"
	resp, err := client.GetRequest(coin, m)
	if err != nil {
		return 0, err
	}
	i, err := strconv.Atoi(string(resp))
	if err != nil {
		return 0, err
	}
	return i, nil
}

func (client *CryptoIDClient) BlockCount(coin string) (int, error) {
	m := make(map[string]string)
	m["q"] = "getblockcount"
	resp, err := client.GetRequest(coin, m)
	if err != nil {
		return 0, err
	}
	i, err := strconv.Atoi(string(resp))
	if err != nil {
		return 0, err
	}
	return i, nil
}

func (client *CryptoIDClient) ListUnspent(coin string, address string) (*ListUnspentResponse, error) {
	m := make(map[string]string)
	m["q"] = "unspent"
	m["active"] = address
	resp, err := client.GetRequest(coin, m)
	if err != nil {
		return nil, err
	}
	utxos := &ListUnspentResponse{}
	err = json.Unmarshal(resp, utxos)
	if err != nil {
		return nil, err
	}
	return utxos, nil
}

func (client *CryptoIDClient) GetBalancesForAddressSlice(coin string, addresses []string) (map[string]float64, error) {
	m := make(map[string]string)
	m["q"] = "getbalances"

	resp, err := client.PostRequest(coin, m, addresses)
	if err != nil {
		return nil, err
	}
	balances := make(map[string]float64)
	err = json.Unmarshal(resp, &balances)
	if err != nil {
		return nil, err
	}
	return balances, nil
}
