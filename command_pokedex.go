package main

import "fmt"

func commandPokedex(config *commandConfig) error {
	if len(config.caughtPokemon) == 0 {
		fmt.Println("you have not caught any Pokemon")
		return nil
	}
	fmt.Println("Your Pokedex:")
	for _, pokemon := range config.caughtPokemon {
		fmt.Println("  -", pokemon.Name)
	}
	return nil
}
