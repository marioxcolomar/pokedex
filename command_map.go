package main

import (
	"errors"
	"fmt"
)

func commandMap(config *commandConfig) error {
	res, err := config.pokeapiClient.ListLocationAreas(config.nextUrl)
	if err != nil {
		return err
	}

	for _, location := range res.Results {
		fmt.Println(location.Name)
	}

	config.nextUrl = res.Next
	config.previousUrl = res.Previous

	return nil
}

func commandMapBack(config *commandConfig) error {
	if config.previousUrl == nil {
		return errors.New("you are on the first page")

	}
	res, err := config.pokeapiClient.ListLocationAreas(config.previousUrl)
	if err != nil {
		return err
	}

	for _, location := range res.Results {
		fmt.Println(location.Name)
	}

	config.nextUrl = res.Next
	config.previousUrl = res.Previous

	return nil
}
