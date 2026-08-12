package cli

import (
	"ai-linux-cmd-assistant/internal/ollama"
	"ai-linux-cmd-assistant/internal/prompt"
	"ai-linux-cmd-assistant/internal/ui"
	"bufio"
	"fmt"
	"log"
	"os"
)

type CLI struct {
	Client *ollama.Client
}

func NewCLI(client *ollama.Client) *CLI {
	return &CLI{
		Client: client,
	}
}

func (cli *CLI) Run() {
	// Print Welcome
	fmt.Println("AI Linux Assistant")
	fmt.Println("Type 'exit' to quit.")

	// Loop Forever
	fmt.Print("> ")
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		input := scanner.Text()
		if input == "exit" {
			fmt.Println("Goodbye!")
			break
		}

		// Getting User's Intent
		userIntent, err := cli.Client.IntentDetector(input)
		if err != nil {
			ui.PrintError(err)
			fmt.Print("> ")
			continue
		}

		// Building Prompt
		prompt := prompt.BuildPrompt(input, userIntent)

		resp, err := cli.Client.Generate(prompt)
		if err != nil {
			ui.PrintError(err)
			fmt.Print("> ")
			continue
		}
		ui.PrintResponse(resp)
		fmt.Print("> ")

	}
	err := scanner.Err()
	if err != nil {
		log.Fatal("An Error Occured While Scanning Input: ", err)
	}
}
