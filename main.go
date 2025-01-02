package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"time"

	"github.com/Drnel/btdv_go_pokedex/internal/pokeapi"
	"github.com/Drnel/btdv_go_pokedex/internal/pokecache"
)

type cliCommand struct {
	name        string
	description string
	callback    func(*config, *pokecache.Cache) error
}

type config struct {
	previous string
	next     string
}

func main() {

	input_scanner := bufio.NewScanner(os.Stdin)
	config := config{
		previous: "",
		next:     "initial",
	}
	cache := pokecache.NewCache(time.Second * 50)
	for {
		fmt.Print("Pokedex > ")
		input_scanner.Scan()
		text := input_scanner.Text()
		input_words := cleanInput(text)
		if len(input_words) == 0 {
			continue
		}
		first_word := input_words[0]
		commands := getCommands()
		command, ok := commands[first_word]
		if ok {
			err := command.callback(&config, &cache)
			if err != nil {
				fmt.Println(err)
			}
		} else {
			fmt.Println("Unknown command")
		}
	}
}

func cleanInput(text string) []string {
	text = strings.TrimSpace(text)
	text = strings.ToLower(text)
	words := strings.Fields(text)
	return words
}

func getCommands() map[string]cliCommand {
	return map[string]cliCommand{
		"exit": {
			name:        "exit",
			description: "Exit the Pokedex",
			callback:    commandExit,
		},
		"help": {
			name:        "help",
			description: "Displays a help message",
			callback:    commandHelp,
		},
		"map": {
			name:        "map",
			description: "Displays next 20 locations",
			callback:    commandMap,
		},
		"mapb": {
			name:        "mapb",
			description: "Displays previous 20 locations",
			callback:    commandMapB,
		},
	}
}
func commandExit(config *config, cache *pokecache.Cache) error {
	fmt.Print("Closing the Pokedex... Goodbye!")
	os.Exit(0)
	return nil
}

func commandHelp(config *config, cache *pokecache.Cache) error {
	fmt.Print("Welcome to the Pokedex!\nUsage:\n\n")
	commands := getCommands()
	for _, command := range commands {
		fmt.Printf("%s: %s\n", command.name, command.description)
	}
	return nil
}

func commandMap(config *config, cache *pokecache.Cache) error {
	url := "https://pokeapi.co/api/v2/location-area?offset=0&limit=20"
	if config.next == "" {
		fmt.Println("you're on the last page")
		return nil
	} else if config.next != "initial" {
		url = config.next
	}
	config.previous, config.next = pokeapi.PrintLANames(url, cache)
	return nil
}
func commandMapB(config *config, cache *pokecache.Cache) error {
	url := ""
	if config.previous == "" {
		fmt.Println("you're on the first page")
		return nil
	} else {
		url = config.previous
	}
	config.previous, config.next = pokeapi.PrintLANames(url, cache)
	return nil
}
