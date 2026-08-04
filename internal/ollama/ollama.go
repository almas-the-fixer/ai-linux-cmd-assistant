package ollama

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

const (
	model = "qwen2.5:3b"
)

func (c *Client) Generate(prompt string) (string, error) {
	request := GenerateRequest{
		Model:  model,
		Prompt: prompt,
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
