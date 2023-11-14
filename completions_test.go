package goopenai

import (
	"bytes"
	"context"
	"encoding/json"
	"io"
	"net/http"
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_CreateChatCompletions_OK(t *testing.T) {
	ctx := context.Background()
	method := http.MethodPost

	b := CreateChatCompletionsRequest{
		Model: "gpt-4",
		Messages: []Message{
			{Role: "user", Content: "Hi!"},
		},
		ResponseFormat: ResponseFormat{
			Type: "text",
		},
	}
	reqJson, err := json.Marshal(b)
	require.NoError(t, err)
	body := bytes.NewReader(reqJson)

	mocked := getMockedClient(getMockedClientParams{
		T:       t,
		Context: ctx,
		URL:     completionsUrl,
		Method:  method,
		Body:    body,
	})
	defer mocked.Controller.Finish()

	expected := CreateChatCompletionsResponse{}
	responseJson, err := json.Marshal(expected)
	require.NoError(t, err)
	responseBody := io.NopCloser(bytes.NewReader(responseJson))

	expectedHttpResponse := &http.Response{Status: "200 OK", Body: responseBody}
	mocked.HttpClient.EXPECT().Do(newRequestMatcher(mocked.Request)).Return(expectedHttpResponse, nil)

	response, err := mocked.Client.CreateChatCompletions(ctx, b)

	// asserts
	require.NoError(t, err)
	require.Equal(t, expected, response)
}
