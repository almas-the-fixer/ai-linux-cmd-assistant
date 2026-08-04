package main

import (
	"ai-linux-cmd-assistant/internal/ollama"
	"fmt"
)

func main() {
	// Make Client
	c := ollama.NewClient()

	// Call Generate
	resp, err := c.Generate("What is Linux in Brief?")
	if err != nil {
		fmt.Println("Error Generating: %w", err)
	}	

	// Print Output
	fmt.Println(resp)
}