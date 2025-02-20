package common

import (
	"log"
	"os"
	"strings"
)

func Input() string {
	bytes, err := os.ReadFile("input.txt")
	if err != nil {
		log.Fatal("Could not find input.txt!")
	}
	return string(bytes)
}

func InputLines() []string {
	text := Input()
	text = strings.TrimRight(text, "\n") // remove possible trailing newlines
	return strings.Split(text, "\n")
}
