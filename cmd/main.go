package main

import (
	"ai-linux-cmd-assistant/internal/cli"
	"ai-linux-cmd-assistant/internal/knowledgebase"
	"ai-linux-cmd-assistant/internal/ollama"
	"fmt"
)

func main() {
	err := knowledgebase.LoadKnowledge()
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	
	// Make Client
	c := ollama.NewClient()
	cli := cli.NewCLI(c)

	cli.Run()
}
