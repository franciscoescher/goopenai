package goopenai

import (
	"context"
	"encoding/json"
)

type CreateEmbeddingsRequest struct {
	Model          string   `json:"model,omitempty"`
	Input          StrArray `json:"input,omitempty"`
	User           string   `json:"user,omitempty"`
	EncodingFormat string   `json:"encoding_format,omitempty"`
}

type CreateEmbeddingsResponse struct {
	Object string                 `json:"object,omitempty"`
	Data   []CreateEmbeddingsData `json:"data,omitempty"`
	Model  string                 `json:"model,omitempty"`
	Usage  CreateEmbeddingsUsage  `json:"usage,omitempty"`
	Error  *Error                 `json:"error"`
}

type CreateEmbeddingsData struct {
	Object    string    `json:"object,omitempty"`
	Embedding []float64 `json:"embedding,omitempty"`
	Index     int       `json:"index,omitempty"`
}

type CreateEmbeddingsUsage struct {
	PromptTokens int `json:"prompt_tokens,omitempty"`
	TotalTokens  int `json:"total_tokens,omitempty"`
}

func (c *Client) CreateEmbeddingsRaw(ctx context.Context, r *CreateEmbeddingsRequest) ([]byte, error) {
	return c.Post(ctx, c.apiBase+embeddingsUrl, r)
}

func (c *Client) CreateEmbeddings(ctx context.Context, r *CreateEmbeddingsRequest) (response *CreateEmbeddingsResponse, err error) {
	raw, err := c.CreateEmbeddingsRaw(ctx, r)
	if err != nil {
		return nil, err
	}

	err = json.Unmarshal(raw, &response)
	return response, err
}
