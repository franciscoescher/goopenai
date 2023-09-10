package goopenai

import (
	"context"
	"encoding/json"
)

type CreateTranslationsRequest struct {
	File           string  `json:"file,omitempty"`
	Model          string  `json:"model,omitempty"`
	Prompt         string  `json:"prompt,omitempty"`
	ResponseFormat string  `json:"response_format,omitempty"`
	Temperature    float64 `json:"temperature,omitempty"`
}

type CreateTranslationsResponse struct {
	Text  string `json:"text,omitempty"`
	Error *Error `json:"error,omitempty"`
}

func (c *Client) CreateTranslationsRaw(ctx context.Context, r CreateTranslationsRequest) ([]byte, error) {
	return c.Post(ctx, translationsUrl, r)
}

func (c *Client) CreateTranslations(ctx context.Context, r CreateTranslationsRequest) (response CreateTranslationsResponse, err error) {
	raw, err := c.CreateTranslationsRaw(ctx, r)
	if err != nil {
		return response, err
	}

	err = json.Unmarshal(raw, &response)
	return response, err
}
