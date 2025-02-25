package product

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
)

type Client struct {
	name string
	path string
}

type GetProductRequest struct {
	Token string `json:"token,omitempty"`
	SKU   uint32 `json:"sku,omitempty"`
}

type GetProductResponse struct {
	Name  string `json:"name,omitempty"`
	Price uint32 `json:"price,omitempty"`
}

type GetProductErrorResponse struct {
	Code    int    `json:"code,omitempty"`
	Message string `json:"message,omitempty"`
}

func (c Client) GetProductInfo(sku uint32, token string) (string, uint32, error) {
	request := GetProductRequest{
		Token: token,
		SKU:   sku,
	}

	data, err := json.Marshal(request)
	if err != nil {
		return "", 0, fmt.Errorf("error marshalling request: %v", err)
	}

	httpRequest, err := http.NewRequest(http.MethodPost, c.path, bytes.NewBuffer(data))
	if err != nil {
		return "", 0, fmt.Errorf("error creating http request: %v", err)
	}

	httpResponse, err := http.DefaultClient.Do(httpRequest)
	if err != nil {
		return "", 0, fmt.Errorf("error making http request: %v", err)
	}

	defer func() {
		_ = httpRequest.Body.Close()
	}()

	if httpResponse.StatusCode != http.StatusOK {
		response := &GetProductErrorResponse{}
		err = json.NewDecoder(httpResponse.Body).Decode(response)
		if err != nil {
			return "", 0, fmt.Errorf("error parsing error response: %v", err)
		}
		return "", 0, fmt.Errorf(response.Message)
	}

	response := &GetProductResponse{}
	err = json.NewDecoder(httpResponse.Body).Decode(response)
	if err != nil {
		return "", 0, fmt.Errorf("error parsing http response: %v", err)
	}
	return response.Name, response.Price, nil
}

func New(name string, basePath string) (*Client, error) {
	const handlerName = "get_product"
	path, err := url.JoinPath(basePath, handlerName)
	if err != nil {
		return nil, fmt.Errorf("%s: incorrect base path: %w", name, err)
	}

	return &Client{
		name: name,
		path: path,
	}, nil
}
