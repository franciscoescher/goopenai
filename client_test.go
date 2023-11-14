package goopenai

import (
	"bytes"
	"context"
	"encoding/json"
	"io"
	"net/http"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/require"
)

func TestCall(t *testing.T) {
	t.Run("Call", func(t *testing.T) {
		ctx := context.Background()
		url := "/fake-url"
		method := http.MethodPost

		b := map[string]string{"name": "test"}
		rJson, err := json.Marshal(b)
		require.NoError(t, err)
		body := bytes.NewReader(rJson)

		mocked := getMockedClient(getMockedClientParams{
			T:       t,
			Context: ctx,
			URL:     url,
			Method:  method,
			Body:    body,
		})
		defer mocked.Controller.Finish()

		expected := &http.Response{Status: "200 OK"}
		mocked.HttpClient.EXPECT().Do(requestMatcher{mocked.Request}).Return(expected, nil)

		response, err := mocked.Client.Call(ctx, method, url, body)

		// asserts
		require.NoError(t, err)
		require.Equal(t, expected, response)
	})
}

type getMockedClientParams struct {
	T       *testing.T
	Context context.Context
	URL     string
	Method  string
	Body    io.Reader
}

type getMockedClientResponse struct {
	Controller *gomock.Controller
	Client     *Client
	Request    *http.Request
	HttpClient *MockHttpClient
}

func getMockedClient(p getMockedClientParams) *getMockedClientResponse {
	apiKey := "fake-api-key"
	org := "fake-org"

	// prepares request
	req, err := http.NewRequestWithContext(p.Context, p.Method, p.URL, p.Body)
	require.NoError(p.T, err)
	req.Header.Add("Authorization", "Bearer "+apiKey)
	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("OpenAI-Organization", org)

	// mocks http client
	ctrl := gomock.NewController(p.T)
	defer ctrl.Finish()
	mockHttpClient := NewMockHttpClient(ctrl)

	// calls client
	c := &Client{
		apiKey:       apiKey,
		organization: org,
		client:       mockHttpClient,
	}
	return &getMockedClientResponse{
		ctrl, c, req, mockHttpClient,
	}
}

type requestMatcher struct {
	req *http.Request
}

func (m requestMatcher) Matches(x interface{}) bool {
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

func (m requestMatcher) String() string {
	return "is an http request with the same method, url, required headers and body"
}
