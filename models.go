package goopenai

import "encoding/json"

const MODELS_URL = "https://api.openai.com/v1/models"

func (c *Client) GetModelsRaw() ([]byte, error) {
	return c.Get(MODELS_URL, nil)
}

func (c *Client) GetModels() (response GetModelsResponse, err error) {
	raw, err := c.GetModelsRaw()
	if err != nil {
		return response, err
	}

	err = json.Unmarshal(raw, &response)
	return response, err
}

type GetModelsResponse struct {
	Data []struct {
		ID          string   `json:"id,omitempty"`
		Object      string   `json:"object,omitempty"`
		OwnedBy     string   `json:"owned_by,omitempty"`
		Permissions []string `json:"permissions,omitempty"`
	} `json:"data,omitempty"`
	Object string `json:"object,omitempty"`

	Error Error `json:"error,omitempty"`
}
