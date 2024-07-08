package main

import (
	"fmt"
	"log"
	"os"
	"path/filepath"

	openai "github.com/sashabaranov/go-openai"
)

var orgID = "org-NtggwnykKVSq9Djj1GoV5TWO"

func main() {
	if err := mainE(); err != nil {
		log.Fatalf("Unexpected error: %v", err)
	}
}

func mainE() error {
	client, err := createOpenAIClient()
	if err != nil {
		return fmt.Errorf("failed to create OpenAI client: %w", err)
	}
	fmt.Println(client)
	return nil
}

func createOpenAIClient() (*openai.Client, error) {
	token, err := loadAuthToken()
	if err != nil {
		return nil, fmt.Errorf("failed to load auth token: %w", err)
	}
	config := openai.DefaultConfig(token)
	config.OrgID = orgID
	client := openai.NewClientWithConfig(config)
	return client, nil
}

func loadAuthToken() (string, error) {
	home, err := os.UserHomeDir()
	if err != nil {
		return "", fmt.Errorf("could not find home directory: %w", err)
	}
	keyFile := filepath.Join(home, ".local", "keys", "cadgpt")
	key, err := os.ReadFile(keyFile)
	if err != nil {
		return "", fmt.Errorf("key file at %q could not be read: %w", keyFile, err)
	}
	return string(key), nil
}
