package goopenai

import (
	"context"
	"encoding/json"
)

type CreateChatCompletionsRequest struct {
	Model            string               `json:"model,omitempty"`
	Messages         []Message            `json:"messages,omitempty"`
	Functions        []CompletionFunciton `json:"functions,omitempty"`
	FunctionCall     *string              `json:"function_call,omitempty"`
	Temperature      float64              `json:"temperature,omitempty"`
	TopP             float64              `json:"top_p,omitempty"`
	N                int                  `json:"n,omitempty"`
	Stream           bool                 `json:"stream,omitempty"`
	Stop             StrArray             `json:"stop,omitempty"`
	MaxTokens        int                  `json:"max_tokens,omitempty"`
	PresencePenalty  float64              `json:"presence_penalty,omitempty"`
	FrequencyPenalty float64              `json:"frequency_penalty,omitempty"`
	LogitBias        map[string]string    `json:"logit_bias,omitempty"`
	User             string               `json:"user,omitempty"`
	ResponseFormat   ResponseFormat       `json:"response_format,omitempty"`
	Seed             int                  `json:"seed,omitempty"`
}

type ResponseFormat struct {
	Type string `json:"type"`
}

type CompletionFunciton struct {
	Name        string `json:"name,omitempty"`
	Description string `json:"description,omitempty"`
	Parameters  []byte `json:"parameters,omitempty"`
}

type CreateChatCompletionsResponse struct {
	ID                string                        `json:"id,omitempty"`
	Object            string                        `json:"object,omitempty"`
	Created           int                           `json:"created,omitempty"`
	Model             string                        `json:"model,omitempty"`
	Choices           []CreateChatCompletionsChoice `json:"choices,omitempty"`
	Usage             CreateChatCompletionsUsave    `json:"usage,omitempty"`
	SystemFingerprint string                        `json:"system_fingerprint,omitempty"`
}

type CreateChatCompletionsChoice struct {
	Index        int      `json:"index,omitempty"`
	Message      *Message `json:"message,omitempty"`
	Delta        *Message `json:"delta,omitempty"`
	FinishReason string   `json:"finish_reason,omitempty"`
}

type CreateChatCompletionsUsave struct {
	PromptTokens     int `json:"prompt_tokens,omitempty"`
	CompletionTokens int `json:"completion_tokens,omitempty"`
	TotalTokens      int `json:"total_tokens,omitempty"`
}

func (c *Client) CreateChatCompletionsRaw(ctx context.Context, r *CreateChatCompletionsRequest) ([]byte, error) {
	if r.ResponseFormat.Type == "" {
		r.ResponseFormat.Type = "text"
	}
	return c.Post(ctx, completionsUrl, r)
}

func (c *Client) CreateChatCompletions(ctx context.Context, r *CreateChatCompletionsRequest) (response *CreateChatCompletionsResponse, err error) {
	raw, err := c.CreateChatCompletionsRaw(ctx, r)
	if err != nil {
		return nil, err
	}

	err = json.Unmarshal(raw, &response)
	return response, err
}
