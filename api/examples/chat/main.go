package main

import (
	"context"
	"fmt"
	"log"

	"github.com/nikolaydimitrov/ollama/api"
)

func main() {
	client, err := api.ClientFromEnvironment()
	if err != nil {
		log.Fatal(err)
	}

	messages := []api.Message{
		api.Message{
			Role:    "system",
			Content: "Provide very brief, concise responses",
		},
		api.Message{
			Role:    "user",
			Content: "Name some unusual animals",
		},
		api.Message{
			Role:    "assistant",
			Content: "Monotreme, platypus, echidna",
		},
		api.Message{
			Role:    "user",
			Content: "which of these is the most dangerous?",
		},
	}

	ctx := context.Background()
	req := &api.ChatRequest{
		Model:    "llama3.1",
		Messages: messages,
		Tools: api.Tools{
			api.Tool{
				Type: "function",
				Function: api.ToolFunction{
					Name:        "get_weather",
					Description: "Get the weather for a location",
					Parameters: api.ToolParameters{
						Type: "object",
						Properties: map[string]api.ToolProperty{
							"location": api.ToolProperty{
								Type:        "string",
								Description: "The location to get the weather for",
							},
						},
					},
				},
			},
		},
	}

	respFunc := func(resp api.ChatResponse) error {
		fmt.Print(resp.Message.Content)
		return nil
	}

	err = client.Chat(ctx, req, respFunc)
	if err != nil {
		log.Fatal(err)
	}
}
