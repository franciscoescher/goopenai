package goopenai

import (
	"context"
	"encoding/json"
)

type CreateImagesEditsRequest struct {
	Image          string `json:"image,omitempty"`
	Mask           string `json:"mask,omitempty"`
	Prompt         string `json:"prompt,omitempty"`
	N              int    `json:"n,omitempty"`
	Size           string `json:"size,omitempty"`
	ResponseFormat string `json:"response_format,omitempty"`
	User           string `json:"user,omitempty"`
}

type CreateImagesEditsResponse struct {
	Created int                     `json:"created,omitempty"`
	Data    []CreateImagesEditsData `json:"data,omitempty"`
}

type CreateImagesEditsData struct {
	URL string `json:"url,omitempty"`
}

func (c *Client) CreateImagesEditsRaw(ctx context.Context, r *CreateImagesEditsRequest) ([]byte, error) {
	return c.Post(ctx, c.apiBase + imagesEditsUrl, r)
}

func (c *Client) CreateImagesEdits(ctx context.Context, r *CreateImagesEditsRequest) (response *CreateImagesEditsResponse, err error) {
	raw, err := c.CreateImagesEditsRaw(ctx, r)
	if err != nil {
		return nil, err
	}

	err = json.Unmarshal(raw, &response)
	return response, err
}
