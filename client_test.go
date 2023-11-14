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
		expected := &http.Response{Status: "200 OK"}
		mockHttpClient.EXPECT().Do(RequestMatcher{req}).Return(expected, nil)

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

type RequestMatcher struct {
	req *http.Request
}

func (m RequestMatcher) Matches(x interface{}) bool {
	// check x type
	parsed, ok := x.(*http.Request)
	if !ok {
		return false
	}
	if m.req.Method != parsed.Method {
		return false
	}
	if m.req.URL.String() != parsed.URL.String() {
		return false
	}
	if m.req.Header.Get("Authorization") != parsed.Header.Get("Authorization") {
		return false
	}
	if m.req.Header.Get("Content-Type") != parsed.Header.Get("Content-Type") {
		return false
	}
	if m.req.Header.Get("OpenAI-Organization") != parsed.Header.Get("OpenAI-Organization") {
		return false
	}
	if m.req.Body != parsed.Body {
		return false
	}
	return true
}

func (m RequestMatcher) String() string {
	return "is an http request with the same method, url, required headers and body"
}
