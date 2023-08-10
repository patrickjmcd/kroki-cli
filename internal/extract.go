package internal

import (
	"fmt"
	"strings"
)

func ExtractContent(input, inputType string) (string, error) {
	startDelimiter := fmt.Sprintf("```%s", inputType)
	endDelimiter := "```"
	start := strings.Index(input, startDelimiter)
	if start == -1 {
		return "", fmt.Errorf("could not find the start delimiter %s", startDelimiter)
	}
	start += len(startDelimiter) // Move the pointer right after the delimiter
	end := strings.Index(input[start:], endDelimiter)
	if end == -1 {
		return "", fmt.Errorf("could not find the end delimiter %s", endDelimiter)
	}
	return input[start : start+end], nil
}
