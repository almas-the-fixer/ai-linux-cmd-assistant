package executor

import (
	"errors"
	"os/exec"
	"strings"
)

func ExecuteCommand(command string) (string, error) {
	// Get Command field from Command Struct AFTER Validation
	if command == "" {
		return "", errors.New("Command Field is Empty Cannot Execute!")
	}
	// Split the Line into <command> <flags> <args> i.e ls -la where ls is cmd and -la is flag / argument
	fields := strings.Fields(command)

	if len(fields) == 0 {
		return "", errors.New("Error Empty Command")
	}
	// Tried to seperate the binary?
	binary := fields[0]

	// tried to seperate the args
	args := fields[1:]

	// Pass Them into exec.Command Execute command if user types y
	cmd := exec.Command(binary, args...)

	// Capture Output and return output with nil error and vice versa if error the empty string and error
	output, err := cmd.CombinedOutput()
	if err != nil {
		return string(output), err
	}
	
	return string(output), nil
}