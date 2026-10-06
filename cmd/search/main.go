// TEST SEARCH PROGRAM
package main

import (
	"fmt"

	"ai-linux-cmd-assistant/internal/weaviate"

	"github.com/joho/godotenv"
)

func main() {
	if err := godotenv.Load(); err != nil {
		fmt.Println("Error loading .env:", err)
		return
	}

	client, err := weaviate.ConnectWeaviate()
	if err != nil {
		fmt.Println("Error connecting to Weaviate:", err)
		return
	}

	results, err := weaviate.SearchChunks(
		client,
		"how do I change file permissions?",
	)
	if err != nil {
		fmt.Println("Search error:", err)
		return
	}

	for i, chunk := range results {
		fmt.Printf("\nResult %d\n", i+1)
		fmt.Println("Source:", chunk.SourceDoc)
		fmt.Println("Content:", chunk.Content)
	}
}