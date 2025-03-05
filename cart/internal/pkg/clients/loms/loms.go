package loms

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
)

type Client struct {
	name string
	path string
}

type StockRequest struct {
	SKU uint32 `json:"sku,omitempty"`
}

type StockResponse struct {
	Count uint64 `json:"count,omitempty"`
}

func (c Client) GetStocks(ctx context.Context, sku uint32) (uint64, error) {
	request := StockRequest{
		SKU: sku,
	}

	data, err := json.Marshal(request)
	if err != nil {
		return 0, fmt.Errorf("error marshalling request: %v", err)
	}

	httpRequest, err := http.NewRequestWithContext(ctx, http.MethodPost, c.path, bytes.NewBuffer(data))
	if err != nil {
		return 0, fmt.Errorf("error creating http request: %v", err)
	}

	httpResponse, err := http.DefaultClient.Do(httpRequest)
	if err != nil {
		return 0, fmt.Errorf("error making http request: %v", err)
	}

	defer func() {
		_ = httpRequest.Body.Close()
	}()

	if httpResponse.StatusCode != http.StatusOK {
		return 0, fmt.Errorf("error getting stocks: status %d", httpResponse.StatusCode)
	}

	response := &StockResponse{}
	err = json.NewDecoder(httpResponse.Body).Decode(response)
	if err != nil {
		return 0, fmt.Errorf("error parsing http response: %v", err)
	}
	return response.Count, nil
}

func New(name string, basePath string) (*Client, error) {
	const handlerName = "stocks"
	path, err := url.JoinPath(basePath, handlerName)
	if err != nil {
		return nil, fmt.Errorf("%s: incorrect base path: %w", name, err)
	}

	return &Client{
		name: name,
		path: path,
	}, nil
}
