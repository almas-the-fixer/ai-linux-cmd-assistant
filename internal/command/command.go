package command

import (
	"bufio"
	"errors"
	"strings"
)

type Command struct {
	Command     string
	Explanation string
}

func ParseResponse(response string) (Command, error) {
	scanner := bufio.NewScanner(strings.NewReader(response))
	var commandText, explanationText string
	for scanner.Scan() {
		line := scanner.Text()
		if strings.HasPrefix(line, "COMMAND: ") {
			commandText = strings.TrimPrefix(line, "COMMAND: ")
		} else if strings.HasPrefix(line, "EXPLANATION: ") {
			explanationText = strings.TrimPrefix(line, "EXPLANATION: ")
		}
	}

	if err := scanner.Err(); err != nil {
		return Command{}, err
	}

	if commandText == "" || explanationText == "" {
		return Command{}, errors.New("Failed To Parse Command")
	}
	
	command := Command{
		Command:     commandText,
		Explanation: explanationText,
	}

	return command, nil
}
