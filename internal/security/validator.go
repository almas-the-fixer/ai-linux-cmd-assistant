package security

import (
	"errors"
	"strings"
)

func ValidateCommand(command string) error {
	normalized := strings.ToLower(strings.TrimSpace(command))

	// Dangerous filesystem deletion
	dangerousPatterns := []string{
		"rm -rf /",
		"rm -rf /*",
		"rm -r -f /",
		"rm -f -r /",

		// Fork bomb
		":(){ :|:& };:",

		// Filesystem destruction
		"wipefs",
		"mkfs",
		"shred",

		// Security disabling
		"ufw disable",
		"iptables -f",

		// Remote code execution
		"| sh",
		"| bash",
		"|sh",
		"|bash",
	}

	for _, pattern := range dangerousPatterns {
		if strings.Contains(normalized, pattern) {
			return errors.New("potentially destructive or unsafe command blocked")
		}
	}

	return nil
}