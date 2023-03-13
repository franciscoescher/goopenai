package goopenai

import (
	"encoding/json"
)

const EMBEDDINGS_URL = "https://api.openai.com/v1/embeddings"

type CreateEmbeddingsRequest struct {
	Model string   `json:"model,omitempty"`
	Input StrArray `json:"input,omitempty"`
	User  string   `json:"user,omitempty"`
}

func (c *Client) CreateEmbeddingsRaw(r CreateEmbeddingsRequest) ([]byte, error) {
	return c.Post(EMBEDDINGS_URL, r)
}

func (c *Client) CreateEmbeddings(r CreateEmbeddingsRequest) (response CreateEmbeddingsResponse, err error) {
	raw, err := c.CreateEmbeddingsRaw(r)
	if err != nil {
		return response, err
	}

	err = json.Unmarshal(raw, &response)
	return response, err
}

type CreateEmbeddingsResponse struct {
	Object string `json:"object,omitempty"`
	Data   []struct {
		Object    string    `json:"object,omitempty"`
		Embedding []float64 `json:"embedding,omitempty"`
		Index     int       `json:"index,omitempty"`
	} `json:"data,omitempty"`
	Model string `json:"model,omitempty"`
	Usage struct {
		PromptTokens int `json:"prompt_tokens,omitempty"`
		TotalTokens  int `json:"total_tokens,omitempty"`
	} `json:"usage,omitempty"`

	Error Error `json:"error"`
}
