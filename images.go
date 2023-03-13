package goopenai

import (
	"encoding/json"
)

const IMAGES_URL = "https://api.openai.com/v1/images/generations"

type CreateImagesRequest struct {
	Prompt         string `json:"prompt,omitempty"`
	N              int    `json:"n,omitempty"`
	Size           string `json:"size,omitempty"`
	ResponseFormat string `json:"response_format,omitempty"`
	User           string `json:"user,omitempty"`
}

func (c *Client) CreateImagesRaw(r CreateImagesRequest) ([]byte, error) {
	return c.Post(IMAGES_URL, r)
}

func (c *Client) CreateImages(r CreateImagesRequest) (response CreateImagesResponse, err error) {
	raw, err := c.CreateImagesRaw(r)
	if err != nil {
		return response, err
	}

	err = json.Unmarshal(raw, &response)
	return response, err
}

type CreateImagesResponse struct {
	Created int `json:"created,omitempty"`
	Data    []struct {
		URL string `json:"url,omitempty"`
	} `json:"data,omitempty"`

	Error Error `json:"error,omitempty"`
}
