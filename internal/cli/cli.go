package cli

import (
	"ai-linux-cmd-assistant/internal/command"
	"ai-linux-cmd-assistant/internal/executor"
	"ai-linux-cmd-assistant/internal/intent"
	"ai-linux-cmd-assistant/internal/knowledgebase"
	"ai-linux-cmd-assistant/internal/ollama"
	"ai-linux-cmd-assistant/internal/prompt"
	"ai-linux-cmd-assistant/internal/security"
	"ai-linux-cmd-assistant/internal/ui"
	"ai-linux-cmd-assistant/internal/weaviate"
	"bufio"
	"fmt"
	wclient "github.com/weaviate/weaviate-go-client/v5/weaviate"
	"log"
	"os"
)

type CLI struct {
	Client         *ollama.Client
	WeaviateClient *wclient.Client
}

func NewCLI(
	client *ollama.Client,
	weaviateClient *wclient.Client,
) *CLI {
	return &CLI{
		Client:         client,
		WeaviateClient: weaviateClient,
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

		var chunks []knowledgebase.Chunk

		if userIntent != intent.OffTopic &&
			userIntent != intent.Malicious &&
			userIntent != intent.FailedToGetIntent {

			chunks, err = weaviate.SearchChunks(
				cli.WeaviateClient,
				input,
			)
			if err != nil {
				ui.PrintError(err)
				fmt.Print("> ")
				continue
			}
		}
		fmt.Println("Retrieved chunks:", len(chunks))

		for i, chunk := range chunks {
			fmt.Printf("Chunk %d | %s\n", i+1, chunk.SourceDoc)
		}
		// Building Prompt
		prompt := prompt.BuildPrompt(input, userIntent, chunks)
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

				// Confirm With user
				choice := ui.ConfirmExecution(scanner)
				if !choice {
					ui.PrintPermissionDenied()
					fmt.Print("> ")
					continue
				}
				output, err := executor.ExecuteCommand(cmd.Command)
				if err != nil {
					ui.PrintError(err)
					fmt.Print("> ")
					continue
				}
				ui.PrintResponse(output)
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
