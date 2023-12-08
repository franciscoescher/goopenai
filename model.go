package goopenai

import (
	"context"
	"encoding/json"
)

type RetrieveModelResponse struct {
	ID          string   `json:"id,omitempty"`
	Object      string   `json:"object,omitempty"`
	OwnedBy     string   `json:"owned_by,omitempty"`
	Permissions []string `json:"permissions,omitempty"`
}

func (c *Client) RetrieveModelRaw(ctx context.Context, id string) ([]byte, error) {
	return c.Get(ctx, c.apiBase+modelUrl+id, nil)
}

func (c *Client) RetrieveModel(ctx context.Context, id string) (response *RetrieveModelResponse, err error) {
	raw, err := c.RetrieveModelRaw(ctx, id)
	if err != nil {
		return nil, err
	}

	err = json.Unmarshal(raw, &response)
	return response, err
}
