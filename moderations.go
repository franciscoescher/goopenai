package goopenai

import (
	"encoding/json"
)

const MODERATIONS_URL = "https://api.openai.com/v1/moderations"

type CreateModerationsRequest struct {
	Input StrArray `json:"input,omitempty"`
	Model string   `json:"model,omitempty"`
}

func (c *Client) CreateModerationsRaw(r CreateModerationsRequest) ([]byte, error) {
	return c.Post(MODERATIONS_URL, r)
}

func (c *Client) CreateModerations(r CreateModerationsRequest) (response CreateModerationsResponse, err error) {
	raw, err := c.CreateModerationsRaw(r)
	if err != nil {
		return response, err
	}

	err = json.Unmarshal(raw, &response)
	return response, err
}

type CreateModerationsResponse struct {
	ID      string `json:"id,omitempty"`
	Model   string `json:"model,omitempty"`
	Results []struct {
		Categories struct {
			Hate            bool `json:"hate,omitempty"`
			HateThreatening bool `json:"hate/threatening,omitempty"`
			SelfHarm        bool `json:"self-harm,omitempty"`
			Sexual          bool `json:"sexual,omitempty"`
			SexualMinors    bool `json:"sexual/minors,omitempty"`
			Violence        bool `json:"violence,omitempty"`
			ViolenceGraphic bool `json:"violence/graphic,omitempty"`
		} `json:"categories,omitempty"`
		CategoryScores struct {
			Hate            float64 `json:"hate,omitempty"`
			HateThreatening float64 `json:"hate/threatening,omitempty"`
			SelfHarm        float64 `json:"self-harm,omitempty"`
			Sexual          float64 `json:"sexual,omitempty"`
			SexualMinors    float64 `json:"sexual/minors,omitempty"`
			Violence        float64 `json:"violence,omitempty"`
			ViolenceGraphic float64 `json:"violence/graphic,omitempty"`
		} `json:"category_scores,omitempty"`
		Flagged bool `json:"flagged,omitempty"`
	} `json:"results,omitempty"`

	Error Error `json:"error,omitempty"`
}
