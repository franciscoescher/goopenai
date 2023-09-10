package goopenai

import (
	"context"
	"encoding/json"
)

type CreateImagesVariationsRequest struct {
	Image          string `json:"image,omitempty"`
	N              int    `json:"n,omitempty"`
	Size           string `json:"size,omitempty"`
	ResponseFormat string `json:"response_format,omitempty"`
	User           string `json:"user,omitempty"`
}

type CreateImagesVariationsResponse struct {
	Created int                          `json:"created,omitempty"`
	Data    []CreateImagesVariationsData `json:"data,omitempty"`
}

type CreateImagesVariationsData struct {
	URL string `json:"url,omitempty"`
}

func (c *Client) CreateImagesVariationsRaw(ctx context.Context, r CreateImagesVariationsRequest) ([]byte, error) {
	return c.Post(ctx, imagesVariationsUrl, r)
}

func (c *Client) CreateImagesVariations(ctx context.Context, r CreateImagesVariationsRequest) (response CreateImagesVariationsResponse, err error) {
	raw, err := c.CreateImagesVariationsRaw(ctx, r)
	if err != nil {
		return response, err
	}

	err = json.Unmarshal(raw, &response)
	return response, err
}
