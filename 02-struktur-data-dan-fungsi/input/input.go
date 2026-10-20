package input

import (
	"fmt"
	"strings"
)

func InputInt(prompt string) (int, error) {
	fmt.Print(prompt)
	var value int
	if _, err := fmt.Scanln(&value); err != nil {
		return 0, err
	}
	return value, nil
}

func InputString(prompt string) (string, error) {
	fmt.Print(prompt)
	var value string
	if _, err := fmt.Scanln(&value); err != nil {
		return "", err
	}
	return strings.TrimSpace(value), nil
}
