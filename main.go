package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	fmt.Print("Pokedex > ")
	for scanner.Scan() {
		input := strings.ToLower(strings.Trim(scanner.Text(), " "))
		firstWord := strings.Split(input, " ")
		fmt.Printf("Your command was: %s\n", firstWord[0])
	}

	if err := scanner.Err(); err != nil {
		fmt.Printf("Invalid input: %s", err)
	}
}
