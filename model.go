package goopenai

import "encoding/json"

const MODEL_URL = "https://api.openai.com/v1/models/"

func (c *Client) GetModelRaw(id string) ([]byte, error) {
	return c.Get(MODEL_URL+id, nil)
}

func (c *Client) GetModel(id string) (response GetModelResponse, err error) {
	raw, err := c.GetModelRaw(id)
	if err != nil {
		return response, err
	}

	err = json.Unmarshal(raw, &response)
	return response, err
}

type GetModelResponse struct {
	ID          string   `json:"id,omitempty"`
	Object      string   `json:"object,omitempty"`
	OwnedBy     string   `json:"owned_by,omitempty"`
	Permissions []string `json:"permissions,omitempty"`

	Error Error `json:"error,omitempty"`
}
