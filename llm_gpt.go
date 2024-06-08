package main

import (
	"context"
	openai "github.com/sashabaranov/go-openai"
	"os"
)

func ChatQuery(query string) string {
	apiKey := os.Getenv("API_KEY")

	client := openai.NewClient(apiKey)

	response, err := client.CreateChatCompletion(context.Background(), openai.ChatCompletionRequest{
		Model: openai.GPT4o,
		Messages: []openai.ChatCompletionMessage{
			{
				Role:    openai.ChatMessageRoleUser,
				Content: query,
			},
		},
	})
	if err != nil {
		panic(err)
	}

	return response.Choices[0].Message.Content
}
