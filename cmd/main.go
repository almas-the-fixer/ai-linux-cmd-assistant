package main

import (
	"ai-linux-cmd-assistant/internal/cli"
	"ai-linux-cmd-assistant/internal/mongodb"
	"ai-linux-cmd-assistant/internal/ollama"
	"context"
	"fmt"
)

func main() {
	// Try MongoDB Conn
	client, err := mongodb.ConnectAtlasDB()
	if err != nil {
		fmt.Print("an error occured: ", err)
		return
	}
	defer client.Disconnect(context.Background())

	// Make Ollama Client
	c := ollama.NewClient()

	// Make CLI
	cli := cli.NewCLI(c)

	cli.Run()
}
