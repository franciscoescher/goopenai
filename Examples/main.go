package main

import (
	"context"
	"fmt"
	"os"

	"github.com/Alumniminium/goopenai"
)

func main() {
	client := getClient()
	printCompletions(client)
}

func getClient() *goopenai.Client {
	apiKey := os.Getenv("API_KEY")
	organization := os.Getenv("API_ORG")
	return goopenai.NewClient(apiKey, organization, "http://openai.her.st/v1")
}

func printCompletions(client *goopenai.Client) {
	r := goopenai.CreateCompletionsRequest{
		Prompt:      "This is a cake recipe:\n\n1.",
		MaxTokens:   100,
		Temperature: float64(0.1),
		TopP:        0.1,
		Seed:        1,
	}

	completions, err := client.CreateCompletionsRaw(context.Background(), &r)

	if err != nil {
		panic(err)
	}

	fmt.Println(string(completions))
}
