package main

import (
	"fmt"

	"ai-linux-cmd-assistant/internal/knowledgebase"
	"ai-linux-cmd-assistant/internal/weaviate"

	"github.com/joho/godotenv"
)

func main() {
	// Load environment variables once for the indexer.
	if err := godotenv.Load(); err != nil {
		fmt.Println("Error loading .env:", err)
		return
	}

	// Connect to Weaviate Cloud.
	client, err := weaviate.ConnectWeaviate()
	if err != nil {
		fmt.Println("Error connecting to Weaviate:", err)
		return
	}

	// Create the collection if it does not already exist.
	err = weaviate.CreateKnowledgeCollection(client)
	if err != nil {
		fmt.Println("Error creating knowledge collection:", err)
		return
	}

	indexed, err := weaviate.KnowledgeIndexed(client)
	if err != nil {
		fmt.Println("Error checking index:", err)
		return
	}

	if indexed {
		fmt.Println("Knowledge base already indexed. Nothing to do.")
		return
	}

	// Load and split all Markdown documents into chunks.
	chunks, err := knowledgebase.LoadKnowledge()
	if err != nil {
		fmt.Println("Error loading knowledge:", err)
		return
	}

	fmt.Println("Total chunks:", len(chunks))

	// R3: insert chunks into Weaviate will go here.
	err = weaviate.InsertChunks(client, chunks)
	if err != nil {
		fmt.Println("Error inserting chunks:", err)
		return
	}

	fmt.Printf("Inserted %d chunks into Weaviate\n", len(chunks))
	fmt.Println("Indexer finished.")
}
