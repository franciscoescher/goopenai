package goopenai

import (
	"bytes"
	"context"
	"encoding/json"
	"net/http"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/require"
)

func TestCall(t *testing.T) {
	t.Run("Call", func(t *testing.T) {
		ctx := context.Background()
		url := "/test/url"
		method := http.MethodPost
		org := "org"
		apiKey := "api-key"

		// prepares request body
		b := map[string]string{"name": "test"}
		rJson, err := json.Marshal(b)
		require.NoError(t, err)
		body := bytes.NewReader(rJson)

		// prepares request
		req, err := http.NewRequestWithContext(ctx, method, url, body)
		require.NoError(t, err)
		req.Header.Add("Authorization", "Bearer "+apiKey)
		req.Header.Add("Content-Type", "application/json")
		req.Header.Add("OpenAI-Organization", org)

		// mocks http client
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()
		mockHttpClient := NewMockHttpClient(ctrl)
		expected := &http.Response{}
		mockHttpClient.EXPECT().Do(gomock.Any()).Return(expected, nil)

		// calls client
		c := &Client{
			apiKey:       apiKey,
			organization: org,
			client:       mockHttpClient,
		}
		response, err := c.Call(ctx, method, url, body)

		// asserts
		require.NoError(t, err)
		require.Equal(t, expected, response)
	})
}
