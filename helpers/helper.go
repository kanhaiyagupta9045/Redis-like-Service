package helpers

import (
	"errors"
	"strings"
)

func Validate(s string) (string, error) {
	arr := strings.Fields(s)
	validCommands := []string{"SET", "GET", "QPUSH", "QPOP", "BQPOP"}
	isValidCommand := false
	var operation string
	for _, cmd := range validCommands {
		if arr[0] == cmd {
			isValidCommand = true
			operation = cmd
			break
		}
	}

	if !isValidCommand {
		return " ", errors.New("invalid command")
	}
	if len(arr) < 2 {
		return " ", errors.New("invalid command")
	}
	return operation, nil
}
