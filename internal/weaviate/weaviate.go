package weaviate

import (
	"context"
	"fmt"
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
		Properties: []*models.Property{
			{
				Name:     "content",
				DataType: schema.DataTypeText.PropString(),
			},
			{
				Name:     "sourceDoc",
				DataType: schema.DataTypeText.PropString(),
				ModuleConfig: map[string]interface{}{
					"text2vec-weaviate": map[string]interface{}{
						"skip": true,
					},
				},
			},
		},
		VectorConfig: map[string]models.VectorConfig{
			"default": {
				Vectorizer: map[string]interface{}{
					"text2vec-weaviate": map[string]interface{}{
						"model":      "Snowflake/snowflake-arctic-embed-m-v1.5",
						"dimensions": 768,
						"properties": []string{"content"},
					},
				},
			},
		},
	}

	// Avoid failing if the collection already exists.
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
