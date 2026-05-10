package main

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"strings"
)

func cleanInput(text string) []string {
	lower := strings.ToLower(text)
	words := strings.Fields(lower)
	return words
}

func commandExit(config *commandConfig) error {
	fmt.Print("Closing the Pokedex... Goodbye!\n")
	os.Exit(0)
	return fmt.Errorf("program closed")
}

func commandHelp(config *commandConfig) error {
	fmt.Print(`Welcome to the Pokedex!
Usage:

help: Displays a help message
exit: Exit the Pokedex`)
	fmt.Print("\n")
	return nil
}

type commandConfig struct {
	id          string
	nextUrl     string
	previousUrl string
}

type JsonResponse struct {
	Count    int    `json:"count"`
	Next     string `json:"next"`
	Previous string `json:"previous"`
	Results  []struct {
		Name string `json:"name"`
		URL  string `json:"url"`
	} `json:"results"`
}

type PokeResponse struct {
	locations []string
	response  JsonResponse
}

func getPokeApi(url string) PokeResponse {
	response, err := http.Get(url)
	if err != nil {
		log.Fatal(err)
	}
	body, err := io.ReadAll(response.Body)
	response.Body.Close()
	if response.StatusCode > 299 {
		log.Fatalf("Response failed with status code: %d and \nbody: %s\n", response.StatusCode, body)
	}
	if err != nil {
		log.Fatal(err)
	}
	responseJson := JsonResponse{}
	errJson := json.Unmarshal(body, &responseJson)
	if errJson != nil {
		fmt.Println(errJson)
	}
	locations := make([]string, 0, len(responseJson.Results))
	for _, loc := range responseJson.Results {
		locations = append(locations, loc.Name)
	}
	return PokeResponse{locations, responseJson}
}

func commandMap(config *commandConfig) error {
	baseUrl := "https://pokeapi.co/api/v2/location-area/"
	if config.nextUrl != "" {
		baseUrl = config.nextUrl
	}
	res := getPokeApi(baseUrl)

	locationAreas := strings.Join(res.locations, "\n")
	fmt.Printf("%s", locationAreas)
	fmt.Print("\n")

	config.nextUrl = res.response.Next
	config.previousUrl = res.response.Previous

	return nil
}

func commandMapBack(config *commandConfig) error {
	if config.previousUrl == "" {
		fmt.Println("you are on the first page")
		return nil
	}
	baseUrl := config.previousUrl
	res := getPokeApi(baseUrl)

	locationAreas := strings.Join(res.locations, "\n")
	fmt.Printf("%s", locationAreas)
	fmt.Print("\n")

	config.nextUrl = res.response.Next
	config.previousUrl = res.response.Previous

	return nil
}

type cliCommand struct {
	name        string
	description string
	callback    func(config *commandConfig) error
}

func runCommand(command string, config *commandConfig) {
	commandMap := map[string]cliCommand{
		"help": {
			name:        "help",
			description: "Show list of commands",
			callback:    commandHelp,
		},
		"exit": {
			name:        "exit",
			description: "Exit the Pokedex",
			callback:    commandExit,
		},
		"map": {
			name:        "map",
			description: "List of maps in Pokedex",
			callback:    commandMap,
		},
		"mapb": {
			name:        "map",
			description: "List of previous maps in Pokedex",
			callback:    commandMapBack,
		},
	}
	cmd, ok := commandMap[command]
	if !ok {
		fmt.Print("Unknown command\n")
	} else {
		cmd.callback(config)
	}
}
