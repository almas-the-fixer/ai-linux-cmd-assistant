package main

import (
	"ai-linux-cmd-assistant/internal/cli"
	"ai-linux-cmd-assistant/internal/ollama"
	"ai-linux-cmd-assistant/internal/weaviate"
	"fmt"
	"github.com/joho/godotenv"
)

func main() {
	//err := knowledgebase.LoadKnowledge()
	//if err != nil {
	//	fmt.Println("Error:", err)
	//	return
	//}
	if err := godotenv.Load(); err != nil {
		fmt.Println("Error loading .env:", err)
		return
	}
	weaviateClient, err := weaviate.ConnectWeaviate()
	if err != nil {
		fmt.Println("Error connecting to Weaviate:", err)
		return
	}
	// Make Client

	c := ollama.NewClient()
	cli := cli.NewCLI(c, weaviateClient)
	cli.Run()
}
