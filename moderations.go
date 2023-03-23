package goopenai

import (
	"context"
	"encoding/json"
)

const MODERATIONS_URL = "https://api.openai.com/v1/moderations"

type CreateModerationsResponse struct {
	ID      string  `json:"id,omitempty"`
	Model   string  `json:"model,omitempty"`
	Results Results `json:"results,omitempty"`
	Error   Error   `json:"error,omitempty"`
}

type Results []struct {
	Categories     Categories     `json:"categories,omitempty"`
	CategoryScores CategoryScores `json:"category_scores,omitempty"`
	Flagged        bool           `json:"flagged,omitempty"`
}

type CategoryScores struct {
	Hate            float64 `json:"hate,omitempty"`
	HateThreatening float64 `json:"hate/threatening,omitempty"`
	SelfHarm        float64 `json:"self-harm,omitempty"`
	Sexual          float64 `json:"sexual,omitempty"`
	SexualMinors    float64 `json:"sexual/minors,omitempty"`
	Violence        float64 `json:"violence,omitempty"`
	ViolenceGraphic float64 `json:"violence/graphic,omitempty"`
}

type Categories struct {
	Hate            bool `json:"hate,omitempty"`
	HateThreatening bool `json:"hate/threatening,omitempty"`
	SelfHarm        bool `json:"self-harm,omitempty"`
	Sexual          bool `json:"sexual,omitempty"`
	SexualMinors    bool `json:"sexual/minors,omitempty"`
	Violence        bool `json:"violence,omitempty"`
	ViolenceGraphic bool `json:"violence/graphic,omitempty"`
}

type CreateModerationsRequest struct {
	Input StrArray `json:"input,omitempty"`
	Model string   `json:"model,omitempty"`
}

func (c *Client) CreateModerationsRaw(ctx context.Context, r CreateModerationsRequest) ([]byte, error) {
	return c.Post(ctx, MODERATIONS_URL, r)
}

func (c *Client) CreateModerations(ctx context.Context, r CreateModerationsRequest) (response CreateModerationsResponse, err error) {
	raw, err := c.CreateModerationsRaw(ctx, r)
	if err != nil {
		return response, err
	}

	err = json.Unmarshal(raw, &response)
	return response, err
}
