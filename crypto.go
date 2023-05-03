package cryptoid

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"net/http"
	"net/url"
)

type CryptoIDClient struct {
	*http.Client
	BaseURL string
	KEY     string
}

const BASE_URL = "http://chainz.cryptoid.info/"

// The full route of the request is http://chainz.cryptoid.info/{coin}/api.dws
const adEndpoint = "/api.dws"

func NewCryptoIDClient(key string) *CryptoIDClient {
	return &CryptoIDClient{
		Client:  &http.Client{},
		BaseURL: BASE_URL,
		KEY:     key,
	}
}

func (client *CryptoIDClient) GetRequest(coin string, params map[string]string) ([]byte, error) {
	ur := client.BaseURL + coin + adEndpoint
	req, err := http.NewRequest("GET", ur, nil)
	if err != nil {
		return nil, err
	}
	queries := url.Values{}
	for x := range params {
		queries.Add(x, params[x])
	}
	queries.Add("key", client.KEY)
	req.URL.RawQuery = queries.Encode()
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	return body, nil

}

func (client *CryptoIDClient) PostRequest(coin string, params map[string]string, reqBody interface{}) ([]byte, error) {
	ur := client.BaseURL + coin + adEndpoint
	jsonData, err := json.Marshal(reqBody)
	if err != nil {
		return nil, err
	}
	req, err := http.NewRequest("POST", ur, bytes.NewBuffer(jsonData))
	if err != nil {
		return nil, err
	}
	queries := url.Values{}
	for x := range params {
		queries.Add(x, params[x])
	}
	queries.Add("key", client.KEY)
	req.URL.RawQuery = queries.Encode()

	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	return body, nil
}
