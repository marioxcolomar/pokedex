package main

import (
	"fmt"
	"os"
)

func commandExit(config *commandConfig) error {
	fmt.Print("Closing the Pokedex... Goodbye!\n")
	os.Exit(0)
	return fmt.Errorf("program closed")
}
