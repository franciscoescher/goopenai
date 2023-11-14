package goopenai

import (
	"context"
	"encoding/json"
)

type ListModelsResponse struct {
	Data   []ListModelsData `json:"data,omitempty"`
	Object string           `json:"object,omitempty"`
}

type ListModelsData struct {
	ID          string   `json:"id,omitempty"`
	Object      string   `json:"object,omitempty"`
	OwnedBy     string   `json:"owned_by,omitempty"`
	Permissions []string `json:"permissions,omitempty"`
}

func (c *Client) ListModelsRaw(ctx context.Context) ([]byte, error) {
	return c.Get(ctx, modelsUrl, nil)
}

func (c *Client) ListModels(ctx context.Context) (response *ListModelsResponse, err error) {
	raw, err := c.ListModelsRaw(ctx)
	if err != nil {
		return nil, err
	}

	err = json.Unmarshal(raw, &response)
	return response, err
}
