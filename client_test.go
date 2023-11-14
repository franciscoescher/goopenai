package goopenai

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/require"
)

func Test_Call(t *testing.T) {
	ctx := context.Background()
	url := "/fake-url"
	method := http.MethodPost

	b := map[string]string{"name": "test"}
	rJson, err := json.Marshal(b)
	require.NoError(t, err)

	mocked := getMockedClient(getMockedClientParams{
		T:       t,
		Context: ctx,
		URL:     url,
		Method:  method,
		Body:    bytes.NewReader(rJson),
	})
	defer mocked.Controller.Finish()

	expected := &http.Response{Status: "200 OK"}
	mocked.HttpClient.EXPECT().Do(newRequestMatcher(mocked.Request)).Return(expected, nil)

	response, err := mocked.Client.Call(ctx, method, url, bytes.NewReader(rJson))

	// asserts
	require.NoError(t, err)
	require.Equal(t, expected, response)
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

// asserts interface
var _ gomock.Matcher = (*requestMatcher)(nil)

func newRequestMatcher(req *http.Request) gomock.Matcher {
	return &requestMatcher{req: req}
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

	buf1, _ := io.ReadAll(m.req.Body)
	buf2, _ := io.ReadAll(parsed.Body)
	return bytes.Equal(buf1, buf2)
}

func (m requestMatcher) String() string {
	bodyString, _ := io.ReadAll(m.req.Body)
	fields := map[string]string{
		"\nMethod":                      m.req.Method,
		"\nURL":                         m.req.URL.String(),
		"\nBody":                        string(bodyString),
		"\nHeader[Authorization]":       m.req.Header.Get("Authorization"),
		"\nHeader[Content-Type]":        m.req.Header.Get("Content-Type"),
		"\nHeader[OpenAI-Organization]": m.req.Header.Get("OpenAI-Organization"),
	}
	return fmt.Sprintf("is an http request matching the data: %v", fields)
}
