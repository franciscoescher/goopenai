package goopenai

import (
	"context"
	"encoding/json"
)

const COMPLETIONS_URL = "https://api.openai.com/v1/chat/completions"

type CreateCompletionsRequest struct {
	Model            string            `json:"model,omitempty"`
	Messages         []Message         `json:"messages,omitempty"`
	Prompt           StrArray          `json:"prompt,omitempty"`
	Suffix           string            `json:"suffix,omitempty"`
	MaxTokens        int               `json:"max_tokens,omitempty"`
	Temperature      float64           `json:"temperature,omitempty"`
	TopP             float64           `json:"top_p,omitempty"`
	N                int               `json:"n,omitempty"`
	Stream           bool              `json:"stream,omitempty"`
	LogProbs         int               `json:"logprobs,omitempty"`
	Echo             bool              `json:"echo,omitempty"`
	Stop             StrArray          `json:"stop,omitempty"`
	PresencePenalty  float64           `json:"presence_penalty,omitempty"`
	FrequencyPenalty float64           `json:"frequency_penalty,omitempty"`
	BestOf           int               `json:"best_of,omitempty"`
	LogitBias        map[string]string `json:"logit_bias,omitempty"`
	User             string            `json:"user,omitempty"`
}

type CreateCompletionsResponse struct {
	ID      string  `json:"id,omitempty"`
	Object  string  `json:"object,omitempty"`
	Created int     `json:"created,omitempty"`
	Model   string  `json:"model,omitempty"`
	Choices Choices `json:"choices,omitempty"`
	Usage   Usage   `json:"usage,omitempty"`
	Error   Error   `json:"error,omitempty"`
}

type Usage struct {
	PromptTokens     int `json:"prompt_tokens,omitempty"`
	CompletionTokens int `json:"completion_tokens,omitempty"`
	TotalTokens      int `json:"total_tokens,omitempty"`
}

type Choices []struct {
	Text         string      `json:"text,omitempty"`
	Index        int         `json:"index,omitempty"`
	Logprobs     interface{} `json:"logprobs,omitempty"`
	FinishReason string      `json:"finish_reason,omitempty"`
	Message      Message     `json:"message"`
}

func (c *Client) CreateCompletionsRaw(ctx context.Context, r CreateCompletionsRequest) ([]byte, error) {
	return c.Post(ctx, COMPLETIONS_URL, r)
}

func (c *Client) CreateCompletions(ctx context.Context, r CreateCompletionsRequest) (response CreateCompletionsResponse, err error) {
	raw, err := c.CreateCompletionsRaw(ctx, r)
	if err != nil {
		return response, err
	}

	err = json.Unmarshal(raw, &response)
	return response, err
}
