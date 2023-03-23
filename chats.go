package goopenai

import (
	"context"
	"encoding/json"
)

const CHATS_URL = "https://api.openai.com/v1/chat/completions"

type CreateChatsRequest struct {
	Model            string            `json:"model,omitempty"`
	Messages         []Message         `json:"messages,omitempty"`
	Temperature      float64           `json:"temperature,omitempty"`
	TopP             float64           `json:"top_p,omitempty"`
	N                int               `json:"n,omitempty"`
	Stream           bool              `json:"stream,omitempty"`
	Stop             StrArray          `json:"stop,omitempty"`
	MaxTokens        int               `json:"max_tokens,omitempty"`
	PresencePenalty  float64           `json:"presence_penalty,omitempty"`
	FrequencyPenalty float64           `json:"frequency_penalty,omitempty"`
	LogitBias        map[string]string `json:"logit_bias,omitempty"`
	User             string            `json:"user,omitempty"`
}

type CreateChatsResponse struct {
	ID      string  `json:"id,omitempty"`
	Object  string  `json:"object,omitempty"`
	Created int     `json:"created,omitempty"`
	Choices Choices `json:"choices,omitempty"`
	Usage   Usage   `json:"usage,omitempty"`
	Error   Error   `json:"error,omitempty"`
}

type Message struct {
	Role    string `json:"role,omitempty"`
	Content string `json:"content,omitempty"`
}

func (c *Client) CreateChatsRaw(ctx context.Context, r CreateChatsRequest) ([]byte, error) {
	return c.Post(ctx, CHATS_URL, r)
}

func (c *Client) CreateChats(ctx context.Context, r CreateChatsRequest) (response CreateChatsResponse, err error) {
	raw, err := c.CreateChatsRaw(ctx, r)
	if err != nil {
		return response, err
	}

	err = json.Unmarshal(raw, &response)
	return response, err
}
