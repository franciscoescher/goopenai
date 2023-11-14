package main

import (
	"context"
	"fmt"
	"os"

	"github.com/franciscoescher/goopenai"
)

func main() {
	client := getClient()

	writeModels(client)

	printCompletions(client)
}

func getClient() *goopenai.Client {
	apiKey := os.Getenv("API_KEY")
	organization := os.Getenv("API_ORG")
	return goopenai.NewClient(apiKey, organization)
}

func writeModels(client *goopenai.Client) {
	models, err := client.ListModelsRaw(context.Background())
	if err != nil {
		panic(err)
	}
	f, err := os.Create("example/models.json")
	if err != nil {
		panic(err)
	}
	defer f.Close()

	// write models to file
	_, err = f.Write(models)
	if err != nil {
		panic(err)
	}
}

func printCompletions(client *goopenai.Client) {
	r := goopenai.CreateChatCompletionsRequest{
		Model: "gpt-4-1106-preview",
		Messages: []goopenai.Message{
			{
				Role:    "user",
				Content: "Respond with a json with your nickname and model data",
			},
		},
		Temperature: 0.7,
		ResponseFormat: goopenai.ResponseFormat{
			Type: "json_object",
		},
	}

	completions, err := client.CreateChatCompletionsRaw(context.Background(), &r)

	if err != nil {
		panic(err)
	}

	fmt.Println(string(completions))
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
