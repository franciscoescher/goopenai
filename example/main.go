package main

import (
	"fmt"
	"os"

	"github.com/franciscoescher/goopenai"
)

func main() {
	apiKey := os.Getenv("API_KEY")
	organization := os.Getenv("API_ORG")

	client := goopenai.NewClient(apiKey, organization)

	models, err := client.GetModelsRaw()
	if err != nil {
		panic(err)
	}
	fmt.Println(string(models))

	r := goopenai.CreateCompletionsRequest{
		Model: "gpt-3.5-turbo",
		Messages: []goopenai.Message{
			{
				Role:    "user",
				Content: "Say this is a test!",
			},
		},
		Temperature: 0.7,
	}

	completions, err := client.CreateCompletions(r)
	if err != nil {
		panic(err)
	}

	fmt.Println(completions)
	/*
		{
		  "id": "chatcmpl-xxx",
		  "object": "chat.completion",
		  "created": 1678667132,
		  "model": "gpt-3.5-turbo-0301",
		  "usage": {
		    "prompt_tokens": 13,
		    "completion_tokens": 7,
		    "total_tokens": 20
		  },
		  "choices": [
		    {
		      "message": {
		        "role": "assistant",
		        "content": "\n\nThis is a test!"
		      },
		      "finish_reason": "stop",
		      "index": 0
		    }
		  ]
		}
	*/
}
