package security

import (
	"errors"
	"strings"
)

func ValidateCommand(command string) error {
	if strings.Contains(command, "rm -rf /")	{
		return errors.New("Potentially Destructive Command!")
	}
	return nil
}