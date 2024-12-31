package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	input_scanner := bufio.NewScanner(os.Stdin)
	for {
		fmt.Print("Pokedex > ")
		input_scanner.Scan()
		text := input_scanner.Text()
		input_words := cleanInput(text)
		if len(input_words) == 0 {
			continue
		}
		command := input_words[0]
		fmt.Printf("Your command was: %s\n", command)
	}
}

func cleanInput(text string) []string {
	text = strings.TrimSpace(text)
	text = strings.ToLower(text)
	words := strings.Fields(text)
	return words
}
