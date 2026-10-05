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

	// Load and split all Markdown documents into chunks.
	chunks, err := knowledgebase.LoadKnowledge()
	if err != nil {
		fmt.Println("Error loading knowledge:", err)
		return
	}

	fmt.Println("Total chunks:", len(chunks))

	// R3: insert chunks into Weaviate will go here.

	fmt.Println("Indexer finished.")
}