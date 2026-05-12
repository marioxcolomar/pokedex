package main

import (
	"fmt"
)

func commandHelp(config *commandConfig) error {
	fmt.Print(`Welcome to the Pokedex!
Usage:

help: Displays a help message
exit: Exit the Pokedex`)
	fmt.Print("\n")
	return nil
}
