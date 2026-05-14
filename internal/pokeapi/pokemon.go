package pokeapi

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

func (c *Client) GetPokemon(name string) (Pokemon, error) {
	url := baseUrl + "/pokemon/" + name
	if val, ok := c.cache.Get(url); ok {
		cacheJson := Pokemon{}
		err := json.Unmarshal(val, &cacheJson)
		if err != nil {
			return Pokemon{}, err
		}
		return cacheJson, nil
	}

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return Pokemon{}, err
	}
	resp, err := c.httpClient.Do(req)
	if err != nil {
		return Pokemon{}, err
	}
	body, err := io.ReadAll(resp.Body)
	defer resp.Body.Close()
	if err != nil {
		return Pokemon{}, err
	}
	pokemon := Pokemon{}
	errJson := json.Unmarshal(body, &pokemon)
	if errJson != nil {
		fmt.Printf("%s was not found\n", name)
		return Pokemon{}, errJson
	}
	c.cache.Add(url, body)
	return pokemon, nil
}
