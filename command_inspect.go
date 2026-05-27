package main

import "fmt"

func commandInspect(config *commandConfig) error {
	if config.name == nil {
		fmt.Println("missing Pokemon to inspect...")
		return nil
	}
	pokemon, ok := config.caughtPokemon[*config.name]
	if !ok {
		fmt.Println("you have not caught that pokemon")
		return nil
	}
	fmt.Println("Name: ", pokemon.Name)
	fmt.Println("Height: ", pokemon.Height)
	fmt.Println("Weight: ", pokemon.Weight)
	fmt.Println("Stats: ")
	for _, stat := range pokemon.Stats {
		fmt.Printf("  -%s: %d\n", stat.Stat.Name, stat.BaseStat)
	}
	fmt.Println("Types: ")
	for _, pokemonType := range pokemon.Types {
		fmt.Println("  -", pokemonType.Type.Name)
	}
	return nil
}
