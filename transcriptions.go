package goopenai

import (
	"context"
	"encoding/json"
)

const TRANSCRIPTIONS_URL = "https://api.openai.com/v1/audio/transcriptions"

type CreateTranscriptionsRequest struct {
	File           string  `json:"file,omitempty"`
	Model          string  `json:"model,omitempty"`
	Prompt         string  `json:"prompt,omitempty"`
	ResponseFormat string  `json:"response_format,omitempty"`
	Temperature    float64 `json:"temperature,omitempty"`
	Language       string  `json:"language,omitempty"`
}

type CreateTranscriptionsResponse struct {
	Text  string `json:"text,omitempty"`
	Error Error  `json:"error,omitempty"`
}

func (c *Client) CreateTranscriptionsRaw(ctx context.Context, r CreateTranscriptionsRequest) ([]byte, error) {
	return c.Post(ctx, TRANSCRIPTIONS_URL, r)
}

func (c *Client) CreateTranscriptions(ctx context.Context, r CreateTranscriptionsRequest) (response CreateTranscriptionsResponse, err error) {
	raw, err := c.CreateTranscriptionsRaw(ctx, r)
	if err != nil {
		return response, err
	}

	err = json.Unmarshal(raw, &response)
	return response, err
}
