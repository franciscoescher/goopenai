package goopenai

import (
	"context"
	"encoding/json"
)

const MODELS_URL = "https://api.openai.com/v1/models"

type ListModelsResponse struct {
	Data   Data   `json:"data,omitempty"`
	Object string `json:"object,omitempty"`
	Error  Error  `json:"error,omitempty"`
}

func (c *Client) ListModelsRaw(ctx context.Context) ([]byte, error) {
	return c.Get(ctx, MODELS_URL, nil)
}

func (c *Client) ListModels(ctx context.Context) (response ListModelsResponse, err error) {
	raw, err := c.ListModelsRaw(ctx)
	if err != nil {
		return response, err
	}

	err = json.Unmarshal(raw, &response)
	return response, err
}
