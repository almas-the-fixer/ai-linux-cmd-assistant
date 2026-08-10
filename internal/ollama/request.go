package ollama

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

func (c *Client) doRequest(request *GenerateRequest) (*GenerateResponse, error) {
	requestBody, err := json.Marshal(request)

	// Marshal request into JSON
	if err != nil {
		return nil, fmt.Errorf("marshal request: %w", err)
	}

	reader := bytes.NewReader(requestBody)

	// Make Request
	ollamaReq, err := http.NewRequest("POST", c.baseURL+"/api/generate", reader)
	if err != nil {
		return nil, fmt.Errorf("error making request: %w", err)
	}

	// Add Header
	ollamaReq.Header.Add("Content-Type", "application/json")

	// Make Client
	httpClient := &http.Client{}

	// Do Request
	resp, err := httpClient.Do(ollamaReq)
	if err != nil {
		return nil, fmt.Errorf("error sending request: %w", err)
	}
	
	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("ollama returned HTTP status %d", resp.StatusCode)
	}

	defer resp.Body.Close()

	// Read Response
	bodyBytes, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("error reading response: %w", err)
	}

	response := new(GenerateResponse)

	err = json.Unmarshal(bodyBytes, response)
	if err != nil {
		return nil, fmt.Errorf("error unmarshaling response: %w", err)
	}

	return response, nil
}
