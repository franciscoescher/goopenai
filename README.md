# Go OpenAI

This is a Go client library for the OpenAI API.

## Installation

    go get github.com/franciscoescher/goopenai

## Usage

Fiest, you need to create a client with the api key and organization id.

```
client := goopenai.NewClient(apiKey, organization)
```

Then, you can use the client to call the api.

Example:

```
package main

import (
	"context"
  "fmt"

  "github.com/franciscoescher/golinkedin"
)

func main() {
	apiKey := os.Getenv("API_KEY")
	organization := os.Getenv("API_ORG")

	client := goopenai.NewClient(apiKey, organization)

	r := goopenai.CompletionsRequest{
		Model: "gpt-3.5-turbo",
		Messages: []goopenai.Message{
			{
				Role:    "user",
				Content: "Say this is a test!",
			},
		},
		Temperature: 0.7,
	}

	completions, err := client.Completions(r)
	if err != nil {
		panic(err)
	}

	fmt.Println(completions)
}
```

Run this code using:
`API_KEY=<your-api-key> API_ORG=<your-org-id> go run .`

## Note

This library is not complete and not fully tested.

Feel free to contribute.
