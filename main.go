package main

import (
	"context"
	"fmt"
	"log"
	"os"
	"time"

	"github.com/predictionguard/go-client"
)

func main() {
	if err := run(); err != nil {
		log.Fatalln(err)
	}
}

func run() error {
	host := "https://api.predictionguard.com"
	apiKey := os.Getenv("PGKEY")

	logger := func(ctx context.Context, msg string, v ...interface{}) {
		s := fmt.Sprintf("msg: %s", msg)
		for i := 0; i < len(v); i = i + 2 {
			s = s + fmt.Sprintf(", %s: %v", v[i], v[i+1])
		}
		log.Println(s)
	}

	cln := client.New(logger, host, apiKey)

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	// Define max tokens and temperature as pointers
	maxTokens := 1000
	temperature := float32(0.3)

	// Use client.Roles for Role values
	input := client.ChatInput{
		Model: "Hermes-2-Pro-Llama-3-8B",
		Messages: []client.ChatInputMessage{
			{
				Role:    client.Roles.System, // Access the system role from client.Roles
				Content: "You are a helpful coding assistant.",
			},
			{
				Role:    client.Roles.User, // Access the user role from client.Roles
				Content: "Write a Go program that prints out random numbers.",
			},
		},
		MaxTokens:   &maxTokens,   // Use pointer for MaxTokens
		Temperature: &temperature, // Use pointer for Temperature
	}

	resp, err := cln.Chat(ctx, input)
	if err != nil {
		return fmt.Errorf("ERROR: %w", err)
	}

	fmt.Println(resp.Choices[0].Message.Content)

	return nil
}
