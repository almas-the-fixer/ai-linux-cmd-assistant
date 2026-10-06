package weaviate

import (
	"ai-linux-cmd-assistant/internal/embedding"
	"ai-linux-cmd-assistant/internal/knowledgebase"
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"

	"github.com/weaviate/weaviate-go-client/v5/weaviate"
	"github.com/weaviate/weaviate-go-client/v5/weaviate/auth"
	"github.com/weaviate/weaviate/entities/models"
	"github.com/weaviate/weaviate/entities/schema"
)

func ConnectWeaviate() (*weaviate.Client, error) {
	host := os.Getenv("WEAVIATE_HOST")
	apiKey := os.Getenv("WEAVIATE_API_KEY")

	cfg := weaviate.Config{
		Host:       host,
		Scheme:     "https",
		AuthConfig: auth.ApiKey{Value: apiKey},
	}

	client, err := weaviate.NewClient(cfg)
	if err != nil {
		return nil, err
	}

	_, err = client.Misc().ReadyChecker().Do(context.Background())
	if err != nil {
		return nil, err
	}

	return client, nil
}

func CreateKnowledgeCollection(client *weaviate.Client) error {
	ctx := context.Background()

	class := &models.Class{
		Class:       "LinuxKnowledge",
		Description: "Linux and Bash knowledge chunks",
		Vectorizer:  "none",
		Properties: []*models.Property{
			{
				Name:     "content",
				DataType: schema.DataTypeText.PropString(),
			},
			{
				Name:     "sourceDoc",
				DataType: schema.DataTypeText.PropString(),
			},
		},
	}

	exists, err := client.Schema().ClassExistenceChecker().
		WithClassName("LinuxKnowledge").
		Do(ctx)
	if err != nil {
		return fmt.Errorf("failed checking LinuxKnowledge collection: %w", err)
	}

	if exists {
		return nil
	}

	err = client.Schema().ClassCreator().
		WithClass(class).
		Do(ctx)
	if err != nil {
		return fmt.Errorf("failed creating LinuxKnowledge collection: %w", err)
	}

	return nil
}

func InsertChunks(client *weaviate.Client, chunks []knowledgebase.Chunk) error {
	ctx := context.Background()

	objects := make([]*models.Object, 0, len(chunks))

	for _, chunk := range chunks {
		vector, err := embedding.EmbedText(
			chunk.Content,
			"retrieval.passage",
		)
		if err != nil {
			return fmt.Errorf(
				"failed to embed chunk from %s: %w",
				chunk.SourceDoc,
				err,
			)
		}

		objects = append(objects, &models.Object{
			Class: "LinuxKnowledge",
			Properties: map[string]interface{}{
				"content":   chunk.Content,
				"sourceDoc": chunk.SourceDoc,
			},
			Vector: vector,
		})
	}

	_, err := client.Batch().
		ObjectsBatcher().
		WithObjects(objects...).
		Do(ctx)

	if err != nil {
		return fmt.Errorf("failed to insert chunks: %w", err)
	}

	return nil
}

func SearchChunks(client *weaviate.Client, query string) ([]knowledgebase.Chunk, error) {
	host := os.Getenv("WEAVIATE_HOST")
	apiKey := os.Getenv("WEAVIATE_API_KEY")

	if host == "" {
		return nil, fmt.Errorf("WEAVIATE_HOST is not set")
	}

	if apiKey == "" {
		return nil, fmt.Errorf("WEAVIATE_API_KEY is not set")
	}

	// Convert the user's query into the same embedding space
	// used for the stored knowledge chunks.
	queryVector, err := embedding.EmbedText(query, "retrieval.query")
	if err != nil {
		return nil, fmt.Errorf("failed to embed query: %w", err)
	}

	requestBody := struct {
		Vector           []float32 `json:"vector"`
		Limit            int      `json:"limit"`
		ReturnProperties []string `json:"returnProperties"`
		ReturnMetadata   []string `json:"returnMetadata"`
	}{
		Vector:           queryVector,
		Limit:            3,
		ReturnProperties: []string{"content", "sourceDoc"},
		ReturnMetadata:   []string{"distance"},
	}

	body, err := json.Marshal(requestBody)
	if err != nil {
		return nil, fmt.Errorf("failed to marshal search request: %w", err)
	}

	url := fmt.Sprintf(
		"https://%s/v1/search/LinuxKnowledge/near-vector",
		host,
	)

	req, err := http.NewRequest(
		http.MethodPost,
		url,
		bytes.NewReader(body),
	)
	if err != nil {
		return nil, fmt.Errorf("failed to create search request: %w", err)
	}

	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", "Bearer "+apiKey)

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, fmt.Errorf("failed to send search request: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode < 200 || resp.StatusCode >= 300 {
		responseBody, _ := io.ReadAll(resp.Body)

		return nil, fmt.Errorf(
			"Weaviate search failed with status %d: %s",
			resp.StatusCode,
			string(responseBody),
		)
	}

	var searchResponse struct {
		Results []struct {
			Properties struct {
				Content   string `json:"content"`
				SourceDoc string `json:"sourceDoc"`
			} `json:"properties"`
		} `json:"results"`
	}

	if err := json.NewDecoder(resp.Body).Decode(&searchResponse); err != nil {
		return nil, fmt.Errorf("failed to decode Weaviate search response: %w", err)
	}

	chunks := make([]knowledgebase.Chunk, 0, len(searchResponse.Results))

	for _, result := range searchResponse.Results {
		chunks = append(chunks, knowledgebase.Chunk{
			Content:   result.Properties.Content,
			SourceDoc: result.Properties.SourceDoc,
		})
	}

	return chunks, nil
}

func KnowledgeIndexed(client *weaviate.Client) (bool, error) {
	ctx := context.Background()

	objects, err := client.Data().ObjectsGetter().
		WithClassName("LinuxKnowledge").
		WithLimit(1).
		Do(ctx)

	if err != nil {
		return false, fmt.Errorf("failed checking LinuxKnowledge objects: %w", err)
	}

	return len(objects) > 0, nil
}
