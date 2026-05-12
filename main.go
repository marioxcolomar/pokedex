package main

import (
	"bufio"
	"fmt"
	"os"
	"pokedex/internal/pokeapi"
	"strings"
	"time"
)

func main() {
	pokeClient := pokeapi.NewClient(5*time.Second, time.Minute*5)
	scanner := bufio.NewScanner(os.Stdin)
	config := &commandConfig{
		pokeapiClient: pokeClient,
	}
	fmt.Print("Pokedex > ")
	for scanner.Scan() {
		input := strings.ToLower(strings.Trim(scanner.Text(), " "))
		inputs := strings.Split(input, " ")
		command := inputs[0]
		if len(inputs) > 1 {
			config.id = &inputs[1]
		}
		runCommand(command, config)
	}

	if err := scanner.Err(); err != nil {
		fmt.Printf("Invalid input: %s", err)
	}
}
