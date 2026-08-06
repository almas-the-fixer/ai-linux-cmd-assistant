package main

import (
	"ai-linux-cmd-assistant/internal/cli"
	"ai-linux-cmd-assistant/internal/ollama"
)

func main() {
	// Make Client
	c := ollama.NewClient()
	cli := cli.NewCLI(c)

	cli.Run()
}