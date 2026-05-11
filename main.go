package main

import (
	"bufio"
	"fmt"
	"os"
	"pokedex/internal/pokecache"
	"strings"
	"time"
)

func main() {
	pokeClient := pokecache.NewCache(5 * time.Second)
	scanner := bufio.NewScanner(os.Stdin)
	config := commandConfig{}
	fmt.Print("Pokedex > ")
	for scanner.Scan() {
		input := strings.ToLower(strings.Trim(scanner.Text(), " "))
		command := strings.Split(input, " ")[0]
		runCommand(command, &config, pokeClient)
	}

	if err := scanner.Err(); err != nil {
		fmt.Printf("Invalid input: %s", err)
	}
}
