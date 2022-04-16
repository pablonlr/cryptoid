package cryptoid

import (
	"io/ioutil"
	"net/http"
	"net/url"
)

type CryptoIDClient struct {
	*http.Client
	BaseURL    string
	KEY        string
	adEndpoint string
}

const BASE_URL = "http://chainz.cryptoid.info/"

func NewCryptoIDClient(key string) *CryptoIDClient {
	return &CryptoIDClient{
		Client:  &http.Client{},
		BaseURL: BASE_URL,
		KEY:     key,

		//The full route of the request is http://chainz.cryptoid.info/{coin}/api.dws
		adEndpoint: "/api.dws",
	}
}

func (client *CryptoIDClient) GetRequest(coin string, params map[string]string) ([]byte, error) {
	ur := client.BaseURL + coin + client.adEndpoint
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
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	return body, nil

}
