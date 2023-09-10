package goopenai

import (
	"context"
	"encoding/json"
)

const IMAGES_EDITS_URL = "https://api.openai.com/v1/images/edits"

type CreateImagesEditsRequest struct {
	Image          string `json:"image,omitempty"`
	Mask           string `json:"mask,omitempty"`
	Prompt         string `json:"prompt,omitempty"`
	N              int    `json:"n,omitempty"`
	Size           string `json:"size,omitempty"`
	ResponseFormat string `json:"response_format,omitempty"`
	User           string `json:"user,omitempty"`
}

func (c *Client) CreateImagesEditsRaw(ctx context.Context, r CreateImagesEditsRequest) ([]byte, error) {
	return c.Post(ctx, IMAGES_EDITS_URL, r)
}

func (c *Client) CreateImagesEdits(ctx context.Context, r CreateImagesEditsRequest) (response CreateImagesEditsResponse, err error) {
	raw, err := c.CreateImagesEditsRaw(ctx, r)
	if err != nil {
		return response, err
	}

	err = json.Unmarshal(raw, &response)
	return response, err
}

type CreateImagesEditsResponse struct {
	Created int `json:"created,omitempty"`
	Data    []struct {
		URL string `json:"url,omitempty"`
	} `json:"data,omitempty"`

	Error *Error `json:"error,omitempty"`
}
