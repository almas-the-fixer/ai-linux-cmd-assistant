package ollama

import (
	"ai-linux-cmd-assistant/internal/intent"
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
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

	requestBody, err := json.Marshal(request)
	if err != nil {
		return "", fmt.Errorf("marshal request: %w", err)
	}

	reader := bytes.NewReader(requestBody)

	ollamaReq, err := http.NewRequest("POST", c.baseURL+"/api/generate", reader)
	if err != nil {
		return "", fmt.Errorf("error making request: %w", err)
	}

	ollamaReq.Header.Add("Content-Type", "application/json")

	httpClient := &http.Client{}

	resp, err := httpClient.Do(ollamaReq)
	if err != nil {
		return "", fmt.Errorf("error sending request: %w", err)
	}
	if resp.StatusCode != 200 {
		return "", fmt.Errorf("http error: %w", err)
	}
	defer resp.Body.Close()

	bodyBytes, err := io.ReadAll(resp.Body)
	if err != nil {
		return "", fmt.Errorf("error reading response: %w", err)
	}

	modelResponse := new(GenerateResponse)

	err = json.Unmarshal(bodyBytes, modelResponse)
	if err != nil {
		return "", fmt.Errorf("error unmarshaling response: %w", err)
	}

	return modelResponse.Response, nil
}

func (c *Client) IntentDetector(userPrompt string) (intent.Intent, error) {
	request := GenerateRequest{
		Model:  model,
		Prompt: intent.BuildIntentPrompt(userPrompt),
		Stream: false,
	}

	requestBody, err := json.Marshal(request)
	if err != nil {
		return intent.FailedToGetIntent, fmt.Errorf("marshal request: %w", err)
	}

	reader := bytes.NewReader(requestBody)

	ollamaReq, err := http.NewRequest("POST", c.baseURL+"/api/generate", reader)
	if err != nil {
		return intent.FailedToGetIntent, fmt.Errorf("error making request: %w", err)
	}

	ollamaReq.Header.Add("Content-Type", "application/json")

	httpClient := &http.Client{}

	resp, err := httpClient.Do(ollamaReq)
	if err != nil {
		return intent.FailedToGetIntent, fmt.Errorf("error sending request: %w", err)
	}
	if resp.StatusCode != 200 {
		return intent.FailedToGetIntent, fmt.Errorf("http error: %w", err)
	}
	defer resp.Body.Close()

	bodyBytes, err := io.ReadAll(resp.Body)
	if err != nil {
		return intent.FailedToGetIntent, fmt.Errorf("error reading response: %w", err)
	}

	modelResponse := new(GenerateResponse)

	err = json.Unmarshal(bodyBytes, modelResponse)
	if err != nil {
		return intent.FailedToGetIntent, fmt.Errorf("error unmarshaling response: %w", err)
	}

	label := strings.TrimSpace(modelResponse.Response)
	
	// For debugging
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
