package main

import (
	"fmt"
	"pokedex/internal/pokeapi"
	"strings"
)

func cleanInput(text string) []string {
	lower := strings.ToLower(text)
	words := strings.Fields(lower)
	return words
}

type commandConfig struct {
	pokeapiClient pokeapi.Client
	id            *string
	nextUrl       *string
	previousUrl   *string
}

type cliCommand struct {
	name        string
	description string
	callback    func(*commandConfig) error
}

func runCommand(command string, config *commandConfig) {
	commandMap := map[string]cliCommand{
		"help": {
			name:        "help",
			description: "Show list of commands",
			callback:    commandHelp,
		},
		"exit": {
			name:        "exit",
			description: "Exit the Pokedex",
			callback:    commandExit,
		},
		"map": {
			name:        "map",
			description: "List of maps in Pokedex",
			callback:    commandMap,
		},
		"mapb": {
			name:        "map back",
			description: "List of previous maps in Pokedex",
			callback:    commandMapBack,
		},
		"explore": {
			name:        "explore",
			description: "Explore a given location area in Pokedex",
			callback:    commandExplore,
		},
	}
	cmd, ok := commandMap[command]
	if !ok {
		fmt.Print("Unknown command\n")
	} else {
		cmd.callback(config)
	}
}
