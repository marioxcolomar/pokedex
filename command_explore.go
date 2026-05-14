package main

import "fmt"

func commandExplore(config *commandConfig) error {
	res, err := config.pokeapiClient.ListPokemonPerLocationArea(*config.name)
	if err != nil {
		return err
	}
	for _, encounter := range res.PokemonEncounters {
		fmt.Println(encounter.Pokemon.Name)
	}
	return nil
}
