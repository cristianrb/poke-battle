package api

import (
	"encoding/json"
	"io"
	"net/http"
	"time"
)

type HTTPClient interface {
	Do(endpoint string, response any) error
}

type DefaultHTTPClient struct {
	baseURL string
	client  *http.Client
}

func NewHTTPClient(baseURL string) *DefaultHTTPClient {
	return &DefaultHTTPClient{
		baseURL: baseURL,
		client:  &http.Client{Timeout: 10 * time.Second},
	}
}

func (c *DefaultHTTPClient) Do(endpoint string, response any) error {
	req, err := http.NewRequest(http.MethodGet, c.baseURL+endpoint, nil)
	if err != nil {
		return err
	}

	resp, err := c.client.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return err
	}

	return json.Unmarshal(body, &response)
}
