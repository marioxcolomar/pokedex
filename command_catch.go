package main

import (
	"fmt"
	"math"
	"math/rand"
)

func commandCatch(config *commandConfig) error {
	if config.name == nil {
		fmt.Println("missing Pokemon to catch...")
		return nil
	}
	fmt.Printf("Throwing a Pokeball at %s...\n", *config.name)
	pokemon, err := config.pokeapiClient.GetPokemon(*config.name)
	if err != nil {
		return err
	}
	if pokemon.Name == "" {
		return fmt.Errorf("Pokemon %s not found", *config.name)
	}
	top := int(math.Round(float64(pokemon.BaseExperience) * 0.8))
	chance := rand.Intn(pokemon.BaseExperience)
	if chance < top {
		fmt.Printf("%s escaped!\n", pokemon.Name)
		return nil
	}
	fmt.Printf("%s was caught!\n", pokemon.Name)
	config.caughtPokemon[pokemon.Name] = pokemon
	return nil
}
