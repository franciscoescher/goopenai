package goopenai

import (
	"context"
	"encoding/json"
)

const TRANSLATIONS_URL = "https://api.openai.com/v1/audio/translations"

type CreateTranslationsRequest struct {
	File           string  `json:"file,omitempty"`
	Model          string  `json:"model,omitempty"`
	Prompt         string  `json:"prompt,omitempty"`
	ResponseFormat string  `json:"response_format,omitempty"`
	Temperature    float64 `json:"temperature,omitempty"`
}

func (c *Client) CreateTranslationsRaw(ctx context.Context, r CreateTranslationsRequest) ([]byte, error) {
	return c.Post(ctx, TRANSLATIONS_URL, r)
}

func (c *Client) CreateTranslations(ctx context.Context, r CreateTranslationsRequest) (response CreateTranslationsResponse, err error) {
	raw, err := c.CreateTranslationsRaw(ctx, r)
	if err != nil {
		return response, err
	}

	err = json.Unmarshal(raw, &response)
	return response, err
}

type CreateTranslationsResponse struct {
	Text string `json:"text,omitempty"`

	Error Error `json:"error,omitempty"`
}
