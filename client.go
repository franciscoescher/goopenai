package goopenai

import (
	"bytes"
	"context"
	"encoding/json"
	"io"
	"net/http"
	"strings"

	"github.com/google/go-querystring/query"
)

type Client struct {
	apiKey       string
	Organization string
}

// Asserts client follows client interface
var _ ClientInterface = (*Client)(nil)

// NewClient creates a new client
func NewClient(apiKey string, organization string) *Client {
	return &Client{
		apiKey:       apiKey,
		Organization: organization,
	}
}

// Post makes a post request
func (c *Client) Post(ctx context.Context, url string, input any) (response []byte, err error) {
	response = make([]byte, 0)

	rJson, err := json.Marshal(input)
	if err != nil {
		return response, err
	}

	resp, err := c.Call(ctx, http.MethodPost, url, bytes.NewReader(rJson))
	if err != nil {
		return response, err
	}
	defer resp.Body.Close()

	response, err = io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	var errorResponse struct {
		Error *Error `json:"error,omitempty"`
	}
	err = json.Unmarshal(response, &errorResponse)
	if errorResponse.Error != nil {
		return nil, errorResponse.Error
	}

	return response, err
}

// Get makes a get request
func (c *Client) Get(ctx context.Context, url string, input any) (response []byte, err error) {
	if input != nil {
		vals, _ := query.Values(input)
		query := vals.Encode()

		if query != "" {
			sb := strings.Builder{}
			sb.WriteString(url)
			sb.WriteString("?")
			sb.WriteString(query)
			url = sb.String()
		}
	}

	resp, err := c.Call(ctx, http.MethodGet, url, nil)
	if err != nil {
		return response, err
	}
	defer resp.Body.Close()

	response, err = io.ReadAll(resp.Body)
	return response, err
}

// Call makes a request
func (c *Client) Call(ctx context.Context, method string, url string, body io.Reader) (response *http.Response, err error) {
	req, err := http.NewRequestWithContext(ctx, method, url, body)
	if err != nil {
		return response, err
	}

	sb := strings.Builder{}
	sb.WriteString("Bearer ")
	sb.WriteString(c.apiKey)
	authHeader := sb.String()

	req.Header.Add("Authorization", authHeader)
	req.Header.Add("Content-Type", "application/json")
	if c.Organization != "" {
		req.Header.Add("OpenAI-Organization", c.Organization)
	}

	client := &http.Client{}
	resp, err := client.Do(req)
	return resp, err
}
