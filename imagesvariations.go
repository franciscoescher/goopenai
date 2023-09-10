package goopenai

import (
	"context"
	"encoding/json"
)

const IMAGES_VARIATIONS_URL = "https://api.openai.com/v1/images/variations"

type CreateImagesVariationsRequest struct {
	Image          string `json:"image,omitempty"`
	N              int    `json:"n,omitempty"`
	Size           string `json:"size,omitempty"`
	ResponseFormat string `json:"response_format,omitempty"`
	User           string `json:"user,omitempty"`
}

func (c *Client) CreateImagesVariationsRaw(ctx context.Context, r CreateImagesVariationsRequest) ([]byte, error) {
	return c.Post(ctx, IMAGES_VARIATIONS_URL, r)
}

func (c *Client) CreateImagesVariations(ctx context.Context, r CreateImagesVariationsRequest) (response CreateImagesVariationsResponse, err error) {
	raw, err := c.CreateImagesVariationsRaw(ctx, r)
	if err != nil {
		return response, err
	}

	err = json.Unmarshal(raw, &response)
	return response, err
}

type CreateImagesVariationsResponse struct {
	Created int `json:"created,omitempty"`
	Data    []struct {
		URL string `json:"url,omitempty"`
	} `json:"data,omitempty"`

	Error *Error `json:"error,omitempty"`
}
