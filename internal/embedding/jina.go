package embedding

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
)

const jinaEmbeddingURL = "https://api.jina.ai/v1/embeddings"
const jinaModel = "jina-embeddings-v5-text-nano"

type embeddingRequest struct {
	Model      string   `json:"model"`
	Task       string   `json:"task"`
	Dimensions int      `json:"dimensions"`
	Input      []string `json:"input"`
}

type embeddingResponse struct {
	Data []struct {
		Embedding []float32 `json:"embedding"`
	} `json:"data"`
}

func EmbedText(text string, task string) ([]float32, error) {
	apiKey := os.Getenv("JINA_API_KEY")
	if apiKey == "" {
		return nil, fmt.Errorf("JINA_API_KEY is not set")
	}

	requestBody := embeddingRequest{
		Model:      jinaModel,
		Task:       task,
		Dimensions: 768,
		Input:      []string{text},
	}

	body, err := json.Marshal(requestBody)
	if err != nil {
		return nil, fmt.Errorf("failed to encode embedding request: %w", err)
	}

	req, err := http.NewRequest(
		http.MethodPost,
		jinaEmbeddingURL,
		bytes.NewReader(body),
	)
	if err != nil {
		return nil, fmt.Errorf("failed to create embedding request: %w", err)
	}

	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", "Bearer "+apiKey)

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, fmt.Errorf("failed to send embedding request: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode < 200 || resp.StatusCode >= 300 {
		responseBody, _ := io.ReadAll(resp.Body)

		return nil, fmt.Errorf(
			"Jina embedding request failed with status %d: %s",
			resp.StatusCode,
			string(responseBody),
		)
	}

	var result embeddingResponse

	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		return nil, fmt.Errorf("failed to decode embedding response: %w", err)
	}

	if len(result.Data) == 0 || len(result.Data[0].Embedding) == 0 {
		return nil, fmt.Errorf("Jina returned an empty embedding")
	}

	return result.Data[0].Embedding, nil
}