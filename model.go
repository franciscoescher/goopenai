package goopenai

import (
	"context"
	"encoding/json"
)

const MODEL_URL = "https://api.openai.com/v1/models/"

func (c *Client) GetModelRaw(ctx context.Context, id string) ([]byte, error) {
	return c.Get(ctx, MODEL_URL+id, nil)
}

func (c *Client) GetModel(ctx context.Context, id string) (response GetModelResponse, err error) {
	raw, err := c.GetModelRaw(ctx, id)
	if err != nil {
		return response, err
	}

	err = json.Unmarshal(raw, &response)
	return response, err
}

type GetModelResponse struct {
	ID          string   `json:"id,omitempty"`
	Object      string   `json:"object,omitempty"`
	OwnedBy     string   `json:"owned_by,omitempty"`
	Permissions []string `json:"permissions,omitempty"`

	Error Error `json:"error,omitempty"`
}
