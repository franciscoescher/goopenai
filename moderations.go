package goopenai

import (
	"context"
	"encoding/json"
)

type CreateModerationsRequest struct {
	Input StrArray `json:"input,omitempty"`
	Model string   `json:"model,omitempty"`
}

type CreateModerationsResponse struct {
	ID      string                    `json:"id,omitempty"`
	Model   string                    `json:"model,omitempty"`
	Results []CreateModerationsResult `json:"results,omitempty"`
}

type CreateModerationsResult struct {
	Categories     CreateModerationsCategories       `json:"categories,omitempty"`
	CategoryScores CreateModerationsCategoriesScores `json:"category_scores,omitempty"`
	Flagged        bool                              `json:"flagged,omitempty"`
}

type CreateModerationsCategories struct {
	Hate            bool `json:"hate,omitempty"`
	HateThreatening bool `json:"hate/threatening,omitempty"`
	SelfHarm        bool `json:"self-harm,omitempty"`
	Sexual          bool `json:"sexual,omitempty"`
	SexualMinors    bool `json:"sexual/minors,omitempty"`
	Violence        bool `json:"violence,omitempty"`
	ViolenceGraphic bool `json:"violence/graphic,omitempty"`
}

type CreateModerationsCategoriesScores struct {
	Hate            float64 `json:"hate,omitempty"`
	HateThreatening float64 `json:"hate/threatening,omitempty"`
	SelfHarm        float64 `json:"self-harm,omitempty"`
	Sexual          float64 `json:"sexual,omitempty"`
	SexualMinors    float64 `json:"sexual/minors,omitempty"`
	Violence        float64 `json:"violence,omitempty"`
	ViolenceGraphic float64 `json:"violence/graphic,omitempty"`
}

func (c *Client) CreateModerationsRaw(ctx context.Context, r CreateModerationsRequest) ([]byte, error) {
	return c.Post(ctx, moderationsUrl, r)
}

func (c *Client) CreateModerations(ctx context.Context, r CreateModerationsRequest) (response CreateModerationsResponse, err error) {
	raw, err := c.CreateModerationsRaw(ctx, r)
	if err != nil {
		return response, err
	}

	err = json.Unmarshal(raw, &response)
	return response, err
}
