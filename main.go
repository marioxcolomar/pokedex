package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	config := commandConfig{}
	fmt.Print("Pokedex > ")
	for scanner.Scan() {
		input := strings.ToLower(strings.Trim(scanner.Text(), " "))
		command := strings.Split(input, " ")[0]
		runCommand(command, &config)
	}

	if err := scanner.Err(); err != nil {
		fmt.Printf("Invalid input: %s", err)
	}
}
