package provider

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
)

type Client struct {
	baseURL    string
	httpClient *http.Client
	apiKey     string
}

func NewClient(baseURL string, apiKey string) *Client {
	return &Client{
		baseURL:    baseURL,
		httpClient: &http.Client{},
		apiKey:     apiKey,
	}
}

func (c *Client) GetData(endpoint string) (json.RawMessage, error) {
	req, err := http.NewRequest("GET", c.baseURL+endpoint, nil)
	if err != nil {
		return nil, err
	}

	req.Header.Set("x-api-key", c.apiKey)
	resp, err := c.httpClient.Do(req)
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
