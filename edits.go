package goopenai

import (
	"context"
	"encoding/json"
)

type CreateEditsRequest struct {
	Model       string  `json:"model,omitempty"`
	Input       string  `json:"input,omitempty"`
	Instruction string  `json:"instruction,omitempty"`
	N           int     `json:"n,omitempty"`
	Temperature float64 `json:"temperature,omitempty"`
	TopP        float64 `json:"top_p,omitempty"`
}

type CreateEditsResponse struct {
	Object  string              `json:"object,omitempty"`
	Created int                 `json:"created,omitempty"`
	Choices []CreateEditsChoice `json:"choices,omitempty"`
	Usage   CreateEditsUsage    `json:"usage,omitempty"`
}

type CreateEditsChoice struct {
	Text  string `json:"text,omitempty"`
	Index int    `json:"index,omitempty"`
}

type CreateEditsUsage struct {
	PromptTokens     int `json:"prompt_tokens,omitempty"`
	CompletionTokens int `json:"completion_tokens,omitempty"`
	TotalTokens      int `json:"total_tokens,omitempty"`
}

func (c *Client) CreateEditsRaw(ctx context.Context, r *CreateEditsRequest) ([]byte, error) {
	return c.Post(ctx, editsUrl, r)
}

func (c *Client) CreateEdits(ctx context.Context, r *CreateEditsRequest) (response *CreateEditsResponse, err error) {
	raw, err := c.CreateEditsRaw(ctx, r)
	if err != nil {
		return nil, err
	}

	err = json.Unmarshal(raw, &response)
	return response, err
}
