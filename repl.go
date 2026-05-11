package main

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"pokedex/internal/pokecache"
	"strings"
)

func cleanInput(text string) []string {
	lower := strings.ToLower(text)
	words := strings.Fields(lower)
	return words
}

func commandExit(config *commandConfig, client *pokecache.Cache) error {
	fmt.Print("Closing the Pokedex... Goodbye!\n")
	os.Exit(0)
	return fmt.Errorf("program closed")
}

func commandHelp(config *commandConfig, client *pokecache.Cache) error {
	fmt.Print(`Welcome to the Pokedex!
Usage:

help: Displays a help message
exit: Exit the Pokedex`)
	fmt.Print("\n")
	return nil
}

type commandConfig struct {
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

func getPokeApi(url string, client *pokecache.Cache) PokeResponse {
	if val, ok := client.Get(url); ok {
		cacheJson := JsonResponse{}
		err := json.Unmarshal(val, &cacheJson)
		if err != nil {
			cacheLocations := make([]string, 0, len(cacheJson.Results))
			for _, loc := range cacheJson.Results {
				cacheLocations = append(cacheLocations, loc.Name)
			}
			return PokeResponse{locations: cacheLocations, response: cacheJson}
		}
	}

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
	client.Add(url, body)
	return PokeResponse{locations, responseJson}
}

func commandMap(config *commandConfig, client *pokecache.Cache) error {
	baseUrl := "https://pokeapi.co/api/v2/location-area/"
	if config.nextUrl != "" {
		baseUrl = config.nextUrl
	}
	res := getPokeApi(baseUrl, client)

	locationAreas := strings.Join(res.locations, "\n")
	fmt.Printf("%s", locationAreas)
	fmt.Print("\n")

	config.nextUrl = res.response.Next
	config.previousUrl = res.response.Previous

	return nil
}

func commandMapBack(config *commandConfig, client *pokecache.Cache) error {
	if config.previousUrl == "" {
		fmt.Println("you are on the first page")
		return nil
	}
	baseUrl := config.previousUrl
	res := getPokeApi(baseUrl, client)

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
	callback    func(config *commandConfig, client *pokecache.Cache) error
}

func runCommand(command string, config *commandConfig, client *pokecache.Cache) {
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
		cmd.callback(config, client)
	}
}
