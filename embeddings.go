package goopenai

import (
	"context"
	"encoding/json"
)

const EMBEDDINGS_URL = "https://api.openai.com/v1/embeddings"

type CreateEmbeddingsRequest struct {
	Model string   `json:"model,omitempty"`
	Input StrArray `json:"input,omitempty"`
	User  string   `json:"user,omitempty"`
}

type CreateEmbeddingsResponse struct {
	Object string `json:"object,omitempty"`
	Data   Data   `json:"data,omitempty"`
	Model  string `json:"model,omitempty"`
	Usage  Usage  `json:"usage,omitempty"`
	Error  Error  `json:"error"`
}

type Data []struct {
	Object      string    `json:"object,omitempty"`
	Embedding   []float64 `json:"embedding,omitempty"`
	Index       int       `json:"index,omitempty"`
	URL         string    `json:"url,omitempty"`
	ID          string    `json:"id,omitempty"`
	OwnedBy     string    `json:"owned_by,omitempty"`
	Permissions []string  `json:"permissions,omitempty"`
}

func (c *Client) CreateEmbeddingsRaw(ctx context.Context, r CreateEmbeddingsRequest) ([]byte, error) {
	return c.Post(ctx, EMBEDDINGS_URL, r)
}

func (c *Client) CreateEmbeddings(ctx context.Context, r CreateEmbeddingsRequest) (response CreateEmbeddingsResponse, err error) {
	raw, err := c.CreateEmbeddingsRaw(ctx, r)
	if err != nil {
		return response, err
	}

	err = json.Unmarshal(raw, &response)
	return response, err
}
