package ollama

import (
	"ai-linux-cmd-assistant/internal/intent"
	"fmt"
	"strings"
)

const (
	model = "qwen2.5:3b"
)

// Type
type Client struct {
	baseURL string
}

// Constructor that makes a Valid Client
func NewClient() *Client {
	return &Client{
		baseURL: "http://localhost:11434",
	}
}

// Method of a Client, i.e A Client can Generate response.
func (c *Client) Generate(userPrompt string) (string, error) {
	request := GenerateRequest{
		Model:  model,
		Prompt: userPrompt,
		Stream: false,
	}

	response, err := c.doRequest(&request)
	if err != nil {
		return "", fmt.Errorf("error making request to Ollama: %w", err)
	}

	return response.Response, nil
}

func (c *Client) IntentDetector(userPrompt string) (intent.Intent, error) {
	request := GenerateRequest{
		Model:  model,
		Prompt: intent.BuildIntentPrompt(userPrompt),
		Stream: false,
	}

	response, err := c.doRequest(&request)
	if err != nil {
		return intent.FailedToGetIntent, fmt.Errorf("error making request to Ollama: %w", err)
	}
	// For debugging
	label := strings.TrimSpace(response.Response)

	fmt.Println("User's Intent: ", label)
	switch label {
	case "HOW_TO":
		return intent.HowTo, nil
	case "EXPLAIN_COMMAND":
		return intent.ExplainCommand, nil
	case "TROUBLESHOOT":
		return intent.Troubleshoot, nil
	case "OFF_TOPIC":
		return intent.OffTopic, nil
	case "MALICIOUS":
		return intent.Malicious, nil
	default:
		return intent.FailedToGetIntent, fmt.Errorf("unexpected intent label: %q", label)
	}
}
