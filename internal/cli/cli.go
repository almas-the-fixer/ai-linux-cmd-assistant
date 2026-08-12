package cli

import (
	"ai-linux-cmd-assistant/internal/command"
	"ai-linux-cmd-assistant/internal/intent"
	"ai-linux-cmd-assistant/internal/ollama"
	"ai-linux-cmd-assistant/internal/prompt"
	"ai-linux-cmd-assistant/internal/security"
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

		if userIntent == intent.GenerateCommand {
			// Need to get user's intent before parsing and parse only when it is GENERATE COMMAND
			cmd, err := command.ParseResponse(resp)
			if err != nil {
				ui.PrintError(err)
				fmt.Print("> ")
				continue
			} else {
				err := security.ValidateCommand(cmd.Command)
				if err != nil {
					ui.PrintError(err)
					fmt.Print("> ")
					continue
				}
				ui.PrintCommand(cmd.Command, cmd.Explanation)
				fmt.Print("> ")
			}
		} else {
			// Print General Responses .. i.e TROUBLESHOOT, HOW_TO etc
			ui.PrintResponse(resp)
			fmt.Print("> ")
		}

	}
	err := scanner.Err()
	if err != nil {
		log.Fatal("An Error Occured While Scanning Input: ", err)
	}
}
