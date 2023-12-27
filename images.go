package goopenai

import (
	"context"
	"encoding/json"
)

type CreateImagesRequest struct {
	Prompt         string `json:"prompt,omitempty"`
	N              int    `json:"n,omitempty"`
	Size           string `json:"size,omitempty"`
	ResponseFormat string `json:"response_format,omitempty"`
	User           string `json:"user,omitempty"`
}

type CreateImagesResponse struct {
	Created int                `json:"created,omitempty"`
	Data    []CreateImagesData `json:"data,omitempty"`
}

type CreateImagesData struct {
	URL string `json:"url,omitempty"`
}

func (c *Client) CreateImagesRaw(ctx context.Context, r *CreateImagesRequest) ([]byte, error) {
	return c.Post(ctx, c.apiBase + imagesUrl, r)
}

func (c *Client) CreateImages(ctx context.Context, r *CreateImagesRequest) (response *CreateImagesResponse, err error) {
	raw, err := c.CreateImagesRaw(ctx, r)
	if err != nil {
		return nil, err
	}

	err = json.Unmarshal(raw, &response)
	return response, err
}
